package calc

import "testing"

func TestFib(t *testing.T) {
	testSuites := []struct {
		in       uint64
		expected int64
	}{
		{},
		{in: 12, expected: 144},
		{in: 24, expected: 46368},
		{in: 15, expected: 610},
	}

	for _, test := range testSuites {
		exist, err := Fib(test.in)
		if err != nil {
			t.Error("failed to execute a command 'fib': ", err)
		} else if exist != test.expected {
			t.Error("failed to execute a command 'fib'. Got ", exist, ", but expected is ", test.expected)
		}
	}

}
