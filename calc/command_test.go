package calc

import (
	"testing"
	"reflect"
)

func TestParseCommandType(t *testing.T) {
	testSuites := []struct {
		in       string
		expected CommandType
	}{
		{expected: UNKNOWN},
		{in: "add", expected: ADD},
		{in: "adD", expected: ADD},
		{in: "sub", expected: SUB},
		{in: "div", expected: DIV},
		{in: "dIv", expected: DIV},
		{in: "mul", expected: MUL},
		{in: "fib", expected: FIB},
		{in: "Fib", expected: FIB},
		{in: "kilimanjaro", expected: UNKNOWN},
	}

	for _, test := range testSuites {
		exist := ParseCommandType(test.in)
		if exist != test.expected {
			t.Error("failed to convert string '", test.in, "' to a command type. Got ", exist, ",  but expected is ", test.expected)
		}
	}
}

func TestNewCommand(t *testing.T) {
	testSuites := []struct {
		in       []string
		expected *Command
		err      bool
	}{
		{err: true},
		{in: []string{"add", "4", "5", "6"}, expected: &Command{command: ADD, values: []int64{4, 5, 6}}},
		{in: []string{"sub", "1"}, expected: &Command{command: SUB, values: []int64{1}}},
		{in: []string{"mul", "15", "-3", "-1", "10", "1000"}, expected: &Command{command: MUL, values: []int64{15, -3, -1, 10, 1000}}},
		{in: []string{"div", "5", "-34", "0"}, expected: &Command{command: DIV, values: []int64{5, -34, 0}}},
		{in: []string{"fib", "5"}, expected: &Command{command: FIB, values: []int64{5}}},
		{in: []string{"everest", "5", "-34", "0"}, err: true},
		{in: []string{"div"}, err: true},
		{in: []string{"fib", "-5"}, err: true},
		{in: []string{"fib", "-5", "45"}, err: true},
		{in: []string{"div", "-5", "sdfds"}, err: true},
		{in: []string{"div", "-5", ".34"}, err: true},
	}

	for _, test := range testSuites {
		exist, err := NewCommand(test.in)
		if test.err && err == nil {
			t.Error("failed to catch an error for args: ", test.in)
		} else if !test.err && err != nil {
			t.Error("failed to create new command with args ", test.in, ": ", err)
		} else if !reflect.DeepEqual(exist, test.expected) {
			t.Error("failed to create new command with args ", test.in, ". Got ", exist, ", but expected is ", test.expected)
		}
	}
}
