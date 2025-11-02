package task1

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(tests)
	expected := "fl"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
