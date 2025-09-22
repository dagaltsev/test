package main

import "fmt"

type Person struct {
	Name     string
	Lastname string
	Age      int
	Email    string
}

func main() {
	var persons []Person

	//fmt.Println(persons)

	// Create David person object
	david := Person{
		Name:     "David",
		Lastname: "Agaltsev",
		Age:      16,
		Email:    "david.agaltsev@example.com",
	}

	// Add david to persons slice
	persons = append(persons, david)

	// Create Timo person object
	timo := Person{
		Name:     "Timo",
		Lastname: "Agaltsev",
		Age:      20,
		Email:    "timo.agaltsev@example.com",
	}

	// Add Timo to persons slice
	persons = append(persons, timo)

	fmt.Println(GeneratePersons())
}

// Deliver 5 Persons (David, Timo, Max, Sabina, Stani)
func GeneratePersons() []Person {
	var persons []Person

	return persons
}
