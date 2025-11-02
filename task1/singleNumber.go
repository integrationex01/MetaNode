package task1

import "fmt"

// 给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。

// 方法一：使用哈希表记录每个数字出现的次数，最后遍历哈希表找到只出现一次的数字
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func singleNumber1(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		if value, exists := m[num]; exists {
			m[num] = value + 1
		} else {
			m[num] = 1
		}
	}

	for key, value := range m {
		if value == 1 {
			return key
		}
	}

	fmt.Printf("map.len=%d\n", len(m))

	fmt.Println("No single number found")
	return -1
}

// 方法二：使用异或运算的性质，a^a=0，a^0=a，利用这一性质可以将成对出现的数字抵消掉，最终剩下的就是只出现一次的数字
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func singleNumber(nums []int) (single int) {
	for _, num := range nums {
		single ^= num
	}
	return single
}
