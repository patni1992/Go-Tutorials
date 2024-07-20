package main

import "fmt"

func main() {

	// Printing
	fmt.Print("Hello!")
	fmt.Print("Go is awesome")

	fmt.Print("Hello!\n")
	fmt.Print("Go is awesome")

	fmt.Print("My name is", "Patrik")

	name := "Patrik"
	lang := "Go"

	fmt.Println("My name is", name)
	fmt.Println("I love the" + lang + " programming language")

	// Formatting

	countBurgers := 2

	fmt.Printf("My name is %s and I had %d burgers today!", name, countBurgers)

	sentence := fmt.Sprintf("My name is %s and I had %d burgers today!", name, age)

	fmt.Println(sentence)

	// Reading user input

	var age int

	fmt.Println("Enter your name: ")
	fmt.Scan(&name)

	fmt.Println("Enter your age: ")
	fmt.Scan(&age)

	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)

	var favColor1 string
	var favColor2 string

	fmt.Println("Enter your two favorite colors: ")
	fmt.Scan(&favColor1, &favColor2)
	fmt.Println("favorite colors ", favColor1, favColor2)

	fmt.Println("Enter your two favorite colors: ")
	fmt.Scanln(&favColor1, &favColor2)
	fmt.Println("favorite colors ", favColor1, favColor2)

	fmt.Println("Enter your two favorite colors separated by a comma (e.g., red,blue): ")
	fmt.Scanf("%s,%s", &favColor1, &favColor2)
	fmt.Printf("Your favorite colors are: %s and %s.\n", favColor1, favColor2)

}
