package main

import "fmt"

func main() {
	const (
		Unkown = 0
		Female = 1
		Male   = 2
	)
	const (
		a = iota
		b
		c
		d
		e
	)
	fmt.Println(Unkown, Female, Male, Male)
	fmt.Println(a, b, c, d, e)
}
