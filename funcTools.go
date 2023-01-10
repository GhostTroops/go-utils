package go_utils

import (
	"log"
	"runtime"
	"strings"
)

// 将该方法放到方法中运行，就可以打印出所有调用该方法的链路出来
func PrintCaller() {
	var i = 0
	for i < 2000 {
		i++
		if pc, file, line, ok := runtime.Caller(i); ok {
			fc := runtime.FuncForPC(pc)
			log.Printf("<-%s %s file:%s (line:%d)\n", strings.Repeat(":", i-1), fc.Name(), line, file) // , runtime.CallersFrames([]uintptr{pc})
			if "main.main" == fc.Name() {
				break
			}
		} else {
			break
		}
	}
}
