package main

import "fmt"

//全局变量定义，必须有var关键字
var aa = 3
var cc = 4
var bb = true

//cc := 123, 函数外面定义不能使用
//变量定义的作用域是在函数里面
// 定义变量

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)

}

// 变量初始值
func variableInitialVale() {
	var a int = 3
	var x, z, e int = 3, 4, 7 //赋值多个
	var s string = "abc"
	fmt.Println(a, s, x, z, e)
}

//type判断
func variableTypeDeduction() {
	var a, b, c = 1, "abc", 0.1
	fmt.Println(a, b, c)
}

//简单的写法
func variableShorter() {
	a, b, c := 1, true, 0.1 // :=定义一个变量
	b = false               //定义后只能定义相同类型的，不能重复定义
	fmt.Println(a, b, c)
}

//匿名变量
func GetData() (int, int) {
	return 100, 200
}

func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInitialVale()
	variableTypeDeduction()
	variableShorter()
	cc, aa = aa, cc
	a, _ := GetData()
	_, b := GetData()
	fmt.Println(a, b)
	fmt.Println(aa, cc, bb)
}
