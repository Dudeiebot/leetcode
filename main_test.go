package main

import (
	//	"fmt"
	//	"reflect"
	"testing"
)

//
//
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

// func TestIsNosPalindrome(t *testing.T) {
// 	tests := []struct {
// 		x    int
// 		want bool
// 	}{
// 		{121, true},
// 		{-121, true},
// 		{10, false},
// 		{21312, true},
// 	}
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%v", tc.x), func(t *testing.T) {
// 			got := isNosPalindrome(tc.x)
// 			if got != tc.want {
// 				t.Fatalf("got %v, want %v", got, tc.want)
// 			}
// 		})
// 	}
// }

// func TestRom2Int(t *testing.T) {
// 	tests := []struct {
// 		s    string
// 		want int
// 	}{
// 		{"MCDIX", 1409},
// 		{"III", 3},
// 		{"MCMIII", 1903},
// 	}
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%v", tc.s), func(t *testing.T) {
// 			got := rom2Int(tc.s)
// 			if got != tc.want {
// 				t.Fatalf("got %v, want %v", got, tc.want)
// 			}
// 		})
// 	}
// }

// func TestLongestPrefix(t *testing.T) {
// 	tests := []struct {
// 		s    []string
// 		want string
// 	}{
// 		{[]string{"flower", "flow", "flight"}, "f"},
// 		{[]string{"dog", "racecar", "car"}, ""},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%v", tc.s), func(t *testing.T) {
// 			got := longestPrefix(tc.s)
// 			if got != tc.want {
// 				t.Fatalf("got %v, want %v", got, tc.want)
// 			}
// 		})
// 	}
// }

// func TestCheckParenthesis(t *testing.T) {
// 	tests := []struct {
// 		s    string
// 		want bool
// 	}{
// 		{"()", true},
// 		{"()[]{}", true},
// 		{"(]", false},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%v", tc.s), func(t *testing.T) {
// 			got := checkParenthesis(tc.s)
// 			if got != tc.want {
// 				t.Fatalf("got %v, want %v", got, tc.want)
// 			}
// 		})
// 	}
// }

