package task1

import "fmt"

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数

// 方法一：将整数转换为字符串，使用双指针从两端向中间遍历比较字符是否相同
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func isPalindrome1(x int) bool {
	xStr := fmt.Sprintf("%d", x)
	for i := 0; i < len(xStr)/2; i++ {
		if xStr[i] != xStr[len(xStr)-1-i] {
			return false
		}
	}
	return true
}

// 方法二：不将整数转换为字符串，反转整数的一半并与另一半比较
// 时间复杂度：O(log10(n))
// 空间复杂度：O(1)
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	o := x
	r := 0
	for x != 0 {
		d := x % 10
		r = r*10 + d
		x /= 10
	}
	return r == o
}
