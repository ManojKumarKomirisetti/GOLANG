package main

import "fmt"

func main() {
	var x int = 13
	y := 153.0236
	var z float32 = 89
	var w uint = 53
	var k int = 25
	var name string = "Manoj Kumar"

	fmt.Println("Values:", x, y, z, w, k, name)
	fmt.Println("%T", z)
	fmt.Printf("%T", k)
	fmt.Println()
	//fmt.Printf("DataTypes:", "%T", "%T", "%T", "%T", "%T", "%T", x, y, z, w, k, name)
	fmt.Println()
	fmt.Printf("Values are: %T, %T, %T", x, y, z)
}
