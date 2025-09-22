package main

import "fmt"

type Person struct {
	Name     string
	Lastname string
	Age      int
	Email    string
}

func main() {
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

	sabina := Person{
		Name:     "Sabina",
		Lastname: "Agaltsev",
		Age:      44,
		Email:    "sabina.agaltsev@example.com",
	}

	list := GeneratePersonsList(david, timo, max, sabina)

	fmt.Println(list)
}

// Deliver 5 Persons (David, Timo, Max, Sabina, Stani)
func GeneratePersons() []Person {
	var persons []Person

	David := Person{
		Name:     "David",
		Lastname: "Agaltsev",
		Age:      16,
		Email:    "David.agaltsev@example.com",
	}

	persons = append(persons, David)

	Timo := Person{
		Name:     "Timo",
		Lastname: "Agaltsev",
		Age:      20,
		Email:    "Timo.agaltsev@example.com",
	}

	persons = append(persons, Timo)

	Max := Person{
		Name:     "Max",
		Lastname: "Agaltsev",
		Age:      23,
		Email:    "Max.agaltsev@example.com",
	}

	persons = append(persons, Max)

	Sabina := Person{
		Name:     "Sabina",
		Lastname: "Agaltsev",
		Age:      44,
		Email:    "Sabina.agaltsev@example.com",
	}

	persons = append(persons, Sabina)

	Stani := Person{
		Name:     "Stani",
		Lastname: "Agaltsev",
		Age:      45,
		Email:    "Stani.agaltsev@example.com",
	}

	persons = append(persons, Stani)

	return persons
}

func GeneratePersonsList(persons ...Person) []Person {
	var list []Person
	list = append(list, persons...)

	return list
}
