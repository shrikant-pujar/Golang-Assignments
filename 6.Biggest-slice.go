package main

import "fmt"

func main() {
	var Biggest, temp int
	array := []int{10, 19, 12, 13, 41}
	slice := array[1:4]

	for _, element := range slice {
		if element > temp {
			temp = element
			Biggest = temp
		}
	}
	fmt.Println("Biggest number of slice is: ", Biggest)
}
