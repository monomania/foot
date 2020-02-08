package utils

import (
	"runtime"
	"strings"
	"time"
)

/**
获取当前的函数名称
*/
func RunFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	funcName := f.Name()
	funcName = funcName[strings.LastIndex(funcName, ".")+1:]
	return funcName
}

//获取相差时间
func GetHourDiffer(t1, t2 time.Time) int64 {
	var hour int64
	if t1.After(t2) {
		diff := t1.Unix() - t2.Unix()
		hour = diff / 3600
		return hour
	} else {
		diff := t2.Unix() - t1.Unix()
		hour = diff / 3600
		return 0 - hour
	}
}
