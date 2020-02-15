package main

import (
	"TestGo/pk1"
	"TestGo/utils"
	"TestGo/utils/timeutils"
	"fmt"
	"math/rand"
	"time"
)

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

	//var age1 int = 12
	//var x1 int
	//fmt.Println("请输入:")
	//fmt.Scanln(&x1)
	//fmt.Println(x1, age1)

	for i := 1; i <= 5; i++ {
		fmt.Println("hello world")
	}

	t1 := time.Now()
	z := time.RFC1123Z
	fmt.Println(t1)
	fmt.Println(z)

	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 10; i++ {
		fmt.Println(rand.Intn(100))
	}
	aaa := []int{1, 2, 3}
	fmt.Println(aaa)

	// 包调用
	utils.Count()
	timeutils.PringTime()

	pk1.MyTest()
	utils.MyTest2()

}
