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

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v %v", tc.nums, tc.target), func(t *testing.T) {
			got := searchInsert(tc.nums, tc.target)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

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

func TestFindComplement(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{5, 2},
		{1, 0},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.n), func(t *testing.T) {
			got := findComplement(tc.n)
			if got != tc.want {
				t.Fatalf(" got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestLicenseFormatting(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want string
	}{
		{"5F3Z-2e-9-w", 4, "5F3Z-2E9W"},
		{"2-5g-3-J", 2, "2-5G-3J"},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v %v", tc.s, tc.k), func(t *testing.T) {
			got := licenseFormatting(tc.s, tc.k)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestConsecutivesOnes(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 0, 1, 1, 0, 1, 1, 1}, 3},
		{[]int{1, 0, 1, 1, 0, 1}, 2},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.nums), func(t *testing.T) {
			got := findMaxConsecutivesOne(tc.nums)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestConstructRectangle(t *testing.T) {
	tests := []struct {
		area int
		want []int
	}{
		{4, []int{2, 2}},
		{37, []int{37, 1}},
		{122122, []int{427, 286}},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.area), func(t *testing.T) {
			got := construtRectangle(tc.area)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestFindPoisonedDuration(t *testing.T) {
	tests := []struct {
		timeSeries []int
		duration   int
		want       int
	}{
		{[]int{1, 4}, 2, 4},
		{[]int{1, 2}, 2, 3},
		{[]int{4, 5}, 3, 4},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v, %v", tc.timeSeries, tc.duration), func(t *testing.T) {
			got := findPoisonedDuration(tc.timeSeries, tc.duration)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestNextGreaterElement(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{4, 1, 2}, []int{1, 3, 4, 2}, []int{-1, 3, -1}},
		{[]int{2, 4}, []int{1, 2, 3, 4}, []int{3, -1}},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v, %v", tc.nums1, tc.nums2), func(t *testing.T) {
			got := nextGreaterElement(tc.nums1, tc.nums2)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestFindWords(t *testing.T) {
	tests := []struct {
		words []string
		want  []string
	}{
		{[]string{"Hello", "Alaska", "Dad", "Peace"}, []string{"Alaska", "Dad"}},
		{[]string{"omk"}, []string{}},
		{[]string{"adsdf", "sfd"}, []string{"adsdf", "sfd"}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.words), func(t *testing.T) {
			got := findWords(tt.words)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		root *util.TreeNode
		want []int
	}{
		{&util.TreeNode{Val: 0}, []int{0}},
		{
			&util.TreeNode{Val: 1, Left: &util.TreeNode{Val: 2}, Right: &util.TreeNode{Val: 2}},
			[]int{2},
		},
	}

	for _, tc := range tests {
		t.Run(" ", func(t *testing.T) {
			got := findMode(tc.root)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestBase7(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{100, "202"},
		{-7, "-10"},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.n), func(t *testing.T) {
			got := base7(tc.n)
			if got != tc.want {
				t.Fatalf(" got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestFindRelativeRank(t *testing.T) {
	tests := []struct {
		score []int
		want  []string
	}{
		{[]int{5, 4, 3, 2, 1}, []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}},
		{[]int{10, 3, 8, 9, 4}, []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.score), func(t *testing.T) {
			got := findRelativeRank(tt.score)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerfectNum(t *testing.T) {
	tests := []struct {
		num  int
		want bool
	}{
		{6, true},
		{28, true},
		{4, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.num), func(t *testing.T) {
			got := perfectNum(tt.num)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("n=%v", tc.n), func(t *testing.T) {
			got := fibonacci(tc.n)
			if got != tc.want {
				t.Fatalf("Fibonacci() = %v; want %v", got, tc.want)
			}
		})
	}
}

func TestDetectCapitalUse(t *testing.T) {
	tests := []struct {
		S    string
		want bool
	}{
		{"USA", true},
		{"FlaG", false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.S), func(t *testing.T) {
			got := detectCapitalUse(tt.S)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseStr(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want string
	}{
		{"abcdefg", 2, "bacdfeg"},
		{"abcd", 2, "bacd"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.s, tt.k), func(t *testing.T) {
			got := reverseStr(tt.s, tt.k)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindLusLength(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want int
	}{
		{"aba", "cdc", 3},
		{"aaa", "bbb", 3},
		{"aaa", "aaa", -1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.a, tt.b), func(t *testing.T) {
			got := findLusLength(tt.a, tt.b)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckRecords(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"PPALLP", true},
		{"PPALLL", false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.s), func(t *testing.T) {
			got := checkRecords(tt.s)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseWords(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
		{"Mr Ding", "rM gniD"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.s), func(t *testing.T) {
			got := reverseWords(tt.s)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayPairs(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 4, 3, 2}, 4},
		{[]int{6, 2, 6, 5, 1, 2}, 9},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := arrayPairSum(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrixReshape(t *testing.T) {
	tests := []struct {
		mat  [][]int
		r, c int
		want [][]int
	}{
		{[][]int{{1, 2}, {3, 4}}, 1, 4, [][]int{{1, 2, 3, 4}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v, %v", tt.mat, tt.c, tt.r), func(t *testing.T) {
			got := matrixReshape(tt.mat, tt.r, tt.c)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistributeCandy(t *testing.T) {
	tests := []struct {
		candyType []int
		want      int
	}{
		{[]int{1, 1, 2, 2, 3, 3}, 3},
		{[]int{1, 1, 2, 3}, 2},
		{[]int{6, 6, 6, 6}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.candyType), func(t *testing.T) {
			got := distributeCandy(tt.candyType)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindLHS(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 2, 2, 2, 3, 5, 7}, 5},
		{[]int{6, 6, 6, 6}, 0},
		{[]int{1, 2, 3, 4}, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := findLHS(tt.nums)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxCount(t *testing.T) {
	tests := []struct {
		m, n int
		ops  [][]int
		want int
	}{
		{3, 3, [][]int{{2, 2}, {3, 3}}, 4},
		{3, 3, [][]int{}, 9},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v, %v", tt.m, tt.n, tt.ops), func(t *testing.T) {
			got := maxCount(tt.m, tt.n, tt.ops)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindRestaurant(t *testing.T) {
	tests := []struct {
		list1, list2 []string
		want         []string
	}{
		{
			[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
			[]string{"Shogun"},
		},
		{
			[]string{"happy", "sad", "good"},
			[]string{"sad", "happy", "good"},
			[]string{"sad", "happy"},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.list1, tt.list2), func(t *testing.T) {
			got := findRestaurant(tt.list1, tt.list2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf(" got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanPlaceFlower(t *testing.T) {
	tests := []struct {
		flowerBed []int
		n         int
		want      bool
	}{
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.flowerBed, tt.n), func(t *testing.T) {
			got := canPlaceFlower(tt.flowerBed, tt.n)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaximumProducts(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{1, 2, 3, 4}, 24},
		{[]int{-1, -2, -3}, -6},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := maximumProducts(tt.nums)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMaxAverarge(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75000},
		{[]int{5}, 1, 5.0000},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.nums, tt.k), func(t *testing.T) {
			got := findMaxAverage(tt.nums, tt.k)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindErrorNums(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 2, 4}, []int{2, 3}},
		{[]int{1, 1}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := findErrorNums(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJudgeCircle(t *testing.T) {
	tests := []struct {
		moves string
		want  bool
	}{
		{"UD", true},
		{"LL", false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.moves), func(t *testing.T) {
			got := judgeCirce(tt.moves)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindLengthOfLcis(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 5, 4, 7}, 3},
		{[]int{2, 2, 2, 2, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := findLengthOfLCIS(tt.nums)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"aba", true},
		{"abca", true},
		{"abc", false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.s), func(t *testing.T) {
			got := validPalindrome(tt.s)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalPoints(t *testing.T) {
	tests := []struct {
		operations []string
		want       int
	}{
		{[]string{"5", "2", "C", "D", "+"}, 30},
		{[]string{"5", "-2", "4", "C", "D", "9", "+", "+"}, 27},
		{[]string{"1", "C"}, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.operations), func(t *testing.T) {
			got := calPoints(tt.operations)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlternatingBits(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{5, true},
		{7, false},
		{11, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.n), func(t *testing.T) {
			got := hasAlternativesBit(tt.n)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"00110011", 6},
		{"10101", 4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.s), func(t *testing.T) {
			got := countBinarySubstring(tt.s)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindShortestSubArray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 2, 3, 1}, 2},
		{[]int{1, 2, 2, 3, 1, 4, 2}, 6},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := findShortestSubarray(tt.nums)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToLower(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"Hello", "hello"},
		{"here", "here"},
		{"LOVELY", "lovely"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.s), func(t *testing.T) {
			got := toLower(tt.s)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsOneBitsCharacter(t *testing.T) {
	tests := []struct {
		bits []int
		want bool
	}{
		{[]int{1, 0, 0}, true},
		{[]int{1, 1, 1, 0}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.bits), func(t *testing.T) {
			got := isOneBitsCharacter(tt.bits)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
		{[]int{2, 1, -1}, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := pivotIndex(tt.nums)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelfDividingNums(t *testing.T) {
	tests := []struct {
		left, right int
		want        []int
	}{
		{1, 22, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}},
		{47, 85, []int{48, 55, 66, 77}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.left, tt.right), func(t *testing.T) {
			got := selfDividingNumbers(tt.left, tt.right)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNextGreatestLetters(t *testing.T) {
	tests := []struct {
		letter []byte
		target byte
		want   byte
	}{
		{[]byte{'c', 'f', 'j'}, 'a', 'c'},
		{[]byte{'x', 'x', 'y', 'y'}, 'z', 'x'},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.letter, tt.target), func(t *testing.T) {
			got := nextGreatestLetter(tt.letter, tt.target)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		cost []int
		want int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.cost), func(t *testing.T) {
			got := minCostClimbingStairs(tt.cost)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDominantIndex(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 6, 2, 1}, 1},
		{[]int{1, 2, 3, 4}, -1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := dominantIndex(tt.nums)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountPrimesSetBits(t *testing.T) {
	tests := []struct {
		left, right int
		want        int
	}{
		{6, 10, 4},
		{10, 15, 5},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.left, tt.right), func(t *testing.T) {
			got := countPrimeSetBits(tt.left, tt.right)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumJewelInStone(t *testing.T) {
	tests := []struct {
		jewels, stones string
		want           int
	}{
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.jewels, tt.stones), func(t *testing.T) {
			got := numJewelInStone(tt.jewels, tt.stones)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotateString(t *testing.T) {
	tests := []struct {
		s, goal string
		want    bool
	}{
		{"abcde", "cdeab", true},
		{"abcde", "abced", false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.goal, tt.s), func(t *testing.T) {
			got := rotateString(tt.s, tt.goal)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueMorseRepresentation(t *testing.T) {
	tests := []struct {
		words []string
		want  int
	}{
		{[]string{"gin", "zen", "gig", "msg"}, 2},
		{[]string{"a"}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.words), func(t *testing.T) {
			got := uniqueMorseRepresentation(tt.words)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumberOfLines(t *testing.T) {
	tests := []struct {
		width []int
		s     string
		want  []int
	}{
		{
			[]int{
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
			},
			"abcdefghijklmnopqrstuvwxyz",
			[]int{3, 60},
		},
		{
			[]int{
				4,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
				10,
			},
			"bbbcccdddaaa",
			[]int{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.s, tt.width), func(t *testing.T) {
			got := numberOfLines(tt.width, tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMostCommon(t *testing.T) {
	tests := []struct {
		paragraph string
		banned    []string
		want      string
	}{
		{"Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}, "ball"},
		{"a.", []string{}, "a"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.paragraph, tt.banned), func(t *testing.T) {
			got := mostCommonWord(tt.paragraph, tt.banned)
			if got != tt.want {
				t.Fatalf("got  %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortestToChar(t *testing.T) {
	tests := []struct {
		s    string
		c    byte
		want []int
	}{
		{"loveleetcode", 'e', []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}},
		{"aaab", 'b', []int{3, 2, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.s, tt.c), func(t *testing.T) {
			got := shortestToChar(tt.s, tt.c)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToGoatLatin(t *testing.T) {
	tests := []struct {
		sentence string
		want     string
	}{
		{"I speak Goat Latin", "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"},
		{
			"The quick brown fox jumped over the lazy dog",
			"heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.sentence), func(t *testing.T) {
			got := toGoatLatin(tt.sentence)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLargeGroupPosition(t *testing.T) {
	tests := []struct {
		s    string
		want [][]int
	}{
		{"abbxxxxzzy", [][]int{{3, 6}}},
		{"abc", [][]int{}},
		{"abcdddeeeeaabbbcd", [][]int{{3, 5}, {6, 9}, {12, 14}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.s), func(t *testing.T) {
			got := largeGroupPosition(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlipAndInvertImage(t *testing.T) {
	tests := []struct {
		image [][]int
		want  [][]int
	}{
		{[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}, [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
		{
			[][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
			[][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.image), func(t *testing.T) {
			got := flipAndInvertImage(tt.image)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRectangleOverlap(t *testing.T) {
	tests := []struct {
		rec1, rec2 []int
		want       bool
	}{
		{[]int{0, 0, 2, 2}, []int{1, 1, 3, 3}, true},
		{[]int{0, 0, 1, 1}, []int{1, 0, 2, 1}, false},
		{[]int{0, 0, 1, 1}, []int{2, 2, 3, 3}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.rec1, tt.rec2), func(t *testing.T) {
			got := isRectangleOverlap(tt.rec1, tt.rec2)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeArea(t *testing.T) {
	tests := []struct {
		a, b, c, d, e, f, g, h int
		want                   int
	}{
		{-3, 0, 3, 4, 0, -1, 9, 2, 45},
		{-2, -2, 2, 2, -2, -2, 2, 2, 16},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("the input can be very long you know"), func(t *testing.T) {
			got := computeArea(tt.a, tt.b, tt.c, tt.d, tt.e, tt.f, tt.g, tt.h)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBackspaceCompare(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"ab#d", "ae#d", true},
		{"a#c", "b", false},
		{"ab##", "c#d#", true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.s, tt.t), func(t *testing.T) {
			got := backspaceCompare(tt.s, tt.t)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLemonadeChange(t *testing.T) {
	tests := []struct {
		bills []int
		want  bool
	}{
		{[]int{5, 5, 5, 10, 20}, true},
		{[]int{5, 5, 10, 10, 10, 20}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.bills), func(t *testing.T) {
			got := lemonadeChange(tt.bills)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransposeMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   [][]int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, [][]int{{1, 4}, {2, 5}, {3, 6}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.matrix), func(t *testing.T) {
			got := transposeMatrix(tt.matrix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryGap(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{22, 2},
		{8, 0},
		{5, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.n), func(t *testing.T) {
			got := binaryGap(tt.n)
			if got != tt.want {
				t.Fatalf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestBodyStrings(t *testing.T) {
	tests := []struct {
		s, goal string
		want    bool
	}{
		{"ab", "ba", true},
		{"ab", "ab", false},
		{"aa", "aa", true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.s, tt.goal), func(t *testing.T) {
			got := buddyStrings(tt.s, tt.goal)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectionArea(t *testing.T) {
	test := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{1, 2}, {3, 4}}, 17},
		{[][]int{{2}}, 5},
		{[][]int{{1, 0}, {0, 2}}, 8},
	}
	for _, tt := range test {
		t.Run(fmt.Sprintf("%v", tt.grid), func(t *testing.T) {
			got := projectionArea(tt.grid)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
