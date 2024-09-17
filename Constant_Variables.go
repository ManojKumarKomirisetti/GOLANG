package main

import (
	"fmt"
)

func main() {
	const name string = "Manoj Kumar"
	fmt.Println("Name of the Student", name)
	// Implicit Typing
	const id = 13.2563
	fmt.Println("Value:", id)
	fmt.Printf("Type:%T", id)
	fmt.Println()
	const fid = 2 / 3
	fmt.Println("Value:", fid) //0
	fmt.Printf("Type:%T", fid)
	const fid1 = 2.0 / 3.0
	fmt.Println()
	fmt.Println("Value:", fid1) //0.66666666666666666
	fmt.Printf("Type:%T", fid1)
	const fid2 float64 = 2 / 3
	fmt.Println()
	fmt.Println("Value:", fid2) //0
	fmt.Printf("Type:%T", fid2)
}
