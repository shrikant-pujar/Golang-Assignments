package main

import "fmt"

func main() {
	var arr = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println("Accessing 2D array elements without range")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Println("Accessing 2D array elements with range")
	for _, item := range arr {
		fmt.Println(item)
	}
}
