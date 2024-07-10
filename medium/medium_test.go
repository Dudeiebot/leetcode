package medium

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGroupAnagram(t *testing.T) {
	tests := []struct {
		strs []string
		want [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.strs), func(t *testing.T) {
			got := GroupAnagram(tt.strs)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.nums, tt.k), func(t *testing.T) {
			got := topKFrequent(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := productExceptSelf(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := longestConsecutive(tt.nums)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoSumII(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.nums, tt.target), func(t *testing.T) {
			got := twoSumII(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, [][]int{}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := threeSum(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.height), func(t *testing.T) {
			got := maxArea(tt.height)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbb", 1},
		{"pwwkew", 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.s), func(t *testing.T) {
			got := lengthOfLongestSubString(tt.s)
			if got != tt.want {
				t.Fatalf("%v, %v", got, tt.want)
			}
		})
	}
}
