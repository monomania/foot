package main

import "tesou.io/platform/foot-parent/foot-api/common/base"

func main() {
	base.Log.Info("日志1")
	base.Log.Info("日志2")
	base.Log.Error("错误的日志")
	base.Log.Info("日志3")

}
