package main

import "fmt"

func main() {
	fmt.Println("Main Method")
	manoj()
	manoj1()
}
func manoj1() {
	var Value int
	var Value2 int32
	var Value3 int64
	var Value4 float32
	var Value5 float64
	fmt.Println("Default Values:", Value, Value2, Value3, Value4, Value5)
}
func manoj() {
	fmt.Println("Manoj Function")
	var name string = "First Variable"
	fmt.Println("The Variable is:", name)
}
