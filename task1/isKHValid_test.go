package task1

import (
	"testing"
)

func TestIsKHValid1(t *testing.T) {
	tests := "()[]{}"
	if isValid1(tests) {
		t.Log("测试通过")
	} else {
		t.Error("测试失败")
	}
}

func TestIsKHValid(t *testing.T) {
	tests := "()[]{}"
	if isValid(tests) {
		t.Log("测试通过")
	} else {
		t.Error("测试失败")
	}
}
