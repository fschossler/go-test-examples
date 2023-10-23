package main

import "testing"

func TestCalculatorAddiction(t *testing.T) {
	cal := Calculable(&Calculator{})
	result := cal.Addition(1, 3)
	got := result
	want := 4

	if got != want {
		t.Errorf("Error. Got %q and the right calculus is %q", got, want)
	}
}

func TestCalculatorSubtraction(t *testing.T) {
	cal := Calculable(&Calculator{})
	result := cal.Subtraction(10, 2)
	got := result
	want := 8

	if got != want {
		t.Errorf("Error. Got %q and the right calculus is %q", got, want)
	}
}

func TestCalculatorDivision(t *testing.T) {
	cal := Calculable(&Calculator{})
	result := cal.Division(10, 5)
	got := result
	want := 2

	if got != want {
		t.Errorf("Error. Got %q and the right calculus is %q", got, want)
	}
}

func TestCalculatorMultiplication(t *testing.T) {
	cal := Calculable(&Calculator{})
	result := cal.Multiplication(5, 10)
	got := result
	want := 50

	if got != want {
		t.Errorf("Error. Got %q and the right calculus is %q", got, want)
	}
}
