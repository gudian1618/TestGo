package utils

import "fmt"

func Count() {
	fmt.Println("调用计数函数")
}

func init() {
	fmt.Println("utils包的另一个初始换函数")
}

func init() {
	fmt.Println("utils包的初始化函数")
}
