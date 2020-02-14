package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var age1 int = 12
	var x1 int
	fmt.Println("请输入:")
	fmt.Scanln(&x1)
	fmt.Println(x1, age1)

	for i := 1; i <= 5; i++ {
		fmt.Println("hello world")
	}

	num1 := rand.Int()
	fmt.Println(num1)
	for i := 0; i < 10; i++ {
		// 种子数固定某一次生成的随机数
		//rand.Seed(1000)
		num := rand.Intn(10)
		fmt.Println(num)
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

}
