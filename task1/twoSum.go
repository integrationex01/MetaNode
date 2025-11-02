package task1

func twoSum(nums []int, target int) []int {

	// 简单解法
	// find different
	/* for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				return []int{i, j}
			}
		}
	}

	return []int{}*/

	// 提高解法
	// create an empty HashTable(already initialized)
	hashTable := map[int]int{}

	// search target-x
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil

}
