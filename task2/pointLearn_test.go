package task2

import "testing"

func Test_receivePointer(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		num  *int
		want *int
	}{
		// TODO: Add test cases.
		{
			name: "Test Case 1",
			num:  func() *int { i := 5; return &i }(),
			want: func() *int { i := 15; return &i }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := receivePointer(tt.num)
			// TODO: update the condition below to compare got with tt.want.
			if got == nil || *got != *tt.want {
				t.Errorf("receivePointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_receiveSlicePointer(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums *[]int
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "Test Case 1",
			nums: &[]int{1, 2, 3},
			want: []int{2, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := receiveSlicePointer(tt.nums)
			// TODO: update the condition below to compare got with tt.want.
			if !equalSlices(got, tt.want) {
				t.Errorf("receiveSlicePointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	} else {
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
	}
	return true
}
