package main

import (
	"tesou.io/platform/foot-parent/foot-spider/launch"
)


func main() {
	//开启SQL输出
	//launch.Before_spider_baseFace()
	//launch.Before_spider_match()
	//launch.Spider_match(4)
	launch.Spider_asiaLastNew(true)
	launch.Spider_baseFace(true)
}
