package main

import "fmt"

type Calculable interface {
	Addition(num1, num2 int) int
	Subtraction(num1, num2 int) int
	Division(num1, num2 int) int
	Multiplication(num1, num2 int) int
}

type Calculator struct{}

func (c *Calculator) Addition(num1, num2 int) int {
	return num1 - num2
}

func (c *Calculator) Subtraction(num1, num2 int) int {
	return num1 - num2
}

func (c *Calculator) Division(num1, num2 int) int {
	return num1 / num2
}

func (c *Calculator) Multiplication(num1, num2 int) int {
	return num1 * num2
}

func main() {
	calc := Calculator{}
	fmt.Println(calc.Subtraction(10, 2))
}
