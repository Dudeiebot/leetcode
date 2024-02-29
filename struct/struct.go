package util

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
