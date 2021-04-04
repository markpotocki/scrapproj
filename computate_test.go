package main

import (
	"fmt"
	"math"
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		name     string
		val      int
		expected int
	}{
		{"n = 1", 1, 1},
		{"n = 5", 5, 120},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := factorial(test.val)
			if actual != test.expected {
				t.Fail()
			}
		})
	}
}

func TestComputeValue(t *testing.T) {
	// compute value is X ^ N / N!
	tests := []struct {
		name     string
		valA     int
		valB     int
		expected float64
	}{
		{name: "2, 1", valA: 2, valB: 1, expected: 2},
		{name: "2, 4", valA: 2, valB: 4, expected: .666666666666666667},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := computeValue(test.valA, test.valB)
			ep := .00000000001
			if math.Abs(actual-test.expected) > ep {
				t.Fail()
			}
		})
	}
}

func BenchmarkGoroutineCompute(b *testing.B) {
	accLimit := 200
	exponent := 2
	for i := 1; i < accLimit; i++ {
		b.Run("GoroutineCompute-Acc"+fmt.Sprint(i), func(b *testing.B) {
			b.ReportAllocs()
			for j := 0; j < b.N; j++ {
				multiThreadCompute(exponent, i)
			}
		})
	}
}

func BenchmarkIterativeCompute(b *testing.B) {
	accLimit := 200
	exponent := 2
	for i := 1; i < accLimit; i++ {
		b.Run("IterativeCompute-Acc"+fmt.Sprint(i), func(b *testing.B) {
			b.ReportAllocs()
			for j := 0; j < b.N; j++ {
				iterativeCompute(exponent, j)
			}
		})
	}
}
