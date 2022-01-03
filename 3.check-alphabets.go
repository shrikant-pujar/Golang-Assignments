package main

import "fmt"

func main() {

	str := "i am @ quest-global"
	vc := 0		//vowel count
	cc := 0		//consonent count
	sc := 0		//special char count
	for _, ch := range str {

		switch ch {
		case 'a', 'e', 'i', 'o', 'u':
			vc++

		case 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z':
			cc++

		case '-', '@', '$', ' ': //i have considered blank-space as special charecter
			sc++
		}
	}
	fmt.Printf("vowel count = %v\n", vc)
	fmt.Printf("consonent count=%v\n", cc)
	fmt.Printf("special charecter=%v\n", sc)
}
