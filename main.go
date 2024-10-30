package main

import (
	"fmt"
	"maps"
)

func main() {
	// Declare a map with a key of type string and a value of type int
	inventory := map[string]int{
		"pens":      100,
		"notebooks": 50,
	}
	// Alternative way to initialize a map
	// inventory := map[string]int{}
	// inventory["pens"] = 100
	// inventory["notebooks"] = 50
	fmt.Println(inventory)

	// Accessing a value from a map
	quantity := inventory["pens"]
	fmt.Println(quantity) // Output will be 100

	// Accessing a value from a map that does not exist
	unknownQuantity := inventory["markers"]
	fmt.Println(unknownQuantity) // Output will be 0, since "markers" does not exist

	// Checking if a key exists in a map
	v, ok := inventory["paper clips"]
	if ok {
		fmt.Println("The inventory has", v, "paper clips")
	} else {
		fmt.Println("Paper clips are not in the inventory")
	}

	// Deleting a key from a map
	delete(inventory, "notebooks")
	_, ok = inventory["notebooks"]
	fmt.Println(ok) // Output will be false

	// Iterating over a map

	inventory = map[string]int{
		"pens":      100,
		"notebooks": 50,
		"erasers":   30,
	}

	for item, quantity := range inventory {
		fmt.Printf("There are %d %s in the inventory\n", quantity, item)
	}

	// Output:
	// There are 100 pens in the inventory
	// There are 50 notebooks in the inventory
	// There are 30 erasers in the inventory

	// Comparing maps
	if inventory == nil {
		fmt.Println("The inventory map is not initialized")
	}

	inventory1 := map[string]int{
		"pens":      100,
		"notebooks": 50,
		"erasers":   30,
	}

	inventory2 := map[string]int{
		"pens":      100,
		"notebooks": 50,
		"erasers":   30,
	}

	if maps.Equal(inventory1, inventory2) {
		fmt.Println("The inventories are equal.")
	} else {
		fmt.Println("The inventories are not equal.")
	}

	// Clearing a map
	fmt.Println(inventory1) // Output will be map[erasers:30 notebooks:50 pens:100]
	clear(inventory1)
	fmt.Println(inventory1) // Output will be map[]

}
