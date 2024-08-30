package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// if statement
	goIsAwesome := true
	if goIsAwesome {
		fmt.Println("I love Go")
	}

	// if, else if, else statement
	language := "Go"
	if language == "Go" {
		fmt.Println("Keep rocking")
	} else if language == "Python" {
		fmt.Println("Solid choice")
	} else {
		fmt.Println("Exploring new languages?")
	}

	// variable declaration inside if statement
	foods := []string{"Pizza", "Sushi", "Burger", "Tacos"}
	if choice := foods[rand.Intn(len(foods))]; choice == "Pizza" {
		fmt.Println("Pizza is delicious! Great choice.")
	} else if choice == "Sushi" {
		fmt.Println("Sushi is a healthy option. Nice!")
	} else if choice == "Burger" {
		fmt.Println("Burger is always a good comfort food.")
	} else {
		fmt.Println("You like", choice, "which is interesting!")
	}

	// switch statement
	day := 6
	switch day {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("Invalid day")
	}

	fruit := "banana"

	// switch statement with multiple cases
	switch fruit {
	case "apple", "cherry":
		fmt.Println("Red")
	case "banana", "lemon":
		fmt.Println("Yellow")
	case "grape", "plum":
		fmt.Println("Purple")
	default:
		fmt.Println("Unknown color")
	}

	temperature := 15

	// blank switch statement
	switch {
	case temperature < 5:
		fmt.Println("Winter")
	case temperature >= 5 && temperature < 15:
		fmt.Println("Autumn")
	case temperature >= 15 && temperature < 25:
		fmt.Println("Spring")
	case temperature >= 25:
		fmt.Println("Summer")
	default:
		fmt.Println("Unusual weather")
	}
}
