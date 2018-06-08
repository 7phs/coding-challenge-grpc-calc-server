package calc

import (
	"testing"
	"math"
	)

const (
	TEST_EPS = 1E-7
)

func TestAdd(t *testing.T) {
	testSuites := []struct {
		in       []int64
		expected int64
	}{
		{},
		{in: []int64{12}, expected: 12},
		{in: []int64{-12, -23}, expected: -35},
		{in: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, expected: 45},
	}

	for _, test := range testSuites {
		exist, err := Add(test.in...)
		if err!=nil {
			t.Error("failed to execute a command 'add': ", err)
		} else if exist!=test.expected {
			t.Error("failed to execute a command 'add'. Got ", exist, ", but expected is ", test.expected)
		}
	}
}

func TestSub(t *testing.T) {
	testSuites := []struct {
		in       []int64
		expected int64
	}{
		{},
		{in: []int64{12}, expected: 12},
		{in: []int64{-12, -23}, expected: 11},
		{in: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, expected: -43},
		{in: []int64{-1, -2, -3, -4, -5, -6, -7, -8, -9, 0}, expected: 43},
	}

	for _, test := range testSuites {
		exist, err := Sub(test.in...)
		if err!=nil {
			t.Error("failed to execute a command 'subtract': ", err)
		} else if exist!=test.expected {
			t.Error("failed to execute a command 'subtract'. Got ", exist, ", but expected is ", test.expected)
		}
	}
}

func TestMul(t *testing.T) {
	testSuites := []struct {
		in       []int64
		expected int64
	}{
		{},
		{in: []int64{12}, expected: 12},
		{in: []int64{-12, -23}, expected: 276},
		{in: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, expected: 0},
		{in: []int64{-1, -2, -3, -4, -5, -6, -7, -8, -9}, expected: -362880},
	}

	for _, test := range testSuites {
		exist, err := Mul(test.in...)
		if err!=nil {
			t.Error("failed to execute a command 'multiply': ", err)
		} else if exist!=test.expected {
			t.Error("failed to execute a command 'multiply'. Got ", exist, ", but expected is ", test.expected)
		}
	}
}

func TestDiv(t *testing.T) {
	testSuites := []struct {
		in       []int64
		expected float64
		err      bool
	}{
		{},
		{in: []int64{12}, expected: 12},
		{in: []int64{-24, -12}, expected: 2.},
		{in: []int64{10000000000, 20, 30, 40, 50, 60, 70, 80, 90}, expected: 0.00027557319223985884},
		{in: []int64{-100000000, -2, -30, -4, -50, -6, -70, -8, -90}, expected: -0.027557319223985893},
		{in: []int64{100, 2, 0, 5}, err: true},
	}

	for _, test := range testSuites {
		exist, err := Div(test.in...)

		if test.err && err==nil {
			t.Error("failed to catch an error while execute a command 'divide' ", test.in)
		} else if !test.err && err!=nil {
			t.Error("failed to execute a command 'divide': ", err)
		} else if math.Abs(exist-test.expected)>TEST_EPS {
			t.Error("failed to execute a command 'divide'. Got ", exist, ", but expected is ", test.expected)
		}
	}
}
