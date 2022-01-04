package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var arr = [8]int{1, 3, 4, 5, 6, 8, 9, 7}

	fmt.Println("Array Elements Before Slice:", arr)

	b := arr[1:6]
	fmt.Println("Array Elements After Slice:", b)

	b = append(b, rand.Intn(100))
	fmt.Println("Appending rand-number to slice", b)

}
