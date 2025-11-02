package task1

import (
	"fmt"
	"testing"
)

func Test_merge1(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		intervals [][]int
		want      [][]int
	}{
		// TODO: Add test cases.
		{
			name: "Test Case 1",

			intervals: [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}},
			want:      [][]int{{1, 10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge1(tt.intervals)
			// TODO: update the condition below to compare got with tt.want.
			if got == nil || fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		intervals [][]int
		want      [][]int
	}{
		// TODO: Add test cases.
		{
			name: "Test Case 1",

			intervals: [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}},
			want:      [][]int{{1, 10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.intervals)
			// TODO: update the condition below to compare got with tt.want.
			if got == nil || fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
