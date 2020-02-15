package pk1

import (
	"TestGo/utils"
	"TestGo/utils/timeutils"
	"fmt"
)

func MyTest1() {
	utils.Count()
	timeutils.PringTime()
}
func init() {
	fmt.Println("pk1包下的init初始化函数")
}
