package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatorAdd(t *testing.T) {
	cal := Calculable(&Calculator{})
	result := cal.Addition(1, 3)

	assert.Equal(t, 4, result)
}

func TestCalculatorSubtraction(t *testing.T) {
	// var cal Calculable
	cal := Calculable(&Calculator{})
	result := cal.Subtraction(10, 3)

	assert.Equal(t, 7, result)
}

func TestCalculatorDivision(t *testing.T) {
	// var cal Calculable
	cal := Calculable(&Calculator{})
	result := cal.Division(10, 5)

	assert.Equal(t, 2, result)
}

func TestCalculatorMultiplication(t *testing.T) {
	// var cal Calculable
	cal := Calculable(&Calculator{})
	result := cal.Multiplication(2, 5)

	assert.Equal(t, 10, result)
}
