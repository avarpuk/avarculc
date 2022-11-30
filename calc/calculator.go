package calculator

import (
	"math"
)

type Operations = map[string]func(float64, float64) float64

type Calculator struct {
	operation Operations
	result    float64
}

func (c *Calculator) GetResult() float64 {
	return c.result
}

func NewCalculator(operations Operations) *Calculator {
	return &Calculator{
		operation: operations,
		result:    0,
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

func (c *Calculator) Run(first float64, operation string, second float64) error {
	if f, ok := c.operation[operation]; ok {
		c.result = f(first, second)
		return nil
	}
	return nil
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
