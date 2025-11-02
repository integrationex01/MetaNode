package task1

import (
	"fmt"
	"testing"
	"time"
)

func TestSingleNumber(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2}
	start := time.Now()
	result := singleNumber(nums)
	elapsed := time.Since(start)
	expected := 4
	if result != expected {
		panic("test failed")
	}
	fmt.Printf("TestSingleNumber took %s\n", elapsed)
}

func TestSingleNumber1(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2}
	start := time.Now()
	result := singleNumber1(nums)
	elapsed := time.Since(start)
	expected := 4
	if result != expected {
		panic("test failed")
	}
	fmt.Printf("TestSingleNumber1 took %s\n", elapsed)
}
