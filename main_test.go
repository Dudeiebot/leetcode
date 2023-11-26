package main

import (
	"fmt"
	"testing"
)

// the test files
func TestDupNum(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 4, 1, 4, 5, 2}, true},
		{[]int{1, 2, 3, 4, 5, 6, 7}, false},
		{[]int{2, 1, 3, 4, 3, 5, 5, 7}, true},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.nums), func(t *testing.T) {
			got := DupNum(tc.nums)
			if got != tc.want {
				t.Fatalf("DupNum() = %v; want %v", got, tc.want)
			}
		})
	}
}
