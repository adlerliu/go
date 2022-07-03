package main

import "fmt"

func main() {
	x := 100
	println(&x)
	x, y := 200, "abc"
	fmt.Println(&x, x)
	fmt.Println(y)
}
