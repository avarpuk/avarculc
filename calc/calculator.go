package calculator

import (
	"math"
)

type Operations = map[string]func(float64, float64) float64

type Calculator struct {
	first     float64
	operation Operations
	second    float64
}

func NewCalculator(operations Operations) *Calculator {
	return &Calculator{
		first:     0,
		operation: operations,
		second:    0,
	}
}

func NewDefaultCalculator() *Calculator {
	operations := Operations{
		"+": sum,
		"-": sub,
		"/": div,
		"*": mul,
	}
	return NewCalculator(operations)
}

func sum(a float64, b float64) float64 {
	return a + b
}

func sub(a float64, b float64) float64 {
	return a - b
}

func div(a float64, b float64) float64 {
	if b == 0 {
		switch {
		case a > 0:
			return math.Inf(1)
		case a < 0:
			return math.Inf(-1)
		default:
			return math.NaN()
		}
	}

	return a / b
}

func mul(a float64, b float64) float64 {
	return a * b
}
