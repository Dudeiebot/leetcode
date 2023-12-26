package main

import (
	"fmt"
	"reflect"
	"testing"
)

// the test files
// func TestDupNum(t *testing.T) {
// 	tests := []struct {
// 		nums []int
// 		want bool
// 	}{
// 		{[]int{1, 2, 3, 4, 1, 4, 5, 2}, true},
// 		{[]int{1, 2, 3, 4, 5, 6, 7}, false},
// 		{[]int{2, 1, 3, 4, 3, 5, 5, 7}, true},
// 	}
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%v", tc.nums), func(t *testing.T) {
// 			got := DupNum(tc.nums)
// 			if got != tc.want {
// 				t.Fatalf("DupNum() = %v; want %v", got, tc.want)
// 			}
// 		})
// 	}
// }

// func TestIsAnagram(t *testing.T) {
// 	tests := []struct {
// 		s, t string
// 		want bool
// 	}{
// 		{"anagram", "nagaram", true},
// 		{"car", "rat", false},
// 	}
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%v and %v", tc.s, tc.t), func(t *testing.T) {
// 			got := isAnagram(tc.s, tc.t)
// 			if got != tc.want {
// 				t.Fatalf("Valid Anagram() = %v; want %v", got, tc.want)
// 			}
// 		})
// 	}
// }
//
// func TestIsPalindrome(t *testing.T) {
// 	tests := []struct {
// 		s    string
// 		want bool
// 	}{
// 		{"A man, a plan, a canal: Panama", true},
// 		{"race a car", false},
// 		{" ", true},
// 	}
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%v", tc.s), func(t *testing.T) {
// 			got := IsPalindrome(tc.s)
// 			if got != tc.want {
// 				t.Fatalf("Valid Palindrome() = %v; want %v", got, tc.want)
// 			}
// 		})
// 	}
// }

// func TestIsSingeNos(t *testing.T) {
// 	tests := []struct {
// 		nums []int
// 		want int
// 	}{
// 		{[]int{2, 2, 1}, 1},
// 		{[]int{4, 1, 2, 1, 2}, 4},
// 		{[]int{1}, 1},
// 	}
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%v", tc.nums), func(t *testing.T) {
// 			got := IsSingleNos(tc.nums)
// 			if got != tc.want {
// 				t.Fatalf("The Single Nos() = %v; want %v", got, tc.want)
// 			}
// 		})
// 	}
// }

// func TestTwoSums(t *testing.T) {
// 	tests := []struct {
// 		nums   []int
// 		target int
// 		want   []int
// 	}{
// 		{[]int{2, 1, 3, 4, 5}, 9, []int{3, 4}},
// 		{[]int{5, 3, 2, 6, 7, 8}, 5, []int{1, 2}},
// 	}
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%v and %v", tc.nums, tc.target), func(t *testing.T) {
// 			got := twoSums(tc.nums, tc.target)
// 			if !reflect.DeepEqual(got, tc.want) {
// 				t.Fatalf("Valid two sums in the slice() = %v; want %v", got, tc.want)
// 			}
// 		})
// 	}
// }
