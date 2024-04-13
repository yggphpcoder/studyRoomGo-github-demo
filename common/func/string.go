package _func

import (
	"strconv"
	"strings"
)

// 获取带中文的字符串中子字符串的实际位置，非字节位置
func UnicodeIndex(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	if result > 0 {
		prefix := []byte(str)[0:result]
		rs := []rune(string(prefix))
		result = len(rs)
	}

	return result
}

func Itoa(string2 string) int {
	i, err := strconv.Atoi(string2)
	if err != nil {
		return 0
	}
	return i
}
