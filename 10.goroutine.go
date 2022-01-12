package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Start Goroutines")

	go func() {

		defer wg.Done()

		for count := 0; count < 1; count++ {
			for char := 0; char <= 100; char++ {
				fmt.Printf("%d ", char)
			}
		}
		fmt.Println()
		fmt.Println()
	}()

	go func() {

		defer wg.Done()

		for count := 0; count < 1; count++ {
			for char := 0; char <= 100; char++ {
				fmt.Printf("%d ", char)
			}
		}
		fmt.Println()
		fmt.Println()
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nEnd of the Program")
	fmt.Println()
}
