package common

import (
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"time"
)

// 带时间打印
func PrintT(msg string, items ...any) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fullFuncName, file, line := getCurrentCaller()
	fname := strings.Split(fullFuncName, ".")
	prefix := fmt.Sprintf("%s %s:%d[func:%s]", timestamp, file, line, fname[len(fname)-1])
	if hasPlaceholder(msg) && len(items) > 0 {
		fmt.Printf(prefix+msg+"\n", items...)
	} else {
		newItems := make([]interface{}, len(items)+1)
		newItems[0] = prefix + msg
		if len(items) > 0 {
			copy(newItems[1:], items)
		}
		fmt.Println(newItems...)
	}
}

// 生成带时间格式日志字符串
func SprintT(msg string, items ...any) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	prefix := timestamp + "  "
	return fmt.Sprintf(prefix+msg, items...)
}

// hasPlaceholder 检查字符串是否包含格式占位符
func hasPlaceholder(s string) bool {
	pattern := `%(?:[-+# 0]*(?:\d+|\*)?(?:\.(?:\d+|\*))?)?[vTtbcdoqxXUeEfFgGps]`

	re := regexp.MustCompile(pattern)

	// 先移除所有 %%（转义的百分号），避免干扰
	s = regexp.MustCompile(`%%`).ReplaceAllString(s, "")

	return re.MatchString(s)
}
func getCurrentCaller() (fullFuncName, file string, line int) {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return "unknown", "unknown", 0
	}
	fullFuncName = runtime.FuncForPC(pc).Name()
	return fullFuncName, file, line
}
