package main

import "fmt"

func main() {
	fmt.Println("The Complex Datatypes:")
	var c complex64 = 30 + 15i
	var c1 complex128 = 55 + 25i
	fmt.Println("Complex32:", c)
	fmt.Println("Complex64:", c1)
}
