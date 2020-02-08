package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"tesou.io/platform/foot-parent/foot-robot/helper"
)

func main() {
	robotgo.Sleep(1)

	title := robotgo.GetTitle()
	fmt.Println("title@@@ ", title)

	util := &helper.RobotHelper{}
	util.Tips()
}

func tes11t(){
	robotgo.Sleep(1)

	title := robotgo.GetTitle()
	fmt.Println("title@@@ ", title)
	for i := 0; i < 1; i++ {
		robotgo.TypeStr("Hello World")
		robotgo.TypeStr("测试")
		// robotgo.TypeString("测试")

		robotgo.TypeStr("山达尔星新星军团, galaxy. こんにちは世界.")


		// ustr := uint32(robotgo.CharCodeAt("测试", 0))
		// robotgo.UnicodeType(ustr)

		//robotgo.KeyTap("enter","control")
		robotgo.KeyTap("enter","control")
		// robotgo.TypeString("en")
		robotgo.KeyTap("i", "alt", "command")

		arr := []string{"alt", "command"}
		robotgo.KeyTap("i", arr)

		robotgo.WriteAll("测试")
		text, err := robotgo.ReadAll()
		if err == nil {
			fmt.Println(text)
		}
	}
}
