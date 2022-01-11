package main

import "fmt"

func main() {

	//Creating array with repeated number
	array := []int{5, 8, 7, 6, 7, 9, 5, 7, 8, 2, 10}
	check(array)

	fmt.Println()
}

func check(arr []int) {
	//Create a map to store number of occurences of number
	newmap := make(map[int]int)

	for _, num := range arr {

		newmap[num] = newmap[num] + 1

	}

	fmt.Println(newmap)
}
