package main

import "fmt"

func main() {
	const PATH string = "http://www.baidu.com"
	const PI = 3.14
	fmt.Println(PATH)
	fmt.Println(PI)

	const (
		MALE       = 0
		FEMALE     = 1
		UNKNOW     = 3
		a      int = 100
		b
	)
	fmt.Println(a, b)

	// 枚举类型
	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINNER = 3
	)
	// iota:可被编译器修改的常量,每定义一个常量自动加一,直到下一个const出现
	const (
		a1 = iota
		a2
		a3
	)
	fmt.Println(a3)

}
