package main

import "fmt"

func main() {
	CheckName2("Sabina")
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
