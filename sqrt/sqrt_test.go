package sqrt

import (
	"testing"
	"math"
)

func TestSqrtZero(t *testing.T) {
	testSqrt(0, t)
}

func TestSqrtPositiveOne(t *testing.T) {
	testSqrt(1, t)
}

func TestSqrtPositiveEven(t *testing.T) {
	testSqrt(2, t)
}

func TestSqrtPositiveOdd(t *testing.T) {
	testSqrt(5, t)
}

func TestSqrtPositiveBigNumber(t *testing.T) {
	testSqrt(1E10, t)
}

func TestSqrtPositiveSmallFloat(t *testing.T) {
	testSqrt(1E-6, t)
}

func TestSqrtPositiveBigFloat(t *testing.T) {
	testSqrt(100.0001, t)
}

func TestSqrtNegativeOne(t *testing.T) {
	testSqrt(-1, t)
}

func TestSqrtNegativeBigNumber(t *testing.T) {
	testSqrt(-1E10, t)
}

func TestSqrtNegativeSmallFloat(t *testing.T) {
	testSqrt(-1E-6, t)
}

func TestSqrtNegativeBigFloat(t *testing.T) {
	testSqrt(-100.0001, t)
}


func testSqrt(x float64, t *testing.T) error {
	const precision = 1E-6

	result := Sqrt(x)
	if x < 0 && !math.IsNaN(result) {
		t.Errorf("sqrt from negative number should NOT a NaN number. Got %.6f", result)
	}

	if result < 0 {
		t.Errorf("sqrt result should NOT be a negative number. Got %.6f", result)
	}

	expected := math.Sqrt(x)


	if math.Abs(result - expected) > precision {
		t.Errorf("the sqrt result %.6f is not equal to expected %.6f", result, expected)
	}

	return nil
}
