package main

import "fmt"

var x = 100

func main() {
	fmt.Println(&x, x)

	x := "abc"
	fmt.Println(&x, x)
}
