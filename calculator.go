// package calculator does simple calculations
package calculator

import "errors"

// Add takes two numbers and returns the result of
// them together
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers a and b, and
// returns the result of subtracting b from a.
func Subtract(a, b float64) float64 {
	return b - a
}

//My own version
func Multiply(a, b float64) float64 {
	return a * b
}

// func Multiply(a, b float64) float64 {
// 	return 0
// }

func Divide(a, b float64) (float64, error) {

	if b == 0 {
		return 0, errors.New("division by zero not allowed")
	}

	return a / b, nil
}
