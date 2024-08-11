package main

import (
	"fmt"
	"slices"
)

func main() {

	// Arrays

	// Declare an array
	var zeroArray [5]int
	var myArray = [3]int{5, 10, 15}
	var cars = [5]string{1: "Volvo", 4: "BMW"}
	var snacks = [...]string{"chips", "popcorn", "peanuts"}
	fmt.Println(zeroArray) // [0 0 0 0 0]
	fmt.Println(myArray)   // [5 10 15]
	fmt.Println(cars)      // [ "" Volvo "" "" BMW]
	fmt.Println(snacks)    // [chips popcorn peanuts]

	// Get the length of an array
	fmt.Println(len(snacks))

	// Compare arrays
	x := [3]int{1, 2, 3}
	y := [3]int{1, 2, 3}

	fmt.Println(x == y) // true
	fmt.Println(x != y) // false

	// Update element of an array
	x[0] = 10
	fmt.Println(x[0]) // 10
	fmt.Println(x)    // 10, 2, 3

	// Slices
	// Declare a slice
	someSlice := []int{5, 10, 15}
	var carsSlice = []string{1: "Volvo", 4: "BMW"}
	var sliceWithNoValue []int

	fmt.Println(someSlice)        // [5 10 15]
	fmt.Println(carsSlice)        // "", Volvo, "", "", BMW
	fmt.Println(sliceWithNoValue) // []

	// Update element of a slice
	someSlice[2] = 100
	fmt.Println(someSlice[2]) // 100

	// Compare slices
	fmt.Println(sliceWithNoValue == nil) // true
	fmt.Println(sliceWithNoValue != nil) // false

	fruits1 := []string{"apple", "banana", "peach"}
	fruits2 := []string{"apple", "banana", "peach"}

	fmt.Println(slices.Equal(fruits1, fruits2)) // true

	// Get the length of a slice
	fmt.Println(len(fruits1))          // 3
	fmt.Println(len(sliceWithNoValue)) // 0

	// Append to a slice
	fruits1 = append(fruits1, "coconut")
	fmt.Println(fruits1) // apple, banana, peach, coconut

	// Append multiple elements to a slice
	fruits1 = append(fruits1, "strawberry", "raspberry")
	fmt.Println(fruits1) // apple, banana, peach, coconut, strawberry, raspberry

	// Append a slice to a slice
	var fruits3 = []string{"durian", "dragonfruit"}
	fruits1 = append(fruits1, fruits3...)
	fmt.Println(fruits1) // apple, banana, peach, coconut, strawberry, raspberry, durian, dragonfruit

	// How slices grow
	var nums []int
	fmt.Println(nums, len(nums), cap(nums)) // [] 0 0
	nums = append(nums, 1)
	fmt.Println(nums, len(nums), cap(nums)) // [1] 1 1
	nums = append(nums, 2)
	fmt.Println(nums, len(nums), cap(nums)) // [1 2] 2 2
	nums = append(nums, 3)
	fmt.Println(nums, len(nums), cap(nums)) // [1 2 3] 3 4
	nums = append(nums, 4)
	fmt.Println(nums, len(nums), cap(nums)) // [1 2 3 4] 4 4
	nums = append(nums, 5)
	fmt.Println(nums, len(nums), cap(nums)) // [1 2 3 4 5] 5 8

	// Make a slice with a specific length and capacity
	nums2 := make([]int, 0, 20)                // Set the length to 0 & capacity to 20
	nums2 = append(nums2, 5, 10, 15, 20)       // nums2 will have a length of 4 & capacity of 20
	fmt.Println(nums2, len(nums2), cap(nums2)) // [5 10 15 20] 4 20

	// Slice a slice
	nums3 := []int{2, 4, 8, 16, 32}
	fmt.Println(nums3[1:4]) // 4, 8, 16
	fmt.Println(nums3[:4])  // 2, 4, 8, 16
	fmt.Println(nums3[1:])  // 4, 8, 16, 32
	fmt.Println(nums3[:])   // 2, 4, 8, 16, 32

	// Be aware when updating a slice, it will update the original slice as well
	nums4 := nums3[1:4]
	nums4[0] = 1337
	fmt.Println(nums3) // 2, 1337, 8, 16, 32
	fmt.Println(nums4) // 1337, 8, 16

	// Remove an element from a slice
	s := []int{10, 20, 30, 40}
	s = append(s[:2], s[2+1:]...)
	fmt.Println(s) // 10, 20, 40

	// Copy a slice
	values := []int{5, 6, 7, 8}
	newValues := make([]int, 4)
	copy(newValues, values)
	newValues[0] = 1337
	fmt.Println(newValues) // 1337, 6 ,7 ,8
	fmt.Println(values)    // 5, 6 ,7 ,8

}
