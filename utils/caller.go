package utils

import (
	"fmt"
	"runtime"
)

// demo: /xxx/xxx/LogX/utils/caller_test.go:9
// n = 0: 打印 调用 GetCallerName的 代码位置
// n = 1: 打印 调用 调用GetCallerName的函数的 代码位置
// ...... 以此类推
func GetCallerName(n int) string {
	n++
	_, file, line, ok := runtime.Caller(n)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s:%d", file, line)
}