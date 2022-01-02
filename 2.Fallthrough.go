package main

import "fmt"

func main() {
    name := "raj"

    switch {
    case name == "manoj":
        fmt.Println("Manoj")
        fallthrough
    case name == "raj":
        fmt.Println("raj")
        fallthrough
    case name == "hello":
        fmt.Println("hello")
    }
}