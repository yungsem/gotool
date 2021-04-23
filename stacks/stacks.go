package stacks

import (
	"runtime/debug"
	"strings"
)

// Stack 格式化 err 的打印信息，加入堆栈信息
// skip 为需要跳过的调用栈数
func Stack(skip int) []string {
	stk := string(debug.Stack())
	stackArr := strings.Split(stk, "\n")
	stackArr = stackArr[3+(skip+1)*2:]

	var files []string
	funcNameMap := make(map[int]string)

	i := 0
	for _, s := range stackArr {
		if strings.Contains(s, "go/src") {
			continue
		}

		if strings.Contains(s, "\t") { // 文件+行号
			idx := strings.LastIndex(s, "+0x")
			if idx == -1 {
				continue
			}
			s = s[0 : idx-1]
			s = strings.ReplaceAll(s, "\t", "")
			files = append(files, s)
		} else { // 包名。方法名
			idx := strings.LastIndex(s, "/")
			if idx != -1 {
				s = s[idx+1:]
			}
			idx2 := strings.Index(s, "+0x")
			if idx2 != -1 {
				s = s[0 : idx2-1]
			}

			funcNameMap[i] = s
			i++
		}
	}

	var sl []string
	for i, s := range files {
		var sb strings.Builder
		sb.WriteString(s)
		sb.WriteString("->")
		sb.WriteString(funcNameMap[i])
		sl = append(sl, sb.String())
	}
	return sl
}
