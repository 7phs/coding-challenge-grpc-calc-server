package main

import (
	"context"
	"time"
	"net"
	"fmt"
	"os"
	"math/rand"

	"google.golang.org/grpc"

	"github.com/7phs/coding-challenge-grpc-calc-server/api"
	"sync"
	"testing"
	"github.com/7phs/coding-challenge-grpc-calc-server/config"
)

const (
	TEST_TRYING_LISTENING = 20
)

type TestGrpcHandler struct {
	timeout time.Duration
	value   int64
	err     error
}

func (s *TestGrpcHandler) SetTimeout(timeout time.Duration) {
	s.timeout = timeout
}

func (s *TestGrpcHandler) SetValue(value int64) {
	s.value = value
}

func (s *TestGrpcHandler) SetError(err error) {
	s.err = err
}

func (s *TestGrpcHandler) Add(ctx context.Context, param *api.Values) (*api.Result, error) {
	time.Sleep(s.timeout)

	return &api.Result{
		Value: s.value,
	}, s.err
}

func (s *TestGrpcHandler) Sub(ctx context.Context, param *api.Values) (*api.Result, error) {
	time.Sleep(s.timeout)

	return &api.Result{
		Value: s.value,
	}, s.err
}

func (s *TestGrpcHandler) Mul(ctx context.Context, param *api.Values) (*api.Result, error) {
	time.Sleep(s.timeout)

	return &api.Result{
		Value: s.value,
	}, s.err
}

func (s *TestGrpcHandler) Div(ctx context.Context, param *api.Values) (*api.ResultDouble, error) {
	time.Sleep(s.timeout)

	return &api.ResultDouble{
		Value: float64(s.value),
	}, s.err
}

func (s *TestGrpcHandler) Fib(ctx context.Context, param *api.Value) (*api.Result, error) {
	time.Sleep(s.timeout)

	return &api.Result{
		Value: s.value,
	}, s.err
}

func SetUpMockServer() (port int, handler *TestGrpcHandler, deferFunc func(), err error) {
	var (
		listener net.Listener
		wait     sync.WaitGroup
	)
	// try to listen a port
	port, listener, err = func() (int, net.Listener, error) {
		for i := 0; i < TEST_TRYING_LISTENING; i++ {
			port := 10000 + rand.Int31n(30000)

			listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
			if err == nil {
				return int(port), listener, nil
			}
		}

		return 0, nil, os.ErrNotExist
	}()

	handler = &TestGrpcHandler{}
	// create a gRPC server object
	server := grpc.NewServer()
	// attach the Ping service to the server
	api.RegisterCalcServer(server, handler)

	// start the server
	wait.Add(1)
	go func() {
		server.Serve(listener)

		wait.Done()
	}()

	deferFunc = func() {
		listener.Close()
		wait.Wait()
	}

	return
}

func TestClient_Execute(t *testing.T) {
	port, handler, deferFunc, err := SetUpMockServer()
	if err != nil {
		t.Error("failed to set up a mock server: ", err)
		return
	}
	defer deferFunc()

	client := NewClient((&config.Config{}).
		SetPort(port))

	testSuites := []struct {
		command     []string
		expected    int64
		timeout     time.Duration
		err         bool
		expectedErr bool
	}{
		{command: []string{"add", "2", "3"}, expected: 5},
		{command: []string{"sub", "100", "10", "5"}, expected: 85},
		{command: []string{"mul", "10", "-10", "-5"}, expected: 500},
		{command: []string{"div", "100", "10"}, expected: 10},
		{command: []string{"fib", "2"}, expected: 3},
		{command: []string{"div", "100", "0"}, err: true, expectedErr: true},
		{command: []string{"add", "100", "0"}, timeout: 200 * time.Millisecond, expectedErr: true},
	}

	for _, test := range testSuites {
		handler.SetValue(test.expected)

		if test.timeout > 0 {
			client.timeout = test.timeout
			handler.SetTimeout(test.timeout+ 100 * time.Millisecond)
		} else {
			client.timeout = 50 * time.Millisecond
			handler.SetTimeout(0)
		}

		if test.err {
			handler.SetError(os.ErrInvalid)
		} else {
			handler.SetError(nil)
		}

		command, err := NewCommand(test.command)
		if err!=nil {
			t.Error("failed to create a command for ", test.command, ": ", err)
			continue
		}

		response, err := client.Execute(command)
		if test.expectedErr && err==nil {
			t.Error("failed to catch an error for ", test.command)
		} else if !test.expectedErr && err!=nil {
			t.Error("failed to execute a command ", test.command, ": ", err)
		} else {
			exist := func() int64 {
				switch v := response.(type) {
				case int64:
						return v
				case float64:
						return int64(v)
				default:
					return 0
				}
			}()

			if exist!=test.expected {
				t.Error("failed to execute a command ", test.command, ". Got ", exist, ", but expected is ", test.expected)
			}
		}
	}
}
