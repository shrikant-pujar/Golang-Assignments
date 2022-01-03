package main

import "fmt"

func reverse(s string) string {
	rev := []rune(s)
	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
	return string(rev)
}

func main() {

	str := "Quest-Global"
	strRev := reverse(str)
	fmt.Println(str)
	fmt.Println(strRev)
}