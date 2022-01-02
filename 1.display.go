package main

import "fmt"

func main() {
	n := 5
	for i := 0; i < 5; i++ {
		for j := n; j > 0; j-- {
			fmt.Printf("*")
		}

		fmt.Printf("\n")
		n--
	}
}