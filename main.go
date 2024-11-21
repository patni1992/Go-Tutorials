package main

import (
	"encoding/json"
	"fmt"
)

type Printer struct{}

func (p Printer) Print() {
	fmt.Println("Printing document...")
}

type Scanner struct{}

func (s Scanner) Scan() {
	fmt.Println("Scanning document...")
}

// A struct that combines both Printer and Scanner
type MultiFunctionDevice struct {
	Printer
	Scanner
}

type Library struct {
	Name  string
	Books []Book
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

// Method
func (b Book) Summary() string {
	return fmt.Sprintf("%s by %s, %d pages", b.Title, b.Author, b.Pages)
}

func (b *Book) UpdatePages(newPages int) {
	if newPages > 0 {
		b.Pages = newPages
	} else {
		fmt.Println("Invalid page count. Pages must be greater than zero.")
	}
}

func main() {

	// Declare a struct

	myBook := Book{
		Title:  "The Go Programming Language",
		Author: "Alan A. A. Donovan",
		Pages:  380,
	}

	// Another way to declare a struct

	// var myBook Book
	// myBook.Title = "The Go Programming Language"
	// myBook.Author = "Alan A. A. Donovan"
	// myBook.Pages = 380

	fmt.Println(myBook) // {The Go Programming Language Alan A. A. Donovan 380}

	// Accessing a struct field
	title := myBook.Title
	fmt.Println(title) // The Go Programming Language

	// Modifying a struct field
	myBook.Title = "New Title"
	fmt.Println(myBook.Title) // New Title

	// Nested structs
	myLibrary := Library{
		Name: "City Library",
		Books: []Book{
			{Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Pages: 380},
			{Title: "Go in Action", Author: "William Kennedy", Pages: 300},
		},
	}

	fmt.Printf("%+v", myLibrary) // {Name:City Library Books:[{Title:The Go Programming Language Author:Alan A. A. Donovan Pages:380} {Title:Go in Action Author:William Kennedy Pages:300}]}

	// Using a method
	fmt.Println(myBook.Summary()) // The Go Programming Language by Alan A. A. Donovan, 380 pages

	// Using a method with a pointer receiver
	myBook.UpdatePages(400)
	fmt.Println(myBook.Pages) // 400

	// Comparing structs

	book1 := Book{Title: "Go in Action", Author: "William Kennedy", Pages: 300}
	book2 := Book{Title: "Go in Action", Author: "William Kennedy", Pages: 300}

	if book1 == book2 {
		fmt.Println("The books are identical.")
	} else {
		fmt.Println("The books are different.")
	}

	// Composition
	mfd := MultiFunctionDevice{}
	mfd.Scan()
	mfd.Print()

	// Anonymous structs

	jsonResponse := `{"Name": "Alice", "Age": 30}`

	// Use an anonymous struct directly
	person := struct {
		Name string
		Age  int
	}{}

	err := json.Unmarshal([]byte(jsonResponse), &person)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)

}
