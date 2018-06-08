package main

import (
	"context"
	"time"
	"errors"
	"fmt"

	"google.golang.org/grpc"

	"github.com/7phs/coding-challenge-grpc-calc-server/config"
	"github.com/7phs/coding-challenge-grpc-calc-server/api"
	"github.com/7phs/coding-challenge-grpc-calc-server/calc"
)

// A helper for calling an arithmetic operation a gRPC service.
type Client struct {
	address string
	timeout time.Duration
}

// Creating the new client wrapper, getting and storing parameters from a config
func NewClient(conf *config.Config) *Client {
	return &Client{
		address: conf.Address(),
		timeout: conf.Timeout(),
	}
}

// Make a connection to a service, route a call gRPC using a command type.
func (o *Client) Execute(command *calc.Command) (interface{}, error) {
	conn, err := grpc.Dial(o.address, grpc.WithInsecure())
	if err != nil {
		return 0, errors.New(fmt.Sprint("failed to connect to server: ", err))
	}
	defer conn.Close()

	ctx, _ := context.WithTimeout(context.Background(), o.timeout)

	client := api.NewCalcClient(conn)
	response, err := func() (interface{}, error) {
		switch command.Command() {
		case calc.ADD:
			r, err := client.Add(ctx, &api.Values{Values: command.Values()})
			return r.GetValue(), err
		case calc.SUB:
			r, err := client.Sub(ctx, &api.Values{Values: command.Values()})
			return r.GetValue(), err
		case calc.MUL:
			r, err := client.Mul(ctx, &api.Values{Values: command.Values()})
			return r.GetValue(), err
		case calc.DIV:
			r, err := client.Div(ctx, &api.Values{Values: command.Values()})
			return r.GetValue(), err
		case calc.FIB:
			r, err := client.Fib(ctx, &api.Value{Value: uint64(command.Values()[0])})
			return r.GetValue(), err
		}

		return 0, errors.New("unknown command")
	}()

	if err != nil {
		return 0, err
	}

	return response, err
}
