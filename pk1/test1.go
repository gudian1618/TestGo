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
	fmt.Println("utils包下的test文件中的初始化函数")
}
