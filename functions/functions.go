package main

import "fmt"

func main() {
	PrintGreeting("David")

	a := 100
	b := 200
	sum := CalculateSum(a, b)
	fmt.Printf("The sum of %d and %d is %d\n", a, b, sum)

	x, y := 5, 10
	x, y = SwapNumbers(x, y)
	fmt.Printf("After swapping: x = %d, y = %d\n", x, y)
}

func PrintGreeting(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func CalculateSum(a, b int) int {
	return a + b
}

func SwapNumbers(x, y int) (int, int) {
	return y, x
}
