package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatorAddition(t *testing.T) {
	testCases := []struct {
		num1, num2, expected int
	}{
		{1, 2, 3},
		{3, 2, 5},
		{32123, 3, 32126},
		{10, 10, 20},
	}

	for _, testCase := range testCases {
		t.Run("Test cases for Addition", func(t *testing.T) {
			calculator := Calculable(Calculator{})
			result := calculator.Addition(testCase.num1, testCase.num2)
			assert.Equal(t, testCase.expected, result)
		})
	}
}

func TestCalculatorSubtraction(t *testing.T) {
	testCases := []struct {
		num1, num2, expected int
	}{
		{1, 2, -1},
		{3, 2, 1},
		{32123, 3, 32120},
		{10, 10, 0},
	}

	for _, testCase := range testCases {
		t.Run("Test cases for Subtraction", func(t *testing.T) {
			calculator := Calculable(Calculator{})
			result := calculator.Subtraction(testCase.num1, testCase.num2)
			assert.Equal(t, testCase.expected, result)
		})
	}
}

func TestCalculatorDivision(t *testing.T) {
	testCases := []struct {
		num1, num2, expected int
	}{
		{10, 5, 2},
		{50, 10, 5},
		{1000, 2, 500},
		{10, 10, 1},
	}

	for _, testCase := range testCases {
		t.Run("Test cases for Division", func(t *testing.T) {
			calculator := Calculable(Calculator{})
			result := calculator.Division(testCase.num1, testCase.num2)
			assert.Equal(t, testCase.expected, result)
		})
	}
}

func TestCalculatorMultiplication(t *testing.T) {
	testCases := []struct {
		num1, num2, expected int
	}{
		{10, 5, 50},
		{50, 10, 500},
		{1000, 2, 2000},
		{10, 10, 100},
	}

	for _, testCase := range testCases {
		t.Run("Test cases for Multiplication", func(t *testing.T) {
			calculator := Calculable(Calculator{})
			result := calculator.Multiplication(testCase.num1, testCase.num2)
			assert.Equal(t, testCase.expected, result)
		})
	}
}