// func TestMergeTwoLists(t *testing.T) {
// 	// We first of all create a linked list from an array, and (the linked list is showing that everything in the array is linked), the we convert the link to arr and the merge it
//
// 	createList := func(arr []int) *ListNode {
// 		tempNode := &ListNode{}
// 		currentNode := tempNode
// 		for _, val := range arr {
// 			currentNode.Next = &ListNode{Val: val}
// 			currentNode = currentNode.Next
// 		}
// 		return tempNode.Next
// 	}
//
// 	list2Arr := func(list *ListNode) []int {
// 		res := make([]int, 0)
// 		for list != nil {
// 			res = append(res, list.Val)
// 			list = list.Next
// 		}
// 		return res
// 	}
//
// 	tests := []struct {
// 		list1 []int
// 		list2 []int
// 		want  []int
// 	}{
// 		{[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
// 		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
// 		{[]int{}, []int{0}, []int{0}},
// 		{[]int{}, []int{}, []int{}},
// 	}
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("merge %v and %v", tc.list1, tc.list2), func(t *testing.T) {
// 			list1 := createList(tc.list1)
// 			list2 := createList(tc.list2)
// 			got := list2Arr(mergeTwoLists(list1, list2))
// 			if !reflect.DeepEqual(got, tc.want) {
// 				t.Fatalf("got %v, mergeTwoLists want %v", got, tc.want)
// 			}
// 		})
// 	}
// }

// func TestRemoveDuplicates(t *testing.T) {
// 	tests := []struct {
// 		nums []int
// 		want int
// 	}{
// 		{[]int{1, 1, 2}, 2},
// 		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%v", tc.nums), func(t *testing.T) {
// 			got := removeDuplicates(tc.nums)
// 			if got != tc.want {
// 				t.Fatalf("got %v, want %v", got, tc.want)
// 			}
// 		})
// 	}
// }
//
// func TestRemoveElements(t *testing.T) {
// 	tests := []struct {
// 		nums []int
// 		val  int
// 		want int
// 	}{
// 		{[]int{3, 1, 2, 3}, 3, 2},
// 		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%v", tc.nums), func(t *testing.T) {
// 			got := removeElements(tc.nums, tc.val)
// 			if got != tc.want {
// 				t.Fatalf("got %v, want %v", got, tc.want)
// 			}
// 		})
// 	}
// }

// func TestStrStr(t *testing.T) {
// 	tests := []struct {
// 		mainString, substring string
// 		want                  int
// 	}{
// 		{"sadbutsad", "sad", 0},
// 		{"leetcode", "leeto", -1},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(
// 			fmt.Sprintf("the mainstring %v and substring %v", tc.mainString, tc.substring),
// 			func(t *testing.T) {
// 				got := strstr(tc.mainString, tc.substring)
// 				if got != tc.want {
// 					t.Fatalf("got %v, want %v", got, tc.want)
// 				}
// 			},
// 		)
// 	}
// }

// func TestSearchInsert(t *testing.T) {
// 	tests := []struct {
// 		nums   []int
// 		target int
// 		want   int
// 	}{
// 		{[]int{1, 3, 5, 6}, 5, 2},
// 		{[]int{1, 3, 5, 6}, 2, 1},
// 		{[]int{1, 3, 5, 6}, 7, 4},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%v %v", tc.nums, tc.target), func(t *testing.T) {
// 			got := searchInsert(tc.nums, tc.target)
// 			if got != tc.want {
// 				t.Fatalf("got %v, want %v", got, tc.want)
// 			}
// 		})
// 	}
// }

// func TestLengthOfLastWord(t *testing.T) {
// 	tests := []struct {
// 		s    string
// 		want int
// 	}{
// 		{"Hello World", 5},
// 		{"   fly me   to   the moon  ", 4},
// 		{"luffy is still joyboy", 6},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%v", tc.s), func(t *testing.T) {
// 			got := lengthOfLastWord(tc.s)
// 			if got != tc.want {
// 				t.Fatalf("got %v, want %v", got, tc.want)
// 			}
// 		})
// 	}
// }

// func TestPlusOne(t *testing.T) {
// 	tests := []struct {
// 		digits []int
// 		want   []int
// 	}{
// 		{[]int{1, 2, 3}, []int{1, 2, 4}},
// 		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
// 		{[]int{9}, []int{1, 0}},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%v", tc.digits), func(t *testing.T) {
// 			got := plusOne(tc.digits)
// 			if !reflect.DeepEqual(got, tc.want) {
// 				t.Fatalf("got %v, want %v", got, tc.want)
// 			}
// 		})
// 	}
// }

// func TestAddBinary(t *testing.T) {
// 	tests := []struct {
// 		a, b string
// 		want string
// 	}{
// 		{"11", "1", "100"},
// 		{"1010", "1011", "10101"},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%v %v", tc.a, tc.b), func(t *testing.T) {
// 			got := addBinary(tc.a, tc.b)
// 			if got != tc.want {
// 				t.Fatalf("got %v, want %v", got, tc.want)
// 			}
// 		})
// 	}
// }

// func TestMySqrt(t *testing.T) {
// 	tests := []struct {
// 		n    int
// 		want int
// 	}{
// 		{16, 4},
// 		{8, 2},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%v", tc.n), func(t *testing.T) {
// 			got := mySqrt(tc.n)
// 			if got != tc.want {
// 				t.Fatalf("got %v, want %v", got, tc.want)
// 			}
// 		})
// 	}
// }

// func TestClimbStairs(t *testing.T) {
// 	tests := []struct {
// 		n    int
// 		want int
// 	}{
// 		{0, 0},
// 		{1, 1},
// 		{3, 3},
// 		{40, 165580141},
// 	}
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("n=%v", tc.n), func(t *testing.T) {
// 			got := climbStairs(tc.n)
// 			if got != tc.want {
// 				t.Fatalf("Fibonacci() = %v; want %v", got, tc.want)
// 			}
// 		})
// 	}
// }

// func TestDelDuplicates(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		input    *ListNode
// 		expected *ListNode
// 	}{
// 		{
// 			name: "No duplicates",
// 			input: &ListNode{
// 				Val:  1,
// 				Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}},
// 			},
// 			expected: &ListNode{
// 				Val:  1,
// 				Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}},
// 			},
// 		},
// 		{
// 			name: "Duplicates in the middle",
// 			input: &ListNode{
// 				Val:  1,
// 				Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
// 			},
// 			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
// 		},
// 		{
// 			name: "Duplicates at the end",
// 			input: &ListNode{
// 				Val:  1,
// 				Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: nil}},
// 			},
// 			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
// 		},
// 		{
// 			name:     "Empty list",
// 			input:    nil,
// 			expected: nil,
// 		},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			result := delDuplicates(tc.input)
// 			if !reflect.DeepEqual(result, tc.expected) {
// 				t.Errorf("Expected %v, got %v", tc.expected, result)
// 			}
// 		})
// 	}
// }

// func TestMerge(t *testing.T) {
// 	tests := []struct {
// 		nums1 []int
// 		nums2 []int
// 		m     int
// 		n     int
// 		want  []int
// 	}{
// 		{[]int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}, 3, 3, []int{1, 2, 2, 3, 5, 6}},
// 		{[]int{4, 5, 6, 0, 0, 0}, []int{1, 2, 3}, 3, 3, []int{1, 2, 3, 4, 5, 6}},
// 		{[]int{0}, []int{1}, 0, 1, []int{1}},
// 		{[]int{1}, []int{}, 1, 0, []int{1}},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(" ", func(t *testing.T) {
// 			merge(tc.nums1, tc.m, tc.nums2, tc.n)
// 			if !reflect.DeepEqual(tc.nums1, tc.want) {
// 				t.Fatalf("got %v, want %v", tc.nums1, tc.want)
// 				// nums1 get modified and can be represented as our got
// 			}
// 		})
// 	}
// }

// func TestInorderTraversal(t *testing.T) {
// 	tests := []struct {
// 		root *TreeNode
// 		want []int
// 	}{
// 		{(*TreeNode)(nil), []int{}},
// 		{&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, []int{2, 1, 3}},
// 		{
// 			&TreeNode{
// 				Val:   1,
// 				Left:  &TreeNode{Val: 3},
// 				Right: &TreeNode{Val: 2},
// 			},
// 			[]int{3, 1, 2},
// 		},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(" ", func(t *testing.T) {
// 			got := inorderTraversal(tc.root)
// 			if !reflect.DeepEqual(got, tc.want) {
// 				t.Fatalf("got %v, want %v", got, tc.want)
// 			}
// 		})
// 	}
// }
//
// func TestSameTree(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		p    *TreeNode
// 		q    *TreeNode
// 		want bool
// 	}{
// 		{"the first", nil, nil, true},
// 		{
// 			"the second",
// 			&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
// 			&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}},
// 			false,
// 		},
// 		{
// 			"the third",
// 			&TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 3}},
// 			&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
// 			false,
// 		},
// 		{
// 			"the third",
// 			&TreeNode{Val: 1, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}},
// 			&TreeNode{Val: 1, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}},
// 			true,
// 		},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			got := isSameTree(tc.p, tc.q)
// 			if got != tc.want {
// 				t.Fatalf("got %v, want %v", got, tc.want)
// 			}
// 		})
// 	}
// }

// func TestIsSymmetry(t *testing.T) {
// 	tests := []struct {
// 		root *TreeNode
// 		want bool
// 	}{
// 		{
// 			&TreeNode{
// 				Val: 1,
// 				Left: &TreeNode{
// 					Val:   2,
// 					Left:  &TreeNode{Val: 3},
// 					Right: &TreeNode{Val: 4},
// 				},
// 				Right: &TreeNode{
// 					Val:   2,
// 					Left:  &TreeNode{Val: 4},
// 					Right: &TreeNode{Val: 3},
// 				},
// 			},
// 			true, // Symmetric tree
// 		},
// 		{
// 			&TreeNode{
// 				Val: 1,
// 				Left: &TreeNode{
// 					Val:   2,
// 					Right: &TreeNode{Val: 3},
// 				},
// 				Right: &TreeNode{
// 					Val:   2,
// 					Right: &TreeNode{Val: 3},
// 				},
// 			},
// 			false, // Asymmetric tree
// 		},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(" ", func(t *testing.T) {
// 			got := isSymmetry(tc.root)
// 			if got != tc.want {
// 				t.Fatalf("got %v, want %v", got, tc.want)
// 			}
// 		})
// 	}
// }
