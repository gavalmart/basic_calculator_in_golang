package calculator_test

import (
	"calculator"
	"testing"
)

// func TestAdd(t *testing.T) {
// 	t.Parallel()
// 	var want float64 = 4
// 	got := calculator.Add(2, 2)

// 	if want != got {
// 		t.Errorf("want %f, got %f", want, got)
// 	}
// }

func TestAdd(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("Add(%f, %f):  want %f, got %f", tc.want, got, tc.want, got)
		}
	}

}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(2, 4)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestProduct(t *testing.T) {
	t.Parallel()
	var want float64 = 10
	got := calculator.Multiply(2, 5)

	if got != want {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: -1, b: -1, want: 1},
		{a: 10, b: 2, want: 5},
		{a: 10, b: 0, want: 5},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)

		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}

		if tc.want != got {
			t.Errorf("Divide (%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()

	_, err := calculator.Divide(1, 0)

	if err == nil {
		t.Error("want no error for invalid input, got nil")
	}
}
