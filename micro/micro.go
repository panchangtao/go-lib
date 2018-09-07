package micro

import (
	"bytes"
	"container/list"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"text/tabwriter"
	"time"

	"github.com/arcplus/go-lib/log"
)

// go build -ldflags -X
var version, gitCommit, buildDate string

// VersionInfo return visualized version info
func VersionInfo() string {
	buff := &bytes.Buffer{}
	w := tabwriter.NewWriter(buff, 0, 0, 0, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "version: \t"+version)
	fmt.Fprintln(w, "gitCommit: \t"+gitCommit)
	fmt.Fprintln(w, "buildDate: \t"+buildDate) // trailing tab
	w.Flush()
	return buff.String()
}

type Micro interface {
	AddResCloseFunc(f func() error)
	Close()
	ServeGRPC(bindAddr string, server GRPCServer)
	ServeHTTP(bindAddr string, handler http.Handler)
	Start()
}

type micro struct {
	mu            *sync.Mutex
	errChan       chan error
	serveFuncs    []func()
	resCloseFuncs *list.List
}

// New create Micro, moduleName.0 is module name.
func New(moduleName ...string) Micro {
	m := &micro{
		mu:            &sync.Mutex{},
		errChan:       make(chan error, 1),
		serveFuncs:    make([]func(), 0),
		resCloseFuncs: list.New(),
	}

	if len(moduleName) != 0 {
		log.SetAttachment(map[string]string{
			"module": moduleName[0],
		})
	}

	if rds, key := os.Getenv("log_rds"), os.Getenv("log_key"); rds != "" && key != "" {
		level := log.InfoLevel
		if os.Getenv("log_level") == "debug" {
			level = log.DebugLevel
		}

		log.SetOutput(log.RedisWriter(log.RedisConfig{
			Level:  level,
			DSN:    rds,
			LogKey: key,
			Async:  true,
		}))
	}

	m.AddResCloseFunc(log.Close)

	return m
}

func SetLogger(rdsDSN string, moduleName string, mode string) {

}

// Close close all added resource FILO
func (m *micro) Close() {
	m.mu.Lock()
	for m.resCloseFuncs.Len() != 0 {
		e := m.resCloseFuncs.Back()
		if f, ok := e.Value.(func() error); ok && f != nil {
			err := f()
			if err != nil {
				log.Errorf("close resource err: %s", err.Error())
			}
		}
		m.resCloseFuncs.Remove(e)
	}
	m.mu.Unlock()
}

// TODO ln reuse?
func (m *micro) createListener(bindAddr string) (net.Listener, error) {
	m.mu.Lock()
	ln, err := net.Listen("tcp", bindAddr)
	if err != nil {
		m.mu.Unlock()
		return nil, err
	}
	m.mu.Unlock()

	m.AddResCloseFunc(func() error {
		err := ln.Close()
		if err != nil {
			if _, ok := err.(*net.OpError); ok {
				return nil
			}
			return err
		}
		return nil
	})

	return ln, nil
}

// AddResCloseFunc add resource close func
func (m *micro) AddResCloseFunc(f func() error) {
	m.mu.Lock()
	m.resCloseFuncs.PushBack(f)
	m.mu.Unlock()
}

// GRPCServer
type GRPCServer interface {
	Serve(net.Listener) error
	GracefulStop()
}

// ServeGRPC is helper func to start gRPC server
func (m *micro) ServeGRPC(bindAddr string, server GRPCServer) {
	m.serveFuncs = append(m.serveFuncs, func() {
		ln, err := m.createListener(bindAddr)
		if err != nil {
			m.errChan <- err
			return
		}

		m.AddResCloseFunc(func() error {
			server.GracefulStop()
			return nil
		})

		err = server.Serve(ln)
		if err != nil {
			m.errChan <- err
		}
	})
}

// TODO other params can optimize
func (m *micro) ServeHTTP(bindAddr string, handler http.Handler) {
	m.serveFuncs = append(m.serveFuncs, func() {
		ln, err := m.createListener(bindAddr)
		if err != nil {
			m.errChan <- err
			return
		}

		server := &http.Server{
			Handler:        handler,
			ReadTimeout:    30 * time.Second,
			WriteTimeout:   30 * time.Second,
			MaxHeaderBytes: 2 << 15, // 64k
		}

		m.AddResCloseFunc(func() error {
			ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
			defer cancelFunc()
			return server.Shutdown(ctx)
		})

		err = server.Serve(ln)
		if err != nil {
			m.errChan <- err
		}
	})
}

// WatchSignal notify signal to stop running
var WatchSignal = []os.Signal{syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT}

// Wait util signal
func (m *micro) Start() {
	defer m.Close()

	for i := range m.serveFuncs {
		go m.serveFuncs[i]()
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, WatchSignal...)
	select {
	case s := <-ch:
		log.Skip(1).Infof("receive stop signal: %s", s)
	case e := <-m.errChan:
		log.Skip(1).Errorf("receive err signal: %s", e)
	}
}
