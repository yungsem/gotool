package errorf

import (
	"runtime/debug"
	"strings"
)

// Format 格式化 err 的打印信息，加入堆栈信息
func Format(err error) string {
	stack := string(debug.Stack())
	stackArr := strings.Split(stack, "\n")
	stackArr = stackArr[5:]

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
			s = s[0:idx]
			s = strings.ReplaceAll(s, "\t", "")
			files = append(files, s)
		} else { // 包名。方法名
			idx := strings.LastIndex(s, "/")
			if idx != -1 {
				s = s[idx+1:]
			}
			idx2 := strings.Index(s, "(")
			if idx2 == -1 {
				continue
			}
			s = s[0:idx2]
			funcNameMap[i] = s
			i++
		}
	}

	var sb strings.Builder
	sb.WriteString(err.Error())
	sb.WriteString("\n")
	for i, s := range files {
		sb.WriteString(" ->")
		sb.WriteString(s)
		sb.WriteString("[")
		sb.WriteString(funcNameMap[i])
		sb.WriteString("]")
		if i != len(files)-1 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
