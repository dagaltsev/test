package main

import "fmt"

func main() {
	PrintGreeting("David")

	a := 100
	b := 200
	sum := CalculateSum(a, b)
	fmt.Printf("The sum of %d and %d is %d\n", a, b, sum)
}

func PrintGreeting(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func CalculateSum(a, b int) int {
	return a + b
}
