package task1

import "sort"

// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

func merge1(intervals [][]int) [][]int {
	// for i := 0; i < len(intervals)-1; i++ {
	// 	if intervals[i+1][0] < intervals[i][0] {
	// 		intervals[i], intervals[i+1] = intervals[i+1], intervals[i]
	// 		i = -1
	// 	}
	// }
	// QuickSort1(intervals, 0, len(intervals)-1)
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})
	if len(intervals) == 0 || len(intervals) == 1 {
		return intervals
	} else {
		for i := 0; i < len(intervals); i++ {
			for j := i + 1; j < len(intervals); j++ {
				if (intervals[j][0] <= intervals[i][1]) && (intervals[j][1] >= intervals[i][0]) {
					// 有重叠，合并
					intervals[i][1] = max(intervals[j][1], intervals[i][1])
					intervals[i][0] = min(intervals[j][0], intervals[i][0])
					// 删除intervals[j]
					intervals = append(intervals[:j], intervals[j+1:]...)
					j--

				}
			}
		}
		return intervals
	}

}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})
	res := [][]int{}
	for _, interval := range intervals {
		n := len(res)
		if n > 0 && interval[0] <= res[n-1][1] {
			res[n-1][1] = max(res[n-1][1], interval[1])
		} else {
			res = append(res, interval)
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func QuickSort1(arr [][]int, left, right int) {
	if left < right {
		i := arrAdjust(arr, left, right)
		QuickSort1(arr, left, i-1)  //调整左边的序列
		QuickSort1(arr, i+1, right) //调整右边的序列
	}
}

//返回调整后基准数的位置
func arrAdjust(arr [][]int, left, right int) int {
	mid := (left + right) / 2
	arr[left], arr[mid] = arr[mid], arr[left] //可以选择中间的数作为基准
	x := arr[left][0]                         //基准
	i, j := left, right
	for i < j {
		//从右向左找小于x的数
		for i < j && arr[j][0] >= x {
			j--
		}
		//从左向右找大于x的数
		for i < j && arr[i][0] <= x {
			i++
		}
		//交换
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// i=j结束扫描
	// 基准数归位，现在左边的序列小于基准数，右边的序列大于基准数
	arr[left], arr[i] = arr[i], arr[left]
	return i
}
