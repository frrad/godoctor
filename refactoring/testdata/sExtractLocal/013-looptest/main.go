package main

import "fmt"

func main() {
	x := 5
	y := 1
	z := 0
	for x = 1; x < y; x++ {
		z = x + y // <<<<< extractLocal,10,6,10,11,newVar,pass
	}
	fmt.Println(z)
	fmt.Println("x", x, "y", y)
}
