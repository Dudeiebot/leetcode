package medium

import (
	"slices"

	"leetCode/graph"
	util "leetCode/struct"
)

func GroupAnagram(strs []string) [][]string {
	hmap := map[[26]int][]string{}
	res := [][]string{}

	for _, word := range strs {
		k := [26]int{}
		for _, char := range word {
			k[char-'a']++
		}
		hmap[k] = append(hmap[k], word)
	}
	for _, val := range hmap {
		res = append(res, val)
	}
	return res
}

func topKFrequent(nums []int, k int) []int {
	seen := make(map[int]int)
	res := make([]int, 0, k)

	for _, tt := range nums {
		seen[tt]++
	}

	freq := make([][]int, len(nums)+1)
	for mk, mv := range seen {
		freq[mv] = append(freq[mv], mk)
	}

	for i := len(freq) - 1; i >= 0; i-- {
		// if freq[i] != nil {
		// 	res = append(res, freq[i]...)
		// 	if len(res) >= k {
		// 		res = res[:k]
		// 		break
		// 	}
		// }

		// another way to do it
		for _, val := range freq[i] {
			if k > 0 {
				res = append(res, val)
				k--
			}
		}
	}
	return res
}

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		res[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= suffix
		suffix *= nums[i]
	}
	return res
}

func longestConsecutive(nums []int) int {
	seen := make(map[int]bool)
	maxLen := 0

	for _, n := range nums {
		seen[n] = true
	}

	for _, n := range nums {
		if !seen[n-1] {
			currLen := 1
			curr := n

			for seen[curr+1] {
				curr++
				currLen++
			}

			if maxLen < currLen {
				maxLen = currLen
			}
		}
	}
	return maxLen
}

func twoSumII(nums []int, target int) []int {
	l, r := 0, len(nums)-1

	for l < r {
		if nums[l]+nums[r] > target {
			r--
		} else if nums[l]+nums[r] < target {
			l++
		} else {
			return []int{l + 1, r + 1}
		}
	}
	return nil
}

func threeSum(nums []int) [][]int {
	res := [][]int{}

	slices.Sort(nums)
	for i, n := range nums {
		if i > 0 && n == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			threeSum := n + nums[l] + nums[r]
			if threeSum > 0 {
				r--
			} else if threeSum < 0 {
				l++
			} else {
				res = append(res, []int{n, nums[l], nums[r]})
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}
	}
	return res
}

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0

	for l < r {
		area := (r - l) * min(height[l], height[r])
		res = max(res, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}

func lengthOfLongestSubString(s string) int {
	seen := make(map[byte]int)
	l, r, res := 0, 0, 0

	for r < len(s) {
		seen[s[r]]++
		for seen[s[r]] > 1 {
			seen[s[l]]--
			l++
		}
		res = max(res, r-l+1)
		r++
	}
	return res
}

func characterReplacement(s string, k int) int {
	res := 0
	maxF := 0
	seen := make(map[byte]int)
	l := 0

	for r := 0; r < len(s); r++ {
		seen[s[r]]++
		maxF = max(maxF, seen[s[r]])
		if (r-l+1)-maxF > k {
			seen[s[l]]--
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}

func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := l + (r-l)/2
		if nums[l] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		}

		if nums[l] <= nums[m] {
			if target >= nums[l] && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if target > nums[m] && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return -1
}

func reOrderList(head *util.ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	var prev *util.ListNode
	for slow != nil {
		slow.Next, prev, slow = prev, slow, slow.Next
	}

	first := head
	for prev.Next != nil {
		first.Next, first = prev, first.Next
		prev.Next, prev = first, prev.Next
	}
}

func removeNthNode(head *util.ListNode, n int) *util.ListNode {
	dummy := &util.ListNode{Next: head}

	slow, fast := dummy, dummy

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func lowestCommonAncestor(root, p, q *util.TreeNode) *util.TreeNode {
	// recursion solution that is more efficient
	// if p.Val < root.Val && q.Val < root.Val {
	// 	return lowestCommonAncestor(root.Left, p, q)
	// } else if p.Val > root.Val && q.Val > root.Val {
	// 	return lowestCommonAncestor(root.Right, p, q)
	// }
	// return root

	// iterative solution that is less efficient
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}

func levelOrder(root *util.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	q := ([]*util.TreeNode{root}) // Initialize the queue with the root node.

	for len(q) > 0 { // Continue processing while there are nodes in the queue.
		curSize := len(q) // We need this because it is the number of nodes at the current level.
		var level []int   // Slice to store the values of nodes at the current level.

		for i := 0; i < curSize; i++ { // current level processing
			node := q[0]
			q := q[1:]                      // removing the first node
			level = append(level, node.Val) // Append the node's value to the current level slice.

			if node.Left != nil { // If the left child exists, add it to the queue.
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

func isValidBST(root *util.TreeNode) bool {
	var isValid func(root, min, max *util.TreeNode) bool

	isValid = func(root, min, max *util.TreeNode) bool {
		if root == nil {
			return true
		}

		if min != nil && root.Val <= min.Val {
			return false
		}
		if max != nil && root.Val >= max.Val {
			return false
		}
		return isValid(root.Left, min, root) && isValid(root.Right, root, max)
	}
	return isValid(root, nil, nil)
}

func kthSmallest(root *util.TreeNode, k int) int {
	stack := []*util.TreeNode{}
	curr := root
	x := 1
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		if x == k {
			return curr.Val
		}
		x++
		stack = stack[:len(stack)-1]
		curr = curr.Right
	}
	return -1
}

func buildTree(preorder, inorder []int) *util.TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	i := slices.Index(inorder, preorder[0])
	lpre := preorder[1 : i+1]
	rpre := preorder[i+1:]
	lin := inorder[:i]
	rin := inorder[i+1:]

	return &util.TreeNode{
		Val:   i,
		Left:  buildTree(lpre, lin),
		Right: buildTree(rpre, rin),
	}
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	temp := []int{}

	var dfs func(i, curr int, temp []int)
	dfs = func(i, curr int, temp []int) {
		if curr == target {
			res = append(res, append([]int{}, temp...))
		}
		if curr > target {
			return
		}
		for j := i; j < len(candidates); j++ {
			temp = append(temp, candidates[j])
			dfs(j, curr+candidates[j], temp)
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0, 0, temp)
	return res
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	islands := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				// choose either IslandBfs and IslandDfs
				// graph.IslandDfs(r, c, grid)
				graph.IslandBfs(grid, r, c)
				islands++
			}
		}
	}
	return islands
}
