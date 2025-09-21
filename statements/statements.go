package main

import "fmt"

func main() {
	CheckName2("Sabina")

	// PrintClass1(10) // "Du gehst in 4te Klasse"
	// PrintClass1(11) // "Du gehst in 5te Klasse"
	// PrintClass1(12) // "Du gehst in 6te Klasse"
	// PrintClass1(13) // "Du gehst in 7te Klasse"
	// PrintClass1(14) // "Du gehst in 8te Klasse"
	// PrintClass1(15) // "Du gehst in 9te Klasse"
	// PrintClass1(16) // "Du gehst in 10te Klasse"
	// PrintClass1(17) // "Du gehst in 11te Klasse"
	// PrintClass1(18) // "Du gehst in 12te Klasse"
}

func CheckName2(name string) {
	switch name {
	case "David": // if name == "David"
		fmt.Println("Hello David")
	case "Timo": // if name == "Timo"
		fmt.Println("Hello Timo")
	case "Max": // if name == "Max"
		fmt.Println("Hello Max")
	default: // else
		fmt.Printf("Hallo %s\n", name)
	}
}

func CheckName1(name string) {
	if name == "David" {
		fmt.Println("Hello David")
	} else if name == "Timo" {
		fmt.Println("Hello Timo")
	} else if name == "Max" {
		fmt.Println("Hello Max")
	} else {
		fmt.Printf("Hallo %s\n", name)
	}
}

func CanDrive1(age int) {
	if age < 17 {
		fmt.Println("Du darfst noch nicht Auto fahren")
	} else if age == 17 {
		fmt.Println("Du darfst mit deinen Eltern Auto fahren")
	} else {
		fmt.Println("Du darfst alleine Auto fahren")
	}
}

func CanDrive2(age int) {
	switch {
	case age < 17: // if age < 17
		fmt.Println("Du darfst noch nicht Auto fahren")
	case age == 17: // if age == 17
		fmt.Println("Du darfst mit deinen Eltern Auto fahren")
	default: // else
		fmt.Println("Du darfst alleine Auto fahren")
	}
}

func ValidateAge(age int) {
	if age < 18 {
		fmt.Println("You are not allowed to vote")
	} else {
		fmt.Println("You are allowed to vote")
	}
}

// create func
// which accepts age as int
// based on age you print class for primary school (age 10 -> "Du gehst in 4te Klasse", etc)
// if, else if, else
func PrintClass1(age int) {}

// switch case
func PrintClass2(age int) {}
