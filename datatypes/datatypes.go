package main

import "fmt"

func main() {
	// string
	// zeichenkette
	name := "David"
	lastname := "Agaltsev"
	age := "16"
	moto := "I love cookies"

	fmt.Printf("name variable type: %T, value: %s\n", name, name)
	fmt.Printf("lastname variable type: %T, value: %s\n", lastname, lastname)
	fmt.Printf("age variable type: %T, value: %s\n", age, age)
	fmt.Printf("moto variable type: %T, value: %s\n", moto, moto)
	fmt.Println()

	output := fmt.Sprintf("I am %s %s, %s years old. My moto is: %s", name, lastname, age, moto)
	fmt.Println(output)
	fmt.Println()

	moto = "I love \"cookies\""
	fmt.Println(moto)
	fmt.Println()

	// integer
	// ganzzahl
	agenumber := 16

	fmt.Printf("agenumber variable type: %T, value: %d\n", agenumber, agenumber)
	fmt.Println()

	output = fmt.Sprintf("I am %s %s, %d years old. My moto is: %s", name, lastname, agenumber, moto)
	fmt.Println(output)
	fmt.Println()

	// float
	// kommazahl
	exactage := 16.1

	fmt.Printf("exactage variable type: %T, value: %f\n", exactage, exactage)
	fmt.Println()

	output = fmt.Sprintf("I am %s %s, %f years old. My moto is: %s", name, lastname, exactage, moto)
	fmt.Println(output)
	fmt.Println()

	// boolean
	// true or false
	isbig := true
	issmall := false

	fmt.Printf("isbig variable type: %T, value: %t\n", isbig, isbig)
	fmt.Printf("issmall variable type: %T, value: %t\n", issmall, issmall)
	fmt.Println()

	// byte used mostly for JSON
	input := []byte(`{"name":"David", "lastname":"Agaltsev"}`)

	fmt.Printf("input variable type: %T, value: %s\n", input, string(input))
	fmt.Println()
}
