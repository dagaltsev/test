package main

import "fmt"

func main() {
	CheckName2("Sabina")

	PrintClass2(18) // "Du gehst in 4te Klasse"
	// PrintClass1(11) // "Du gehst in 5te Klasse"
	// PrintClass1(12) // "Du gehst in 6te Klasse"
	// PrintClass1(13) // "Du gehst in 7te Klasse"
	// PrintClass1(14) // "Du gehst in 8te Klasse"
	// PrintClass1(15) // "Du gehst in 9te Klasse"
	// PrintClass1(16) // "Du gehst in 10te Klasse"
	// PrintClass1(17) // "Du gehst in 11te Klasse"
	// PrintClass1(18) // "Du gehst in 12te Klasse"

	fmt.Println(CompareStrings("David geht in die 10te Klasse", "David geht in die 10te Klasse"))
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
func PrintClass1(age int) {
	if age == 10 {
		fmt.Println("Du gehst in die 4te Klasse")
	} else if age == 11 {
		fmt.Println("Du gehst in die 5te Klasse")
	} else if age == 12 {
		fmt.Println("Du gehst in die 6te Klasse")
	} else if age == 13 {
		fmt.Println("Du gehst in die 7te Klasse")
	} else if age == 14 {
		fmt.Println("Du gehst in die 8te Klasse")
	} else if age == 15 {
		fmt.Println("Du gehst in die 9te Klasse")
	} else if age == 16 {
		fmt.Println("Du gehst in die 10te Klasse")
	} else if age == 17 {
		fmt.Println("Du gehst in die 11te Klasse")
	} else if age == 18 {
		fmt.Println("Du gehst in die 12te Klasse")
	} else if age > 18 {
		fmt.Println("Du gehst nicht mehr zur Schule")
	}
}

// switch case
func PrintClass2(age int) {
	switch age {
	case 10:
		fmt.Println("Du gehst in die 4 Klasse")
	case 11:
		fmt.Println("Du gehst in die 5 Klasse")
	case 12:
		fmt.Println("Du gehst in die 6 Klasse")
	case 13:
		fmt.Println("Du gehst in die 7 Klasse")
	case 14:
		fmt.Println("Du gehst in die 8 Klasse")
	case 15:
		fmt.Println("Du gehst in die 9 Klasse")
	case 16:
		fmt.Println("Du gehst in die 10 Klasse")
	case 17:
		fmt.Println("Du gehst in die 11 Klasse")
	case 18:
		fmt.Println("Du gehst in die 12 Klasse")
	default:
		fmt.Println("Du gehst nicht mehr zur Schule")
	}
}

func CompareStrings(given, expected string) bool {
	return given == expected
}
