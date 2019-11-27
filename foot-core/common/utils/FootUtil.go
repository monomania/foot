package utils

import (
	"runtime"
	"strings"
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
