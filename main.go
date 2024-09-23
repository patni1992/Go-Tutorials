package main

import (
	"errors"
	"fmt"
)

// Simple function to say hello
func sayHello() {
	fmt.Println("Hello, Go!")
}

// Function with parameters and return value
// func add(a int, b int) int {
// 	return a + b
// }

// multiple parameters same type
func add(a, b int) int {
	return a + b
}

// Function with multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// Function as parameter
func apply(slice []int, fn func(int) int) {
	for i, v := range slice {
		slice[i] = fn(v)
	}
}

func double(n int) int {
	return n * 2
}

func square(n int) int {
	return n * n
}

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

// counter is a function that returns another function, which has a closure
func counter() func() int {
	x := 0 // x is captured by the returned function
	return func() int {
		x++ // Each call to this function will increment x
		return x
	}
}

// Function with named return value
// func add(a, b int) (sum int) {
// 	sum = a + b
// 	return
// }

func complicatedFunction(a, b int) (result int, err error) {
	if a < 0 {
		err = errors.New("a is negative")
		return
	}
	result = a + b
	if result > 100 {
		err = errors.New("too large")
		return
	}
	// more logic
	return
}

func main() {
	// Call sayHello function 3 times
	sayHello()
	sayHello()
	sayHello()

	// Call add function with parameters 2 and 5
	fmt.Println(add(2, 5))

	// Call divide function with parameters 10 and 0
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Call apply function with parameters nums and double, square
	nums := []int{1, 2, 3, 4, 5}
	apply(nums, double)
	fmt.Println(nums)
	nums = []int{1, 2, 3, 4, 5}
	apply(nums, square)
	fmt.Println(nums)

	// Call sum function with parameters 100, 15, 20, 5, 7 & with slice of numbers
	sum(100, 15, 20, 5, 7)
	numbers := []int{5, 8, 10, 20}
	sum(numbers...)

	// Anonymous function
	func() { fmt.Println("Hello!") }()
	// anonymous function assigned to variable
	greet := func(name string) { fmt.Println("Hello,", name) }
	greet("Go")

	// Call counter function (closure)
	increment := counter()

	fmt.Println(increment()) // Output: 1
	fmt.Println(increment()) // Output: 2
	fmt.Println(increment()) // Output: 3
}
