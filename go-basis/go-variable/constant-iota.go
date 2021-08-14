package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)

	const (
		l = 1 << iota //<< 表示左移的意思
		j = 3 << iota
		k
		o
	)
	fmt.Println("l=", l)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("o=", o)

}
