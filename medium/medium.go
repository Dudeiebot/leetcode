package medium

import (
	"slices"
	"sort"
	"strings"

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

func cloneGraph(node *util.Node) *util.Node {
	if node == nil {
		return nil
	}

	copies := make(map[*util.Node]*util.Node)
	var dfs func(n *util.Node)

	dfs = func(n *util.Node) {
		if n == nil {
			return
		}

		if _, exists := copies[n]; exists {
			return
		}
		newNode := &util.Node{Val: n.Val}
		copies[n] = newNode

		for _, child := range n.Children {
			dfs(child)
			newNode.Children = append(newNode.Children, copies[child])
		}
	}
	dfs(node)
	return copies[node]
}

func pacificAtlantic(heights [][]int) [][]int {
	res := make([][]int, 0)

	if len(heights) == 0 || len(heights[0]) == 0 {
		return res
	}

	rows := len(heights)
	cols := len(heights[0])

	pac := make([][]bool, rows)
	atl := make([][]bool, rows)

	for r := 0; r < rows; r++ {
		pac[r] = make([]bool, cols)
		atl[r] = make([]bool, cols)
	}

	var dfs func(r, c int, visit [][]bool, prevHeight int)
	dfs = func(r, c int, visit [][]bool, prevHeight int) {
		if r < 0 || c < 0 || r == rows || c == cols || heights[r][c] < prevHeight || visit[r][c] {
			return
		}
		visit[r][c] = true
		dfs(r+1, c, visit, heights[r][c])
		dfs(r-1, c, visit, heights[r][c])
		dfs(r, c+1, visit, heights[r][c])
		dfs(r, c-1, visit, heights[r][c])
	}

	for c := 0; c < cols; c++ {
		dfs(0, c, pac, heights[0][c])
		dfs(rows-1, c, atl, heights[rows-1][c])
	}

	for r := 0; r < rows; r++ {
		dfs(r, 0, pac, heights[r][0])
		dfs(r, cols-1, atl, heights[r][cols-1])
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if pac[r][c] && atl[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}
	return res
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	visited := make([]bool, numCourses)
	visiting := make([]bool, numCourses)

	var dfs func(course int) bool

	dfs = func(course int) bool {
		if visited[course] {
			return true
		}

		if visiting[course] {
			return false
		}

		visiting[course] = true
		for _, p := range prerequisites {
			if p[0] == course {
				if !dfs(p[1]) {
					return false
				}
			}
		}
		visiting[course] = false
		visited[course] = true
		return true
	}
	for i := 0; i < numCourses; i++ {
		if !visited[i] {
			if !dfs(i) {
				return false
			}
		}
	}
	return true
}

// Dynamic Programming they can be very silly and get harder with time also
func rob(nums []int) int {
	rob1 := 0
	rob2 := 0

	for _, n := range nums {
		temp := max(n+rob1, rob2)
		rob1 = rob2
		rob2 = temp
	}
	return rob2
}

func robII(nums []int) int {
	var helper func(nums []int) int

	helper = func(nums []int) int {
		h1, h2 := 0, 0
		for _, n := range nums {
			maxM := max(h1+n, h2)
			h1 = h2
			h2 = maxM
		}
		return h2
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(helper(nums[1:]), helper(nums[:len(nums)-1]))
}

func longestPalindrome(s string) string {
	res := ""

	for i := 0; i < len(s); i++ {
		// odd length palindrome
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if (r - l + 1) > len(res) {
				res = s[l : r+1]
			}
			l--
			r++
		}
		l, r = i, i+1
		// even length palindrome
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if (r - l + 1) > len(res) {
				res = s[l : r+1]
			}
			l--
			r++
		}
	}
	return res
}

func palindromicSubstring(s string) int {
	res := 0

	for i := 0; i < len(s); i++ {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			res++
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			res++
			l--
			r++
		}
	}
	return res
}

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	n := len(s)
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		// Check one-digit number
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}

		// Check two-digit number
		twoDigit := s[i-2 : i]
		if twoDigit >= "10" && twoDigit <= "26" {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 // Initialize with a value larger than possible
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], 1+dp[i-coin])
			}
		}
	}

	if dp[amount] > amount {
		return -1 // If no solution is possible
	}
	return dp[amount]
}

func coinChangeII(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}

	return dp[amount]
}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSoFar := nums[0]
	minSoFar := nums[0]
	result := maxSoFar

	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		tempMax := max(curr, max(maxSoFar*curr, minSoFar*curr))
		minSoFar = min(curr, min(maxSoFar*curr, minSoFar*curr))

		maxSoFar = tempMax

		result = max(result, maxSoFar)
	}

	return result
}

func wordBreak(s string, wordDict []string) bool {
	q := []string{s}
	memo := make(map[string]bool)
	for len(q) != 0 {
		remaining := q[0]
		q = q[1:] // Dequeue
		if _, ok := memo[remaining]; ok {
			continue
		}
		for _, word := range wordDict {
			if word == remaining {
				return true
			}
			if strings.HasPrefix(remaining, word) {
				q = append(q, remaining[len(word):]) // Enqueue
				memo[remaining] = true
			}
		}
	}
	return false
}

// length of longest increasing subsequence
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] > res[len(res)-1] {
			res = append(res, nums[i])
		} else {
			// Find the first element in res that is >= nums[i]
			j := 0
			for j < len(res) && res[j] < nums[i] {
				j++
			}
			res[j] = nums[i]
		}
	}
	return len(res)
}

// the dp approach
func lengthOfLISDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Initialize dp array with 1s
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	// Compute LIS
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return slices.Max(dp)
}

// unique Paths
// combanitorial or mathematical approach
func uniquePaths(m int, n int) int {
	ans := 1
	for i := 1; i <= m-1; i++ {
		ans = ans * (n - 1 + i) / i
	}
	return ans
}

func uniquePathsDp(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func longestCommonSubsequence(text1, text2 string) int {
	dp := [][]int{}

	for i := 0; i < len(text1)+1; i++ {
		dp = append(dp, make([]int, len(text2)+1))
	}
	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

func maxSubArray(nums []int) int {
	max_end := nums[0]
	max_at_certain := nums[0]

	for i := 1; i < len(nums); i++ {
		max_at_certain = max(max_at_certain+nums[i], nums[i])
		max_end = max(max_end, max_at_certain)
	}
	return max_end
}

func canJump(nums []int) bool {
	reach := 0
	for i := 0; i <= reach; i++ {
		if reach < i+nums[i] {
			reach = i + nums[i]
		}
		if reach >= len(nums)-1 {
			return true
		}
	}
	return false
}

func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}

	for i := 0; i < len(intervals); i++ {
		if newInterval[1] < intervals[i][0] {
			res = append(res, newInterval)
			return append(res, intervals[i:]...)
		} else if newInterval[0] > intervals[i][1] {
			res = append(res, intervals[i])
		} else {
			newInterval = []int{
				min(newInterval[0], intervals[i][0]),
				max(newInterval[1], intervals[i][1]),
			}
		}
	}
	res = append(res, newInterval)
	return res
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	current := intervals[0]

	for j := range intervals {
		if intervals[j][0] > current[1] {
			res = append(res, current)
			current = intervals[j]
		}
		current[1] = max(current[1], intervals[j][1])
	}
	res = append(res, current)
	return res
}
