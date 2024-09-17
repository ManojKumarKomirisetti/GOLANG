package main

import (
	"fmt"
)

func main() {
	const PI = 3.14159265358979323846264338327950288419716939937510582097494459
	fmt.Println(PI) //3.141592653589793
	/*In the above example, we declared a variable named PI and`
	assigned the mathematically correct value to it. The Go compiler
	will loose some precision values, and the output will be the closest float64 value.
	Up to 15 Values.*/
}
