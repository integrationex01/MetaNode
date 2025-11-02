package task1

import "strings"

// 编写一个函数来查找字符串数组中的最长公共前缀。

// 如果不存在公共前缀，返回空字符串 ""。

func longestCommonPrefix(strs []string) string {
	// 1、遍历寻找最短字符串
	// 2、遍历最短字符串的每个字符，依次与其他字符串的对应字符比较
	// 3、如果有不匹配的，返回当前匹配的前缀
	// 4、如果全部匹配，返回最短字符串
	short := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(short) {
			short = strs[i]
		}
	}
	for i := 0; i < len(short); i++ {
		for _, s := range strs {
			if !strings.HasPrefix(s, short[:i+1]) {
				if len(short[:i+1]) == 1 {
					return ""
				}
				return short[:i]
			}
		}
	}

	return short

}
