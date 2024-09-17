package main

import (
	"fmt"
)

func short() {
	fmt.Println("Short Method First Statement!")
	y := 13
	fmt.Println(y)
	y = 15
	fmt.Println(y)
	n := "Manoj"
	fmt.Println(n)

	n = "Manoj Kumar"
	fmt.Println(n)
	var m string = "Tectoro"
	fmt.Println(m)
	m = "TCPL"
	fmt.Println(m)

}
func main() {
	short()
	fmt.Println("Main Method Last Statement!")
}

//ShortHand Notation cannot be written  outside of scope.
