package task1

import (
	"log"
)

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

// 有效字符串需满足：

// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。

func isValid1(s string) bool {
	b := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		b = append(b, s[i])
		l := len(b)
		if l >= 2 && ((b[l-1] == b[l-2]+1) || (b[l-1] == b[l-2]+2)) {
			log.Println(string(b))
			b = b[:l-2]
		}
	}
	log.Println(string(b))
	return len(b) == 0
}

// 栈方案：
func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, char := range s {
		switch char {
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		default:
			if len(stack) == 0 || stack[len(stack)-1] != char {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}

		}
	}
	return len(stack) == 0
}
