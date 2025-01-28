package heap

import (
	"fmt"
	"testing"
)

func TestLastStoneWeight(t *testing.T) {
	tests := []struct {
		stones []int
		want   int
	}{
		{[]int{2, 7, 4, 1, 8, 1}, 1},
		{[]int{1}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.stones), func(t *testing.T) {
			got := LastStoneWeight(tt.stones)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}
