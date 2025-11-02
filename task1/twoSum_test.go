package task1

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums   []int
		target int
		want   []int
	}{
		// TODO: Add test cases.
		{
			name:   "Test Case 1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			// TODO: update the condition below to compare got with tt.want.
			if got == nil || fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
