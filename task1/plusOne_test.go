package task1

import "testing"

func TestPlusOne(t *testing.T) {
	digits := []int{1, 9, 9}
	expected := []int{2, 0, 0}
	result := plusOne(digits)
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Test failed: expected %v, got %v", expected, result)
		}
	}
}
