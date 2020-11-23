package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	//variable()
	variableInit()
}

func variableInit() {
	// 编译器推断类型
	var a, b = 3, 4
	c, d := true, "abc"
	println(a, b)
	println(c, d)
}
func variable() {
	var i int
	var s string
	fmt.Printf("%d %q\n", i, s)
}
