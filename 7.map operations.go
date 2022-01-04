package main

import "fmt"

func main() {
	fruits := map[int]string{
		0: "orange",
		1: "apple",
		2: "banana",
		3: "mango",
		4: "papaya",
	}
	fmt.Println("Content of Maps are:\n", fruits)

	// Adding new key-value pairs in the map
	fruits[5] = "Grapes"
	fruits[6] = "watermilon"
	fmt.Println("\n Map after adding new key-value pair:\n", fruits)

	// Updating values of the map
	fruits[1] = "kiwi"
	fruits[3] = "pear"
	fmt.Println("\n Map after updating values of the map:\n", fruits)

	//delete values of the map
	delete(fruits, 2)
	delete(fruits, 4)
	fmt.Println("\n Map after deleting values of the map:\n", fruits)

	//delete all elements of map
	for item := range fruits {
		delete(fruits, item)
	}
	fmt.Println("\n Map after deleting all elements of map:\n", fruits)
}
