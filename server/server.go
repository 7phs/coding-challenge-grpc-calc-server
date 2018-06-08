package main

import (
	"net"
	"fmt"
	"errors"
	"google.golang.org/grpc"

	"github.com/7phs/coding-challenge-grpc-calc-server/api"
	"sync"
	"github.com/7phs/coding-challenge-grpc-calc-server/config"
)

// A server listened for a gRPC connection and processing requests.
type Server struct {
	listener  net.Listener
	wait      sync.WaitGroup

	errorWait chan struct{}
	err       error
}

// Creates a server and trying listening a port.
func NewServer(config *config.Config) (*Server, error) {
	// create a listener on TCP port 7777
	listener, err := net.Listen("tcp", config.Address())
	if err != nil {
		return nil, errors.New(fmt.Sprint("failed to listen ", config.Address(), ":", err))
	}

	return &Server{
		listener:  listener,
		errorWait: make(chan struct{}),
	}, nil
}

// Raise an error - store an error and close an error channel to process that case.
func (o *Server) raiseError(err error) {
	o.err = err
	close(o.errorWait)
}

// Return a channel to make a signal of an error case.
func (o *Server) WaitError() chan struct{} {
	return o.errorWait
}

// An error happened while server was working.
func (o *Server) Error() error {
	return o.err
}

// Running main goroutine of processing gRPC requests.
func (o *Server) Run() {
	// create a gRPC server object
	server := grpc.NewServer()
	// attach the Ping service to the server
	api.RegisterCalcServer(server, &GrpcHandler{})

	// start the server
	o.wait.Add(1)
	go func() {
		if err := server.Serve(o.listener); err != nil {
			o.raiseError(errors.New(fmt.Sprint("failed to serve:", err)))
		}

		o.wait.Done()
	}()
}

// Shutdown the server.
func (o *Server) Shutdown() {
	o.listener.Close()
}

// Wait for completely stopped.
func (o *Server) Wait() {
	o.wait.Wait()
}
