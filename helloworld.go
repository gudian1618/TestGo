package main

import "fmt"

func main() {
	/*
		变量:一小块内容,可以改变
		使用:
		1.变量的声明:定义
		2.访问:赋值与取值
	*/

	// 原样
	var num1 int = 12
	fmt.Println(num1)

	// 推断类型
	var num2 = "fads"
	fmt.Println(num2)

	// 简短类型,:=只能定义不能改变赋值,
	// 除非多变量中有为定义变量可修改其余赋值
	// 简短定义方式不能定义全局变量
	sum := 100
	fmt.Println(sum)

	// 多个变量同时定义
	var m, n = 100, "ser"
	fmt.Println(m, n)

	// 批量定义一组值
	var (
		name = "sfas"
		age  = 18
	)
	fmt.Println(name, age, &name)

}
