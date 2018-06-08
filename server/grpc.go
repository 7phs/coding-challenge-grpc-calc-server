package main

import (
	"context"
	"time"

	"github.com/7phs/coding-challenge-grpc-calc-server/api"
	"github.com/7phs/coding-challenge-grpc-calc-server/log"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/7phs/coding-challenge-grpc-calc-server/calc"
)

// A handler implemented gRPC calc service
type GrpcHandler struct{}

// Log a message of a result executing a request.
func (s *GrpcHandler) log(name string, param interface{}, v interface{}, duration time.Duration, err error) {
	if err != nil {
		log.Info("operation: "+name+"; error: ", err)
	} else {
		log.Info("operation: "+name+"(", param, ") = ", v, " in ", duration)
	}
}

// Converting an error to gRPC status.
func (s *GrpcHandler) err(err error) error {
	if err==nil {
		return nil
	}

	switch err.(type) {
	case *calc.ErrZeroDiv:
		return status.Error(codes.InvalidArgument, err.Error())
	default:
		return status.Error(codes.Internal, err.Error())
	}
}

func (s *GrpcHandler) Add(ctx context.Context, param *api.Values) (*api.Result, error) {
	start := time.Now()
	v, err := calc.Add(param.GetValues()...)

	s.log("Add", param.GetValues(), v, time.Since(start), err)

	return &api.Result{
		Value: v,
	}, s.err(err)
}

func (s *GrpcHandler) Sub(ctx context.Context, param *api.Values) (*api.Result, error) {
	start := time.Now()
	v, err := calc.Sub(param.GetValues()...)

	s.log("Sub", param.GetValues(), v, time.Since(start), err)

	return &api.Result{
		Value: v,
	}, s.err(err)
}

func (s *GrpcHandler) Mul(ctx context.Context, param *api.Values) (*api.Result, error) {
	start := time.Now()
	v, err := calc.Mul(param.GetValues()...)

	s.log("Mul", param.GetValues(), v, time.Since(start), err)

	return &api.Result{
		Value: v,
	}, s.err(err)
}

func (s *GrpcHandler) Div(ctx context.Context, param *api.Values) (*api.ResultDouble, error) {
	start := time.Now()
	v, err := calc.Div(param.GetValues()...)

	s.log("Div", param.GetValues(), v, time.Since(start), err)

	return &api.ResultDouble{
		Value: v,
	}, s.err(err)
}

func (s *GrpcHandler) Fib(ctx context.Context, param *api.Value) (*api.Result, error) {
	start := time.Now()
	v, err := calc.Fib(param.GetValue())

	s.log("Fib", param.GetValue(), v, time.Since(start), err)

	return &api.Result{
		Value: v,
	}, s.err(err)
}
