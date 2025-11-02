package task1

import "testing"

func Test_removeDuplicates1(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Test Case 1",
			nums: []int{1, 1, 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates1(tt.nums)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("removeDuplicates1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Test Case 1",
			nums: []int{1, 1, 2, 2, 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.nums)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
