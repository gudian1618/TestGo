package main

import (
	_ "TestGo/pk1"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	//const PATH string = "http://www.baidu.com"
	//const PI = 3.14
	//fmt.Println(PATH)
	//fmt.Println(PI)
	//
	//const (
	//	MALE       = 0
	//	FEMALE     = 1
	//	UNKNOW     = 3
	//	a      int = 100
	//	b
	//)
	//fmt.Println(a, b)
	//
	//// 枚举类型
	//const (
	//	SPRING = 0
	//	SUMMER = 1
	//	AUTUMN = 2
	//	WINNER = 3
	//)
	//// iota:可被编译器修改的常量,每定义一个常量自动加一,直到下一个const出现
	//const (
	//	a1 = iota
	//	a2
	//	a3
	//)
	//fmt.Println(a3)
	///*
	//	 	变量:一小块内容,可以改变
	//		使用:
	//		1.变量的声明:定义
	//		2.访问:赋值与取值
	//*/
	//
	//// 原样
	//var num1 int = 12
	//fmt.Println(num1)
	//
	//// 推断类型
	//var num2 = "fads"
	//fmt.Println(num2)
	//
	//// 简短类型,:=只能定义不能改变赋值,
	//// 除非多变量中有为定义变量可修改其余赋值
	//// 简短定义方式不能定义全局变量
	//sum := 100
	//fmt.Println(sum)
	//
	//// 多个变量同时定义
	//var m, n = 100, "ser"
	//fmt.Println(m, n)
	//
	//// 批量定义一组值
	//var (
	//	name = "sfas"
	//	age  = 18
	//)
	//fmt.Println(name, age, &name)
	//
	////var age1 int = 12
	////var x1 int
	////fmt.Println("请输入:")
	////fmt.Scanln(&x1)
	////fmt.Println(x1, age1)
	//
	//for i := 1; i <= 5; i++ {
	//	fmt.Println("hello world")
	//}
	//
	//t1 := time.Now()
	//z := time.RFC1123Z
	//fmt.Println(t1)
	//fmt.Println(z)
	//
	//rand.Seed(time.Now().UnixNano())
	//for i := 1; i <= 10; i++ {
	//	fmt.Println(rand.Intn(100))
	//}
	//aaa := []int{1, 2, 3}
	//fmt.Println(aaa)

	// 包调用

	/*
		init和mian函数
		1.go内置函数,一个初始化一次可以多个;一个程序的入口
		2.无参无返回值,go自动调用,无法手动调用
		3.init可以在不同包下多个,main只能有一个
		4.先执行init,后执行mani,同一个go文件顺序执行
		不同包下,将文件名按字符串进行排序后调用init函数
		如果多个包之间存在依赖的话
		导入顺序与调用初始化顺序相反,为栈结构
		5._操作,表示导入包只执行初始化,不调用该包中任何函数
	*/

	//utils.Count()
	//timeutils.PringTime()
	//
	//pk1.MyTest1()
	//utils.MyTest2()
	//
	//fmt.Println("-------------------")
	//
	//p1 := person.Person{"王二狗", 30, "男"}
	//fmt.Println(p1.Name, p1.Age, p1.Sex)
	//
	//db, err := sql.Open("mysql", "root:test@tcp(127.0.0."+
	//	"1:3306)/db20?charset=utf8")
	//if err != nil {
	//	fmt.Println("错误信息:", db)
	//	return
	//}
	//fmt.Println("连接成功:", db)

	//pk2.MyTest3()

	t1 := time.Now()
	fmt.Println("%T\n", t1)
	fmt.Println(t1)

	time.Date(2020, 2, 15, 23, 23, 23, 23, time.Local)
	s1 := t1.Format("2006年1月2日 15:04:05")
	fmt.Println(s1)

	s2 := t1.Format("2006/01/02")
	fmt.Println(s2)
	fmt.Printf("%T\n", t1)

	year, month, day := t1.Date()
	fmt.Println(year, month, day)

	hour, min, sec := t1.Clock()
	fmt.Println(hour, min, sec)

	yearDay := t1.YearDay()
	fmt.Println(yearDay)
	fmt.Println(t1.Weekday(), t1.Day())

	// 时间戳 距离UTC时的时间差
	fmt.Println(t1.Unix(), t1.UnixNano())

	// 时间间隔
	fmt.Println(t1.Add(time.Minute))

	// 睡眠延迟
	//time.Sleep(1 * time.Second)
	//fmt.Println("****")
	//
	//go printNum()
	//for i := 1; i <= 50; i++ {
	//	fmt.Printf("\t主程序打印字母: A %d\n", i)
	//}
	//time.Sleep(time.Second)

	// 获取goroot目录
	fmt.Println("goroot-->", runtime.GOROOT())
	// 获取操作系统
	fmt.Println("os/platform-->", runtime.GOOS)
	// 获取逻辑cpu数量
	fmt.Println("逻辑cpu的数量:", runtime.NumCPU())

	// gosched
	//go func() {
	//	for i:=1; i<12; i++ {
	//		fmt.Println("goroutine...")
	//	}
	//}()
	//
	//for i:=0; i<4; i++ {
	//	// 让出时间片,先让别的goroutine执行
	//	runtime.Gosched()
	//	fmt.Println("main...")
	//}

	//go func() {
	//	fmt.Println("goroutine开始...")
	//	fun()
	//	fmt.Println("goroutine结束...")
	//}()
	//time.Sleep(3 * time.Second)

	/*
		4个模拟并发
	*/

	//wg.Add(4)
	//
	//go saleTickets("售票口1")
	//go saleTickets("售票口2")
	//go saleTickets("售票口3")
	//go saleTickets("售票口4")
	//
	//wg.Wait()
	//fmt.Println("main结束了")

	//time.Sleep(7*time.Second)

	//wg.Add(2)
	//go fun1()
	//go fun2()
	//
	//fmt.Println("main进入阻塞状态,等待wg中的goroutine结束")
	//wg.Wait()
	//fmt.Println("main结束")

	//ruMutex = new(sync.RWMutex)
	//wg1 = new(sync.WaitGroup)
	//
	//wg1.Add(4)
	//// 多个读同时进行
	//go readData(0)
	//go writeData(1)
	//go readData(2)
	//go writeData(3)
	//
	//wg1.Wait()
	//fmt.Println("main结束...")

	/*
		channel通道
	*/

	//var a chan int // 声明通道变量
	//fmt.Printf("%T,%v\n", a, a)
	//
	//if a == nil {
	//	fmt.Println("为空不能使用")
	//	a = make(chan int)
	//	fmt.Println(a)
	//}
	//test1(a)
	//
	//var ch1 chan bool
	//ch1 = make(chan bool)
	//
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println("子协程中,i:", i)
	//	}
	//	// 向通道内写数据,发送数据到通道
	//	ch1 <- true
	//	fmt.Println("结束...")
	//}()
	//// 通道向外发数据,向通道外读数据
	//data := <- ch1
	//fmt.Println("main...", data)
	//fmt.Println("main结束...")

	//ch2 := make(chan int)
	//
	//go func() {
	//	fmt.Println("子程序开始执行...")
	//	time.Sleep(2 * time.Second)
	//	data := <-ch2
	//	fmt.Println("data:", data)
	//}()
	//ch2 <- 10
	//fmt.Println("main.over..")

	// 死锁
	//ch3 := make(chan int)
	//ch3 <- 100 // 阻塞

	/*
		关闭通道
		子goroutine和住goroutine的读写阻塞与接触阻塞
	*/

	//ch4 := make(chan int)
	//go sendData(ch4)
	//// 读取通道的数据
	//for {
	//	time.Sleep(1 * time.Second)
	//	v, ok := <-ch4
	//	if !ok {
	//		fmt.Println("读取了所有的值..", ok)
	//		break
	//	}
	//	fmt.Println("读取的数据:", v, ok)
	//}
	//fmt.Println("main...over...")

	ch5 := make(chan int)
	go sendData(ch5)
	for v := range ch5 {
		fmt.Println("读取数据:", v)
	}
	fmt.Println("main...over")

}

