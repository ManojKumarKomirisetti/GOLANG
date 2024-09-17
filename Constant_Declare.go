package main

import (
	"fmt"
)

func declare() {
	const a, b, c, d = 12, 15.5, 10 + 35i, "Manoj Kumar"
	fmt.Println("Values of ABCD:", a, b, c, d)
	fmt.Printf("Type:%T,%T,%T,%T", a, b, c, d)
	fmt.Println("Next Line:-----------")
	const p, q, r, s float64 = 12, 16, 25, 63.32
	fmt.Printf("Type:%T,%T,%T,%T", p, q, r, s)
	fmt.Println("Values of WXYZ:", p, q, r, s)
	fmt.Println("Next Line:-----------")
	const w, x, y, z complex128 = 12, 16, 25, 63.32
	fmt.Println("Values of WXYZ:", w, x, y, z)
	fmt.Printf("Type:%T,%T,%T,%T", w, x, y, z)
}

func main() {
	declare()
}
