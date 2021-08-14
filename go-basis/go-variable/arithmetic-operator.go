package main

import "fmt"

func main() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行c的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行c的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行c的值为 %d\n", c)
}
