package main

import "tesou.io/platform/foot-parent/foot-core/launch"

func main() {

	//生成数据库表
	launch.GenTable()
	//清空数据表
	//launch.TruncateTable()
}