func sendData(ch4 chan int) {
	// 发送方: 发送数据
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		ch4 <- i
	}
	close(ch4)
}

func test1(ch chan int) {
	fmt.Printf("%T,%v\n", ch, ch)
}

var wg1 *sync.WaitGroup
var ruMutex *sync.RWMutex

func writeData(i int) {
	defer wg1.Done()

	fmt.Println(i, "开始写...")
	ruMutex.Lock() // 写操作上锁
	fmt.Println(i, "正在写...")
	time.Sleep(2 * time.Second)
	ruMutex.Unlock()
	fmt.Println(i, "写结束...")

}

func readData(i int) {
	defer wg1.Done()
	fmt.Println(i, "开始读...")
	ruMutex.RLock() // 读操作上锁
	fmt.Println(i, "正在读取数据...")
	time.Sleep(2 * time.Second)
	ruMutex.RUnlock() // 读操作解锁
	fmt.Println(i, "读结束...")
}

// 创建一把锁
var mutex sync.Mutex

var ticket = 10

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for {
		// 上锁
		mutex.Lock()    // g2
		if ticket > 0 { // g1
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出", ticket)
			ticket--
		} else {
			mutex.Unlock() // 条件不满足也要解锁
			fmt.Println(name, "售完,没有了...")
			break
		}
		// 解锁
		mutex.Unlock()
	}
}

var wg sync.WaitGroup

func fun1() {
	for i := 1; i < 10; i++ {
		fmt.Println("fun1中打印", i)
	}
	// 计数值减一
	wg.Done()
}

func fun2() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Println("fun2中打印", i)
	}
}

func fun() {
	defer fmt.Println("defer...")
	// 返回这次函数调用
	//return
	// 跳出这次go协程
	runtime.Goexit()
	fmt.Println("fun函数...")

}

func printNum() {
	for i := 1; i <= 50; i++ {
		fmt.Printf("子程序中打印数字%d\n", i)
	}
}
