package calc

// Calculates a sum of all integer arguments and return it.
// Doesn't check overflow, but will return 'err' when implementing it.
func Add(values ... int64) (result int64, err error) {
	for _, v := range values {
		result += v
	}

	return
}

// Calculates subtraction of all integer arguments from the first arguments and return it.
// Doesn't check overflow, but will return 'err' when implementing it.
func Sub(values ... int64) (result int64, err error) {
	if len(values) == 0 {
		return
	}

	result = values[0]

	for _, v := range values[1:] {
		result -= v
	}

	return
}

// Multiplies of all integer arguments and return the result.
// Doesn't check overflow, but will return 'err' when implementing it.
func Mul(values ... int64) (result int64, err error) {
	if len(values) == 0 {
		return
	}

	result = values[0]

	for _, v := range values[1:] {
		if v == 0 {
			result = 0
			break
		}

		result *= v
	}

	return
}

// Calculates division the first arguments by another integer arguments and return the result.
// Doesn't check overflow, but will return 'err' when implementing it.
// Will return error 'ErrZeroDiv' if one of arguments is zero.
func Div(values ... int64) (result float64, err error) {
	if len(values) == 0 {
		return
	}

	result = float64(values[0])

	for _, v := range values[1:] {
		if v == 0 {
			result = .0
			err = NewErrZeroDiv("division by zero")
			break
		}

		result /= float64(v)
	}

	return
}
