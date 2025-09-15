package main

import "fmt"

func main() {
	// key => value pair
	// array index 0 => David :: names[0]
	// array index 1 => Timo  :: names[1]
	// array index 2 => Max   :: names[2]
	names := []string{"David", "Timo", "Max"}

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// array index 0 => 1 :: numbers[0]
	// array index 1 => 2 :: numbers[1]
	// array index 2 => 3 :: numbers[2]
	// array index 3 => 4 :: numbers[3]
	// array index 4 => 5 :: numbers[4]
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(numbers)

	// Calculate sum for all numbers values
	sum := numbers[0] + numbers[1] + numbers[2] + numbers[3] + numbers[4]

	fmt.Printf("The sum for numbers array is %d\n", sum)

	fmt.Println("---------------------------------------------------------------------------------------------")
	// given an array with names: Max, Timo, David
	n := []string{"Max", "Timo", "David"}

	// given array with ages: 23, 20, 16
	a := []int{23, 20, 16}

	// create 3 outputs for each brother and his age
	fmt.Printf("%s is %d years old\n", n[0], a[0])
	fmt.Printf("%s is %d years old\n", n[1], a[1])
	fmt.Printf("%s is %d years old\n", n[2], a[2])
}
