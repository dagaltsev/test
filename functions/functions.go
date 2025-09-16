package main

import "fmt"

func main() {
	PrintGreeting("David")
	PrintGreeting2("David", 16)
	PrintGreeting2("Timo", 20)
	PrintGreeting2("Max", 23)
	PrintGreeting2("Stani", 45)
	PrintGreeting2("Milan", 1)

	a := 100
	b := 200
	sum := CalculateSum(a, b)
	fmt.Printf("The sum of %d and %d is %d\n", a, b, sum)

	x, y := 5, 10
	x, y = SwapNumbers(x, y)
	fmt.Printf("After swapping: x = %d, y = %d\n", x, y)

	fmt.Println(ConvertNamesToArray("David", "Timo", "Max"))
	fmt.Println(ConvertNamesToArray("Stani", "Sabina", "Milan"))
	fmt.Println(ConvertNamesToArray("Marina", "Irina", "Karina"))
}

func CalculateSum(a, b int) int {
	return a + b
}

func SwapNumbers(x, y int) (int, int) {
	return y, x
}

// create function to calculate sum for 5 integers
func CalculateSum2(a, b, c, d, e int) int {
	return a + b + c + d + e
}

func PrintGreeting(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// create func to accept name and age and print "Hello <name>, you are <age> years old"
func PrintGreeting2(name string, age int) {
	fmt.Printf("Hello %s, you are %d years old\n", name, age)
}

// create func which accept 3 names and returns array of those names => []string{"David","Timo","Max"}
func ConvertNamesToArray(name1, name2, name3 string) []string {
	return []string{name1, name2, name3}
}
