package task1

import "testing"

func TestIsPalindrome1(t *testing.T) {
	x := 121
	expected := true
	result := isPalindrome1(x)
	if result != expected {
		t.Errorf("isPalindrome(%d) = %v; want %v", x, result, expected)
	}
}

func TestIsPalindrome(t *testing.T) {
	x := 121
	expected := true
	result := isPalindrome(x)
	if result != expected {
		t.Errorf("isPalindrome(%d) = %v; want %v", x, result, expected)
	}
}
