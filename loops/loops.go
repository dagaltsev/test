package main

import "fmt"

type Person struct {
	Name     string
	Lastname string
	Age      int
	Email    string
}

func main() {

	fmt.Println("Loops in Golang")

	// for loop
	for i := 0; i < 5; i++ {
		//fmt.Println(i)
	}

	// while loop
	j := 10
	for j >= 5 {
		//fmt.Println(j)
		j--
	}

	a := 0
	for true {
		//fmt.Println(a)
		if a == 5 {
			break
		}
		a++
	}

	// range loop (for-each)
	//ages := []int{16, 20, 23, 44, 45, 70}

	// index and value
	// for i, v := range ages {
	// 	fmt.Println(i, v)
	// }

	// Create David person object
	david := Person{
		Name:     "David",
		Lastname: "Agaltsev",
		Age:      16,
		Email:    "david.agaltsev@example.com",
	}

	// Create Timo person object
	timo := Person{
		Name:     "Timo",
		Lastname: "Agaltsev",
		Age:      20,
		Email:    "timo.agaltsev@example.com",
	}

	// Create Max person object
	max := Person{
		Name:     "Max",
		Lastname: "Agaltsev",
		Age:      23,
		Email:    "max.agaltsev@example.com",
	}

	persons := []Person{david, timo, max}

	// for-each loop
	for _, person := range persons {
		out := fmt.Sprintf("%s %s is %d years old. Contact: %s", person.Name, person.Lastname, person.Age, person.Email)
		fmt.Println(out)
	}

	fmt.Println("Loop finished")
}
