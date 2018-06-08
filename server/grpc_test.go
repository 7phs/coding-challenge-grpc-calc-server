package main

import (
	"testing"
	"github.com/7phs/coding-challenge-grpc-calc-server/api"
	"context"
	"github.com/7phs/coding-challenge-grpc-calc-server/log"
	"math"
)

func test_hideLogs() func() {
	level := log.GetLogLevel()
	log.SetLogLevel(log.ERROR)

	return func() {
		log.SetLogLevel(level)
	}
}

func TestGrpcHandler_Add(t *testing.T) {
	defer test_hideLogs()()

	s := &GrpcHandler{}

	testSuites := []struct {
		in       *api.Values
		expected int64
	}{
		{},
		{in: &api.Values{Values: []int64{12}}, expected: 12},
		{in: &api.Values{Values: []int64{-12, -23}}, expected: -35},
		{in: &api.Values{Values: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}}, expected: 45},
	}

	for _, test := range testSuites {
		resp, err := s.Add(context.Background(), test.in)
		if err != nil {
			t.Error("failed to execute a command 'add': ", err)
		} else if exist := resp.GetValue(); exist != test.expected {
			t.Error("failed to execute a command 'add'. Got ", exist, ", but expected is ", test.expected)
		}
	}
}

func TestGrpcHandler_Sub(t *testing.T) {
	defer test_hideLogs()()

	s := &GrpcHandler{}

	testSuites := []struct {
		in       *api.Values
		expected int64
	}{
		{},
		{in: &api.Values{Values: []int64{12}}, expected: 12},
		{in: &api.Values{Values: []int64{-12, -23}}, expected: 11},
		{in: &api.Values{Values: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}}, expected: -43},
		{in: &api.Values{Values: []int64{-1, -2, -3, -4, -5, -6, -7, -8, -9, 0}}, expected: 43},
	}

	for _, test := range testSuites {
		resp, err := s.Sub(context.Background(), test.in)
		if err != nil {
			t.Error("failed to execute a command 'sub': ", err)
		} else if exist := resp.GetValue(); exist != test.expected {
			t.Error("failed to execute a command 'sub'. Got ", exist, ", but expected is ", test.expected)
		}
	}
}

func TestGrpcHandler_Mul(t *testing.T) {
	defer test_hideLogs()()

	s := &GrpcHandler{}

	testSuites := []struct {
		in       *api.Values
		expected int64
	}{
		{},
		{in: &api.Values{Values: []int64{12}}, expected: 12},
		{in: &api.Values{Values: []int64{-12, -23}}, expected: 276},
		{in: &api.Values{Values: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}}, expected: 0},
		{in: &api.Values{Values: []int64{-1, -2, -3, -4, -5, -6, -7, -8, -9}}, expected: -362880},
	}

	for _, test := range testSuites {
		resp, err := s.Mul(context.Background(), test.in)
		if err != nil {
			t.Error("failed to execute a command 'mul': ", err)
		} else if exist := resp.GetValue(); exist != test.expected {
			t.Error("failed to execute a command 'mul'. Got ", exist, ", but expected is ", test.expected)
		}
	}
}

func TestGrpcHandler_Div(t *testing.T) {
	defer test_hideLogs()()

	s := &GrpcHandler{}

	testSuites := []struct {
		in       *api.Values
		expected float64
		err      bool
	}{
		{},
		{in: &api.Values{Values: []int64{12}}, expected: 12},
		{in: &api.Values{Values: []int64{-24, -12}}, expected: 2.},
		{in: &api.Values{Values: []int64{10000000000, 20, 30, 40, 50, 60, 70, 80, 90}}, expected: 0.00027557319223985884},
		{in: &api.Values{Values: []int64{-100000000, -2, -30, -4, -50, -6, -70, -8, -90}}, expected: -0.027557319223985893},
		{in: &api.Values{Values: []int64{100, 2, 0, 5}}, err: true},
	}

	for _, test := range testSuites {
		resp, err := s.Div(context.Background(), test.in)
		if err != nil {
			if test.err && err == nil {
				t.Error("failed to catch an error while execute a command 'divide' ", test.in)
			} else if !test.err && err != nil {
				t.Error("failed to execute a command 'divide': ", err)
			} else if exist := resp.GetValue(); math.Abs(exist-test.expected) > TEST_EPS {
				t.Error("failed to execute a command 'divide'. Got ", exist, ", but expected is ", test.expected)
			}
		}
	}
}

func TestGrpcHandler_Fib(t *testing.T) {
	defer test_hideLogs()()

	s := &GrpcHandler{}

	testSuites := []struct {
		in       *api.Value
		expected int64
		err      bool
	}{
		{},
		{in: &api.Value{Value: 12}, expected: 144},
		{in: &api.Value{Value: 24}, expected: 46368},
		{in: &api.Value{Value: 15}, expected: 610},
	}

	for _, test := range testSuites {
		resp, err := s.Fib(context.Background(), test.in)
		if err != nil {
			t.Error("failed to execute a command 'fib': ", err)
		} else if exist := resp.GetValue(); exist != test.expected {
			t.Error("failed to execute a command 'fib'. Got ", exist, ", but expected is ", test.expected)
		}
	}
}
