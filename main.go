package main

import "fmt"

func main() {
	var animal1 string = "cat"
	var animal2 = "dog"
	animal3 := "monkey"
	var animal4 string
	animal4 = "lion"

	var defaultNum int
	var defaultBool bool
	var defaultStr string

	fmt.Println(animal1, animal2, animal3, animal4)
	fmt.Println(defaultNum, defaultBool, defaultStr)

	var (
		name    string = "John"
		age     int    = 30
		address string = "123 Go Street"
	)

	fmt.Println(name, age, address)

	const pi = 3.14
	const port = 1337

	fmt.Println(pi, port)

	var num1 int = 1992
	var num2 int64 = 123456789
	var num3 int8 = 127
	var num4 uint = 100

	fmt.Println(num1, num2, num3, num4)

	var amount1 float32 = 99.84
	var amount2 float64 = 109.5
	amount3 := 56.7433

	fmt.Println(amount1, amount2, amount3)

	var isTrue bool = true
	isFalse := false

	fmt.Println(isTrue, isFalse)

}
