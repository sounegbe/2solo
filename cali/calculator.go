package main

import (
	"fmt"
	"math"
)

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

func Power(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}

func Modulus(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("modulus by zero is not allowed")
	}
	return math.Mod(a, b), nil
}

func SquareRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, fmt.Errorf("square root of negative number is not allowed")
	}
	return math.Sqrt(num), nil
}

func Average(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("cannot calculate average of empty list")
	}

	sum := 0.0
	for _, num := range numbers {
		sum += num
	}

	return sum / float64(len(numbers)), nil
}
