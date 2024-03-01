package main

import (
	"fmt"
	util "leetCode/struct"
	"reflect"
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

// func TestPreOrderTraversal(t *testing.T) {
// 	tests := []struct {
// 		root *TreeNode
// 		want []int
// 	}{
// 		{(*TreeNode)(nil), []int{}},
// 		{&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, []int{1, 2, 3}},
// 		{
// 			&TreeNode{
// 				Val:   1,
// 				Left:  &TreeNode{Val: 3},
// 				Right: &TreeNode{Val: 2},
// 			},
// 			[]int{1, 3, 2},
// 		},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(" ", func(t *testing.T) {
// 			got := preOrderTraversal(tc.root)
// 			if !reflect.DeepEqual(got, tc.want) {
// 				t.Fatalf("got %v, want %v", got, tc.want)
// 			}
// 		})
// 	}
// }

// func TestPostOrderTraversal(t *testing.T) {
// 	tests := []struct {
// 		root *TreeNode
// 		want []int
// 	}{
// 		{(*TreeNode)(nil), []int{}},
// 		{&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, []int{2, 3, 1}},
// 		{
// 			&TreeNode{
// 				Val:   1,
// 				Left:  &TreeNode{Val: 3},
// 				Right: &TreeNode{Val: 2},
// 			},
// 			[]int{3, 2, 1},
// 		},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(" ", func(t *testing.T) {
// 			got := postOrderTraversal(tc.root)
// 			if !reflect.DeepEqual(got, tc.want) {
// 				t.Fatalf("got %v, want %v", got, tc.want)
// 			}
// 		})
// 	}
// }
//
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

// func TestMaximumDepth(t *testing.T) {
// 	tests := []struct {
// 		root *TreeNode
// 		want int
// 	}{
// 		{nil, 0},               // Empty tree
// 		{&TreeNode{Val: 1}, 1}, // Single-node tree
// 		{
// 			&TreeNode{
// 				Val:   1,
// 				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
// 				Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 5}},
// 			},
// 			3, // Balanced tree
// 		},
// 		{
// 			&TreeNode{
// 				Val:   1,
// 				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
// 				Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}}},
// 			},
// 			4, // Unbalanced tree
// 		},
// 	}
// 	for _, tc := range tests {
// 		t.Run(" ", func(t *testing.T) {
// 			got := maximumDepth(tc.root)
// 			if got != tc.want {
// 				t.Errorf("maximumDepth(%v) = %v, want %v", tc.root, got, tc.want)
// 			}
// 		})
// 	}
// }

// func TestSortedArrayToBST(t *testing.T) {
// 	tests := []struct {
// 		nums     []int
// 		expected *TreeNode
// 	}{
// 		{
// 			nums: []int{-10, -3, 0, 5, 9},
// 			expected: &TreeNode{
// 				Val:   0,
// 				Left:  &TreeNode{Val: -3, Left: &TreeNode{Val: -10}},
// 				Right: &TreeNode{Val: 9, Left: &TreeNode{Val: 5}},
// 			},
// 		},
// 		// Add more test cases as needed
// 	}
//
// 	for _, tc := range tests {
// 		t.Run("", func(t *testing.T) {
// 			result := sortedArray2BST(tc.nums)
// 			if !reflect.DeepEqual(result, tc.expected) {
// 				t.Errorf("Expected %v, got %v", tc.expected, result)
// 			}
// 		})
// 	}
// }

// func TestIsBalanced(t *testing.T) {
// 	tests := []struct {
// 		root     *TreeNode
// 		expected bool
// 	}{
// 		{nil, true}, // An empty tree is balanced
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
// 			true, // Balanced tree
// 		},
// 		{
// 			&TreeNode{
// 				Val: 1,
// 				Left: &TreeNode{
// 					Val:  2,
// 					Left: &TreeNode{Val: 3},
// 					Right: &TreeNode{
// 						Val:   3,
// 						Left:  &TreeNode{Val: 4},
// 						Right: &TreeNode{Val: 4},
// 					},
// 				},
// 				Right: &TreeNode{Val: 0},
// 			},
// 			false,
// 		},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run("", func(t *testing.T) {
// 			result := isBalanced(tc.root)
// 			if result != tc.expected {
// 				t.Errorf("Expected %v, got %v", tc.expected, result)
// 			}
// 		})
// 	}
// }

// func TestMinDepth(t *testing.T) {
// 	tests := []struct {
// 		root     *TreeNode
// 		expected int
// 	}{
// 		{nil, 0},
// 		{
// 			&TreeNode{
// 				Val:  2,
// 				Left: nil,
// 				Right: &TreeNode{
// 					Val:  3,
// 					Left: nil,
// 					Right: &TreeNode{
// 						Val:  4,
// 						Left: nil,
// 						Right: &TreeNode{
// 							Val:  5,
// 							Left: nil,
// 							Right: &TreeNode{
// 								Val:   6,
// 								Left:  nil,
// 								Right: nil,
// 							},
// 						},
// 					},
// 				},
// 			},
// 			5,
// 		},
// 		{&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, 2},
// 		{
// 			&TreeNode{
// 				Val:   1,
// 				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
// 				Right: &TreeNode{Val: 4},
// 			},
// 			2,
// 		},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run("", func(t *testing.T) {
// 			result := minimumDepth(tc.root)
// 			if result != tc.expected {
// 				t.Errorf("Expected %v, got %v", tc.expected, result)
// 			}
// 		})
// 	}
// }

// func TestHasPathSum(t *testing.T) {
// 	tests := []struct {
// 		root      *TreeNode
// 		targetSum int
// 		expected  bool
// 	}{
// 		{nil, 0, false},
// 		{
// 			&TreeNode{
// 				Val:   1,
// 				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
// 				Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}},
// 			},
// 			7,
// 			true,
// 		},
// 		{&TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 8}}, 13, true},
// 		{&TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 8}}, 10, false},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run("", func(t *testing.T) {
// 			result := hasPathSum(tc.root, tc.targetSum)
// 			if result != tc.expected {
// 				t.Errorf(
// 					"For target sum %d, expected %v, got %v",
// 					tc.targetSum,
// 					tc.expected,
// 					result,
// 				)
// 			}
// 		})
// 	}
// }

// func TestGeneratePascal(t *testing.T) {
// 	tests := []struct {
// 		numRows int
// 		want    [][]int
// 	}{
// 		{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
// 		{1, [][]int{{1}}},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("numsAndRows= %v", tc.numRows), func(t *testing.T) {
// 			got := generatePascal(tc.numRows)
// 			if !reflect.DeepEqual(got, tc.want) {
// 				t.Fatalf("got = %v, want = %v", got, tc.want)
// 			}
// 		})
// 	}
// }
//
// func TestGetRow(t *testing.T) {
// 	tests := []struct {
// 		rowIndex int
// 		want     []int
// 	}{
// 		{4, []int{1, 4, 6, 4, 1}},
// 		{1, []int{1, 1}},
// 		{3, []int{1, 3, 3, 1}},
// 		{0, []int{1}},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("RowIndex= %v", tc.rowIndex), func(t *testing.T) {
// 			got := getRow(tc.rowIndex)
// 			if !reflect.DeepEqual(got, tc.want) {
// 				t.Fatalf("got = %v, want = %v", got, tc.want)
// 			}
// 		})
// 	}
// }

// func TestBestTimeStock(t *testing.T) {
// 	tests := []struct {
// 		prices []int
// 		want   int
// 	}{
// 		{[]int{7, 1, 5, 3, 6, 4}, 5},
// 		{[]int{7, 6, 4, 3, 1}, 0},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%v", tc.prices), func(t *testing.T) {
// 			got := bestTimeStock(tc.prices)
// 			if got != tc.want {
// 				t.Fatalf("got = %v, want = %v", got, tc.want)
// 			}
// 		})
// 	}
// }
//
// func TestMaxProfit(t *testing.T) {
// 	tests := []struct {
// 		prices []int
// 		want   int
// 	}{
// 		{[]int{7, 1, 5, 3, 6, 4}, 7},
// 		{[]int{1, 2, 3, 4, 5}, 4},
// 		{[]int{7, 6, 4, 3, 1}, 0},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%v", tc.prices), func(t *testing.T) {
// 			got := maxProfit(tc.prices)
// 			if got != tc.want {
// 				t.Fatalf("got = %v, want = %v", got, tc.want)
// 			}
// 		})
// 	}
// }

// func TestSingeNum(t *testing.T) {
// 	tests := []struct {
// 		nums []int
// 		want int
// 	}{
// 		{[]int{2, 2, 2, 1}, 1},
// 		{[]int{4, 1, 2, 1, 2, 1, 2}, 4},
// 		{[]int{1}, 1},
// 	}
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%v", tc.nums), func(t *testing.T) {
// 			got := singleNum(tc.nums)
// 			if got != tc.want {
// 				t.Fatalf("The Single Nos() = %v; want %v", got, tc.want)
// 			}
// 		})
// 	}
// }

// func TestHasCycle(t *testing.T) {
// 	tests := []struct {
// 		input    *ListNode
// 		expected bool
// 	}{
// 		{
// 			input: func() *ListNode {
// 				node1 := &ListNode{Val: 1}
// 				node2 := &ListNode{Val: 2}
// 				node3 := &ListNode{Val: 3}
// 				node4 := &ListNode{Val: 4}
//
// 				node1.Next = node2
// 				node2.Next = node3
// 				node3.Next = node4
// 				node4.Next = node2 // Creating a cycle by connecting the last node to the second node
//
// 				return node1
// 			}(),
// 			expected: true,
// 		},
// 		{
// 			input: func() *ListNode {
// 				node5 := &ListNode{Val: 5}
// 				node6 := &ListNode{Val: 6}
// 				node7 := &ListNode{Val: 7}
//
// 				node5.Next = node6
// 				node6.Next = node7
//
// 				return node5
// 			}(),
// 			expected: false,
// 		},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(" ", func(t *testing.T) {
// 			result := hasCycle(tc.input)
// 			if result != tc.expected {
// 				t.Errorf("Expected %v, but got %v", tc.expected, result)
// 			}
// 		})
// 	}
// }

// func TestHasCycle(t *testing.T) {
// 	tests := []struct {
// 		input    *ListNode
// 		expected *ListNode
// 	}{
// 		{
// 			input: func() *ListNode {
// 				node1 := &ListNode{Val: 1}
// 				node2 := &ListNode{Val: 2}
// 				node3 := &ListNode{Val: 3}
// 				node4 := &ListNode{Val: 4}
//
// 				node1.Next = node2
// 				node2.Next = node3
// 				node3.Next = node4
// 				node4.Next = node2 // Creating a cycle by connecting the last node to the second node
//
// 				return node1
// 			}(),
// 			expected: true,
// 		},
// 		{
// 			input: func() *ListNode {
// 				node5 := &ListNode{Val: 5}
// 				node6 := &ListNode{Val: 6}
// 				node7 := &ListNode{Val: 7}
//
// 				node5.Next = node6
// 				node6.Next = node7
//
// 				return node5
// 			}(),
// 			expected: false,
// 		},
// 	}
//
// 	for _, tc := range tests {
// 		t.Run(" ", func(t *testing.T) {
// 			result := hasCycle(tc.input)
// 			if result != tc.expected {
// 				t.Errorf("Expected %v, but got %v", tc.expected, result)
// 			}
// 		})
// 	}
// }

// No test case for detect cycle

// No test case for detect Intersection

// func TestExcelSheetConv(t *testing.T) {
// 	tests := []struct {
// 		columnNumber int
// 		want         string
// 	}{
// 		{1, "A"},
// 		{28, "AB"},
// 		{701, "ZY"},
// 	}
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%v", tc.columnNumber), func(t *testing.T) {
// 			got := excelSheetConv(tc.columnNumber)
// 			if got != tc.want {
// 				t.Fatalf("got %v; want %v", got, tc.want)
// 			}
// 		})
// 	}
// }

// func TestExcelSheetConvNum(t *testing.T) {
// 	tests := []struct {
// 		columnTitle string
// 		want        int
// 	}{
// 		{"A", 1},
// 		{"AB", 28},
// 		{"ZY", 701},
// 	}
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%v", tc.columnTitle), func(t *testing.T) {
// 			got := excelSheetConvNum(tc.columnTitle)
// 			if got != tc.want {
// 				t.Fatalf("got %v; want %v", got, tc.want)
// 			}
// 		})
// 	}
//}

// func TestMajorityNumber(t *testing.T) {
// 	tests := []struct {
// 		nums []int
// 		want int
// 	}{
// 		{[]int{3, 2, 3}, 3},
// 		{[]int{1, 2, 1, 2, 3, 2, 2}, 2},
// 		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
// 	}
// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%v", tc.nums), func(t *testing.T) {
// 			got := majorityNumber(tc.nums)
// 			if got != tc.want {
// 				t.Fatalf("got %v, want %v", got, tc.want)
// 			}
// 		})
// 	}
// }

func TestReverseBit(t *testing.T) {
	tests := []struct {
		num  uint32
		want uint32
	}{
		{0b00000010100101000001111010011100, 0b00111001011110000010100101000000},
		{0b11111111111111111111111111111101, 3221225471},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.num), func(t *testing.T) {
			got := reverseBit(tc.num)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		num  uint32
		want int
	}{
		{0b00000000000000000000000000001011, 3},
		{0b00000000000000000000000010000000, 1},
		{0b11111111111111111111111111111101, 31},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.num), func(t *testing.T) {
			got := hammingWeight(tc.num)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsHappy(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{19, true},
		{2, false},
		{7, true},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.n), func(t *testing.T) {
			got := isHappy(tc.n)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestRemoveLinkedList(t *testing.T) {
	tests := []struct {
		head     *util.ListNode
		val      int
		expected *util.ListNode
	}{
		{
			head: &util.ListNode{
				Val: 1,
				Next: &util.ListNode{
					Val: 4, Next: &util.ListNode{Val: 3, Next: &util.ListNode{Val: 4}},
				},
			},
			val: 4,
			expected: &util.ListNode{
				Val:  1,
				Next: &util.ListNode{Val: 3},
			},
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			result := removeLinkedList(tc.head, tc.val)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestIsIsmorphics(t *testing.T) {
	tests := []struct {
		s, t string
		want bool
	}{
		{"bar", "foo", false},
		{"egg", "add", true},
		{"paper", "title", true},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v and %v", tc.s, tc.t), func(t *testing.T) {
			got := isIsomorphics(tc.s, tc.t)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		head *util.ListNode
		want *util.ListNode
	}{
		{
			head: &util.ListNode{
				Val: 1,
				Next: &util.ListNode{
					Val: 2, Next: &util.ListNode{Val: 3, Next: &util.ListNode{Val: 4}},
				},
			},
			want: &util.ListNode{
				Val: 4,
				Next: &util.ListNode{
					Val: 3, Next: &util.ListNode{Val: 2, Next: &util.ListNode{Val: 1}},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := reverseList(tc.head)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestContainsDuplicates(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := containsNearbyDuplicate(tc.nums, tc.k)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCountNodes(t *testing.T) {
	tests := []struct {
		root *util.TreeNode
		want int
	}{
		{(*util.TreeNode)(nil), 0},
		{&util.TreeNode{Val: 1, Left: &util.TreeNode{Val: 2}, Right: &util.TreeNode{Val: 3}}, 3},
	}

	for _, tc := range tests {
		t.Run(" ", func(t *testing.T) {
			got := countNodes(tc.root)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestInverseNode(t *testing.T) {
	tests := []struct {
		root *util.TreeNode
		want *util.TreeNode
	}{
		{(*util.TreeNode)(nil), nil},
		{
			&util.TreeNode{Val: 1, Left: &util.TreeNode{Val: 2}, Right: &util.TreeNode{Val: 3}},
			&util.TreeNode{Val: 1, Left: &util.TreeNode{Val: 3}, Right: &util.TreeNode{Val: 2}},
		},
	}

	for _, tc := range tests {
		t.Run(" ", func(t *testing.T) {
			got := inverseNode(tc.root)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		nums []int
		want []string
	}{
		{[]int{0, 1, 2, 5, 6, 9}, []string{"0->2", "5->6", "9"}},
		{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.nums), func(t *testing.T) {
			got := summaryRanges(tc.nums)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsPowerOfTwos(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{1, true},
		{16, true},
		{3, false},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.n), func(t *testing.T) {
			got := isPowerOfTwos(tc.n)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsPowerOfThrees(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{-1, false},
		{0, false},
		{27, true},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.n), func(t *testing.T) {
			got := isPowerOfThree(tc.n)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsPowerOfFour(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{1, true},
		{16, true},
		{3, false},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.n), func(t *testing.T) {
			got := isPowerOfFour(tc.n)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsLinkPalindrome(t *testing.T) {
	tests := []struct {
		head *util.ListNode
		want bool
	}{
		{
			head: &util.ListNode{
				Val: 1,
				Next: &util.ListNode{
					Val: 2,
					Next: &util.ListNode{
						Val: 2,
						Next: &util.ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			result := isLinkPalindrome(tc.head)
			if result != tc.want {
				t.Errorf("Expected %v, got %v", tc.want, result)
			}
		})
	}
}

func TestBinaryTreePaths(t *testing.T) {
	// Create a sample binary tree
	root := &util.TreeNode{
		Val: 1,
		Left: &util.TreeNode{
			Val: 2,
			Left: &util.TreeNode{
				Val: 5,
			},
			Right: &util.TreeNode{
				Val: 4,
			},
		},
		Right: &util.TreeNode{
			Val: 3,
		},
	}

	// Define the expected result
	expectedResult := []string{"1->2->5", "1->2->4", "1->3"}

	// Call the function
	result := binaryTreePaths(root)

	// Check if the result matches the expected result
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestAddDigits(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		{1, 1},
		{38, 2},
		{0, 0},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.num), func(t *testing.T) {
			got := addDigits(tc.num)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsUgly(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{6, true},
		{1, true},
		{14, false},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.n), func(t *testing.T) {
			got := isUgly(tc.n)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestMissingNum(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{0, 3, 2}, 1},
		{[]int{1, 3, 4, 5, 6, 7, 0}, 2},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.nums), func(t *testing.T) {
			got := missingNum(tc.nums)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{0, 3, 2, 0, 12}, []int{3, 2, 12, 0, 0}},
		{[]int{0}, []int{0}},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.nums), func(t *testing.T) {
			moveZeroes(tc.nums)
			if !reflect.DeepEqual(tc.nums, tc.want) {
				t.Fatalf("got %v, want %v", tc.nums, tc.want)
			}
		})
	}
}

func TestWordPattern(t *testing.T) {
	tests := []struct {
		pattern string
		s       string
		want    bool
	}{
		{"abba", "dog, cat, cat, dog", true},
		{"bbaa", "cat, dog, cat, dog", false},
	}

	for _, tc := range tests {
		t.Run(" ", func(t *testing.T) {
			got := wordPattern(tc.pattern, tc.s)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCountBits(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{2, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 1, 2}},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.n), func(t *testing.T) {
			got := countBits(tc.n)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestReversStrings(t *testing.T) {
	tests := []struct {
		s    []byte
		want []byte
	}{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{'H', 'a', 'p', 'p', 'y'}, []byte{'y', 'p', 'p', 'a', 'H'}},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.s), func(t *testing.T) {
			reverseStrings(tc.s)
			if !reflect.DeepEqual(tc.s, tc.want) {
				t.Fatalf("got %v, want %v", tc.s, tc.want)
			}
		})
	}
}

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"hello", "holle"},
		{"leetcode", "leotcede"},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.s), func(t *testing.T) {
			got := reverseVowels(tc.s)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{1, 2, 2, 1, 0, 3}, []int{1, 2, 3, 4}, []int{1, 2, 3}},
		{[]int{5, 4, 6, 9, 0, 2}, []int{5, 5, 5, 1, 8, 7}, []int{5}},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v, %v", tc.nums1, tc.nums2), func(t *testing.T) {
			got := intersection(tc.nums1, tc.nums2)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		// Assuming the function accounts for duplicates and the order of the elements in the result is not specified
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v, %v", tc.nums1, tc.nums2), func(t *testing.T) {
			got := intersect(tc.nums1, tc.nums2)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSquare(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{16, true},
		{14, false},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.n), func(t *testing.T) {
			got := validSquare(tc.n)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		ransomNote string
		magazine   string
		want       bool
	}{
		{"a", "b", false},
		{"aa", "aab", true},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v, %v", tc.ransomNote, tc.magazine), func(t *testing.T) {
			got := canConstruct(tc.ransomNote, tc.magazine)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestFirstUniqueChar(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"leetcode", 0},
		{"lovely", 1},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.s), func(t *testing.T) {
			got := firstUniqueChar(tc.s)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestFindTheDifference(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want byte
	}{
		{"abcd", "abcde", 'e'},
		{"", "y", 'y'},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v, %v", tc.s, tc.t), func(t *testing.T) {
			got := findTheDifference(tc.s, tc.t)
			if got != tc.want {
				t.Fatalf(" got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"abc", "ahbgnc", true},
		{"aed", "ayefrd", true},
		{"bfg", "tyfg", false},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v, %v", tc.s, tc.t), func(t *testing.T) {
			got := isSubsequence(tc.s, tc.t)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestReadBinaryWatch(t *testing.T) {
	tests := []struct {
		turnedOn int
		want     []string
	}{
		{
			1,
			[]string{
				"0:01",
				"0:02",
				"0:04",
				"0:08",
				"0:16",
				"0:32",
				"1:00",
				"2:00",
				"4:00",
				"8:00",
			},
		},
		{9, []string{}},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.turnedOn), func(t *testing.T) {
			got := readBinaryWatch(tc.turnedOn)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSumOfLeftLeaves(t *testing.T) {
	tests := []struct {
		root *util.TreeNode
		want int
	}{
		{(*util.TreeNode)(nil), 0},
		{&util.TreeNode{Val: 1, Left: &util.TreeNode{Val: 2}, Right: &util.TreeNode{Val: 3}}, 2},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.root), func(t *testing.T) {
			got := sumOfLeftLeaves(tc.root)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestToHex(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{-1, "ffffffff"},
		{26, "1a"},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.n), func(t *testing.T) {
			got := toHex(tc.n)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abccccdd", 7},
		{"a", 1},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.s), func(t *testing.T) {
			got := longestPalindrome(tc.s)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{3, []string{"1", "2", "Fizz"}},
		{5, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{
			15,
			[]string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
				"Fizz",
				"7",
				"8",
				"Fizz",
				"Buzz",
				"11",
				"Fizz",
				"13",
				"14",
				"FizzBuzz",
			},
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.n), func(t *testing.T) {
			got := fizzbuzz(tc.n)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestThirdMax(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 1, 4}, 2},
		{[]int{3, 2, 1}, 1},
		{[]int{1, 2}, 2},
		{[]int{2, 2, 3, 1}, 1},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.nums), func(t *testing.T) {
			got := thirdMax(tc.nums)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestAddString(t *testing.T) {
	tests := []struct {
		num1, num2 string
		want       string
	}{
		{"111", "123", "234"},
		{"456", "77", "533"},
		{"0", "0", "0"},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v, %v", tc.num1, tc.num2), func(t *testing.T) {
			got := addStrings(tc.num1, tc.num2)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCountSegment(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"we are a king in the hood", 7},
		{"yoo, can you introduce me to this king", 8},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.s), func(t *testing.T) {
			got := countSegments(tc.s)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestArrangeCoins(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{5, 2},
		{8, 3},
		{6, 3},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.n), func(t *testing.T) {
			got := arrangeCoins(tc.n)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestFindAllDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 1, 2}, []int{1}},
		{[]int{4, 3, 5, 6, 7, 3, 2, 1, 1}, []int{3, 1}},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.nums), func(t *testing.T) {
			got := findAllDuplicates(tc.nums)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestFindDisappearedNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
		{[]int{1, 1}, []int{2}},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.nums), func(t *testing.T) {
			got := findDisapperedNumber(tc.nums)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestAssignCookies(t *testing.T) {
	tests := []struct {
		g, s []int
		want int
	}{
		{[]int{1, 2, 3}, []int{1, 1}, 1},
		{[]int{1, 2}, []int{1, 2, 3}, 2},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v, %v", tc.g, tc.s), func(t *testing.T) {
			got := assignCookies(tc.g, tc.s)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestHammingDistance(t *testing.T) {
	tests := []struct {
		x, y int
		want int
	}{
		{1, 4, 2},
		{3, 1, 1},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v, %v", tc.x, tc.y), func(t *testing.T) {
			got := hammingDistance(tc.x, tc.y)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestRepeatedSubString(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"aba", false},
		{"abab", true},
		{"abcabcabcabc", true},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.s), func(t *testing.T) {
			got := reapeatedSubString(tc.s)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}
