package main

import "fmt"

func main() {
	// three component for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Skiping initialization in the header
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	// Skipping the increment in the header
	for i := 0; i < 10; {
		fmt.Println(i)
		if i < 5 {
			i += 2
		} else {
			i += 3
		}
	}

	// while loop
	number := 1
	for number <= 10 {
		fmt.Println(number)
		number++
	}

	// infinite loop
	// for {
	// 	fmt.Println("This loop will run forever.")
	// }

	// break statement
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break // Stops the loop when i reaches 5
		}
		fmt.Println(i)
	}

	// continue statement
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // Skips even numbers
		}
		fmt.Println(i)
	}

	// range loop

	languages := []string{"Go", "Java", "C", "Python", "Ruby"}

	for i, lang := range languages {
		fmt.Println(i)
		fmt.Println(lang)
	}

	// range loop skipping index
	for _, lang := range languages {
		fmt.Println(lang)
	}
}
