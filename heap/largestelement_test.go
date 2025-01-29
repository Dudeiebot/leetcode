package heap

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.nums, tt.k), func(t *testing.T) {
			got := FindKthLargest(tt.nums, tt.k)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}
