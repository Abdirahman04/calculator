package main

import "fmt"

func main() {
	var num1 float64
	var num2 float64

	fmt.Println("Enter two numbers:")

	fmt.Scanln(&num1)
	fmt.Scanln(&num2)

	sum := num1 + num2
	sub := num1 - num2
	mult := num1 * num2
	div := num1 / num2

	fmt.Printf("%v + %v = %v\n", num1, num2, sum)
	fmt.Printf("%v - %v = %v\n", num1, num2, sub)
	fmt.Printf("%v * %v = %v\n", num1, num2, mult)
	fmt.Printf("%v / %v = %v\n", num1, num2, div)
}
