package main

// import (
// 	"strings"
// 	"unicode"
// "math/big"
// )

// duplicate num in a array

// func DupNum(nums []int) bool {
// 	s := make(map[int]bool)
//
// 	for _, num := range nums {
// 		if s[num] {
// 			return true
// 		}
// 		s[num] = true
// 	}
// 	return false
// }

// func DupNum(nums []int) bool {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i] == nums[j] {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// valid anagram
// you can first of all do it by spliting(strings.split), sort it (sort.strings) and then compare/join with (strings.join) or
// func isAnagram(s, t string) bool {
// 	// make our map
// 	c_s := make(map[rune]int)
// 	c_t := make(map[rune]int)
//
// 	// add all the character individually to the map
// 	for _, char := range s {
// 		c_s[char]++
// 	}
//
// 	for _, char := range t {
// 		c_t[char]++
// 	}
//
// 	// range through our map and check if the count of c_s matches the one of c_t
// 	for char, count := range c_s {
// 		if count != c_t[char] {
// 			return false
// 		}
// 	}
// 	return true
// }

// valid Palindrome
// before i wanted to use the rune datatype reverse model but it takes a lot of memory and it is not fast
// so these is me comparing the first and the last and going through it in a single for loop statement
// func IsPalindrome(s string) bool {
// 	s = strings.Map(func(r rune) rune {
// 		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
// 			return -1
// 		}
// 		return unicode.ToLower(r)
// 	}, s)
//
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		if s[i] != s[j] {
// 			return false
// 		}
// 	}
// 	return true
// }

// Single Number in an array
// These is not memory safe and i am just putting these here for the normal understanding of how it works

//func IsSingleNos(nums []int) int {
// 	// res := 0
// 	seen := make(map[int]int)
//
// 	for _, num := range nums {
// 		seen[num]++
// 	}
//
// 	for num, count := range seen {
// 		if count <= 1 {
// 			return num
// 		}
// 	}
// 	return 0
// }

/* The bitwise XOR operation (^) has the property that when applied twice with the same value, it returns the original value. Therefore, when XORing all elements in the slice, the elements that appear an odd number of times will contribute to the final result, and those that appear an even number of times will cancel out.
 */
//The memory safe one

// func IsSingleNos(nums []int) int {
// 	res := 0
//
// 	for _, num := range nums {
// 		res = res ^ num
// 	}
// 	return res
// }

// We are finding the sum of 2 number in an array that forms a target and we return the indices of the numbers
// constant space of time technique O(n)
// func twoSums(nums []int, target int) []int {
// 	seen := make(map[int]int)
//
// 	for i, num := range nums {
// 		diff := target - num
// 		if j, ok := seen[diff]; ok {
// 			return []int{j, i}
// 		}
// 		seen[num] = i
// 	}
// 	return nil
// }
//
// bruteforce techniques O(n2)
// func twoSums(nums []int, target int) []int {
// 	arr := make([]int, 2)
//
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				arr[0] = i
// 				arr[1] = j
// 				return arr
// 			}
// 		}
// 	}
// 	return nil
// }

// nos Palindrome but we dont first convert the nos to string
// func isNosPalindrome(x int) bool {
// 	if x < 0 || (x != 0 && x%10 == 0) {
// 		return false
// 	}
//
// 	temph := 0
// 	for x > temph {
// 		temph = (temph * 10) + (x % 10)
// 		x = x / 10
// 	}
// 	return x == temph || x == temph/10
// }

// func rom2Int(s string) int {
// 	Numeral := map[byte]int{
// 		'I': 1,
// 		'V': 5,
// 		'X': 10,
// 		'L': 50,
// 		'C': 100,
// 		'D': 500,
// 		'M': 1000,
// 	}
//
// 	res := 0
// 	for i := 0; i < len(s)-1; i++ {
// 		if Numeral[s[i]] < Numeral[s[i+1]] {
// 			res -= Numeral[s[i]]
// 		} else {
// 			res += Numeral[s[i]]
// 		}
// 	}
// 	return res + Numeral[s[len(s)-1]]
// }

//this approach have a twist to it and it is not straightforward because you have to be raplacing about
// func rom2Int(s string) int {
// 	Numeral := map[string]int{
// 		"I": 1,
// 		"V": 5,
// 		"X": 10,
// 		"L": 50,
// 		"C": 100,
// 		"D": 500,
// 		"M": 1000,
// 	}
//
// 	res := 0
// 	s = strings.Replace(s, "IV", "IIII", -1)
// 	s = strings.Replace(s, "IX", "VIIII", -1)
// 	s = strings.Replace(s, "XL", "XXXX", -1)
// 	s = strings.Replace(s, "XC", "LXXXX", -1)
// 	s = strings.Replace(s, "CD", "CCCC", -1)
// 	s = strings.Replace(s, "CM", "DCCCC", -1)
//
// 	for _, i := range s {
// 		res += Numeral[string(i)]
// 	}
// 	return res
// }

// func longestPrefix(s []string) string {
// 	if len(s) == 0 {
// 		return ""
// 	}
//
// 	firstLet := s[0]
//
// 	for i := 0; i < len(firstLet); i++ {
// 		for _, words := range s[1:] {
// 			if i == len(words) || words[i] != firstLet[i] {
// 				return firstLet[0:i]
// 			}
// 		}
// 	}
// 	return firstLet
// }

// func checkParenthesis(s string) bool {
// 	stack := make([]rune, 0)
//
// 	pairs := map[rune]rune{
// 		'[': ']',
// 		'{': '}',
// 		'(': ')',
// 	}
//
// 	for _, i := range s {
// 		// we can use the switch case statement also here
// 		if i == '[' || i == '{' || i == '(' {
// 			stack = append(stack, i) // we append to the stack
// 		} else if i == ']' || i == '}' || i == ')' && len(stack) > 0 {
// 			popped := stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]
// 			if pairs[popped] != i { // we pop from the stack and compare it with the else if (i)
// 				return false
// 			}
// 		}
// 	}
// 	return len(stack) == 0
// }

// Merge two sorted linked list
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }
//
// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	tempNode := &ListNode{}
// 	currentNode := tempNode
//
// 	for list1 != nil && list2 != nil {
// 		if list1.Val < list2.Val {
// 			currentNode.Next = list1 // this point to the currentvalue in our tempNode
// 			list1 = list1.Next       // this point to our NextValue
// 		} else {
// 			currentNode.Next = list2
// 			list2 = list2.Next
// 		}
// 		currentNode = currentNode.Next
// 	}
//
// 	// and when any of them come to finish here, we add the remaining to currentNode(which also lead to our tempNode)
// 	if list1 != nil {
// 		currentNode.Next = list1
// 	} else {
// 		currentNode.Next = list2
// 	}
// 	return tempNode.Next
// }

// here we are using a pointer, and the pointer is used to compare to the num
// func removeDuplicates(nums []int) int {
// 	i := 0
//
// 	for k := 0; k < len(nums); k++ {
// 		if nums[k] != nums[i] {
// 			i++
// 			nums[i] = nums[k]
// 		}
// 	}
// 	return i + 1
// }
//
// // same thing here also the pointer is necesay for constant space of time
// func removeElements(nums []int, val int) int {
// 	i := 0
// 	for _, num := range nums {
// 		if num != val {
// 			nums[i] = num
// 			i++
// 		}
// 	}
// 	return i
// }

// func strstr(mainString string, substring string) int {
// 	// this is fast way
// 	// return strings.Index(mainString, substring)
//
// 	k := len(mainString) + 1 - len(substring)
// 	for i := 0; i < k; i++ {
// 		if mainString[i:i+len(substring)] == substring {
// 			return i
// 		}
// 	}
// 	return -1
//
// 	// sadbutsad = 9
// 	// sad = 3
//
// 	// 9 + 1 -3 = 7 //we do this to not get to the final end but like a letter behind the end
//
// 	// s: ad == sad
// }

// func searchInsert(nums []int, target int) int {
// 	// l := 0
// 	// r := len(nums) - 1
// 	//
// 	// for l <= r {
// 	// 	mid := (l + r) / 2
// 	// 	if nums[mid] < target {
// 	// 		l = mid + 1
// 	// 	} else if nums[mid] > target {
// 	// 		r = mid - 1
// 	// 	} else {
// 	// 		return mid
// 	// 	}
// 	// }
// 	// return l
// 	//Binary search using Olog(n)
//
// 	for k := range nums {
// 		if nums[k] == target {
// 			return k
// 		} else if nums[k] > target {
// 			return k
// 		}
// 	}
// 	return len(nums)
// 	// linear Search
// }

// func lengthOfLastWord(s string) int {
// 	l := 0
// 	i := len(s) - 1
//
// 	for i >= 0 && s[i] == ' ' {
// 		i--
// 	}
//
// 	for i >= 0 && s[i] != ' ' {
// 		l++
// 		i--
// 	}
// 	return l
// }

// func plusOne(digits []int) []int {
// 	r := len(digits) - 1
//
// 	for i := r; i >= 0; i-- {
// 		if digits[i] < 9 {
// 			digits[i]++
// 			return digits
// 		}
// 		digits[i] = 0
// 	}
// 	digits = append([]int{1}, digits...)
// 	return digits
// }

// func addBinary(a string, b string) string {
// 	res := ""
// 	carry := 0
//
// 	i, j := len(a)-1, len(b)-1
//
// 	for i >= 0 || j >= 0 || carry > 0 {
// 		sum := carry
//
// 		if i >= 0 {
// 			sum += int(a[i] - '0') // the string here is converted to int here
// 			i--
// 		}
//
// 		if j >= 0 {
// 			sum += int(b[j] - '0')
// 			j--
// 		}
//
// 		rem := sum % 2
// 		res = fmt.Sprintf("%d%s", rem, res)
//
// 		carry = sum / 2
// 	}
// 	return res
// }

// We can use the standard library but in interviews you wont be allowed to use it, so Therefore use the normal formula
// func addBinary(a, b string) string {
// 	num1, _ := new(big.Int).SetString(a, 2)
// 	num2, _ := new(big.Int).SetString(b, 2)
//
// 	sum := new(big.Int).Add(num1, num2)
// 	return sum.Text(2)
// }
//
// there are other standard library to use like strconv.ParseInt and strconv.FormatInt

// func mySqrt(n int) int {
// 	// without using the standard library
// 	l, r := 0, n
//
// 	for l <= r {
// 		mid := (l + r) / 2
//
// 		if mid*mid < n {
// 			l = mid + 1
// 		} else if mid*mid > n {
// 			r = mid - 1
// 		} else {
// 			return mid
// 		}
// 	}
// 	return r
// }

// func climbStairs(n int) int {
// 	// this is related to the fibonacci (dynamic programming) and we can use that recursion for it also till n got to 1 and 2
// 	if n == 1 {
// 		return 1
// 	}
//
// 	n1 := 1
// 	n2 := 1
// 	total := 0
//
// 	for i := 2; i <= n; i++ {
// 		total = n1 + n2
// 		n2 = n1
// 		n1 = total
// 	}
// 	return total
// }

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }
//
// func delDuplicates(head *ListNode) *ListNode {
// 	res := head
//
// 	for head != nil && head.Next != nil {
// 		if head.Next.Val == head.Val {
// 			head.Next = head.Next.Next
// 		} else {
// 			head = head.Next
// 		}
// 	}
// 	return res
// }

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	c := (m + n) - 1
// 	i, j := m-1, n-1
//
// 	for j >= 0 {
// 		if i >= 0 && nums1[i] > nums2[j] {
// 			nums1[c] = nums1[i]
// 			i--
// 		} else {
// 			nums1[c] = nums2[j]
// 			j--
// 		}
// 		c--
// 	}
// }

// This is a long range of binary tree question so we are going to be using this struct all true
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//
// func inorderTraversal(root *TreeNode) []int {
// 	res := make([]int, 0) // create the arr that store the result
// 	stck := make([]*TreeNode, 0)
// 	// create a slack of our datatype
// 	curr := root
//
// 	// note: we are trving the left tree first
// 	for curr != nil || len(stck) > 0 {
// 		for curr != nil {
// 			stck = append(stck, curr)
// 			// the left is being pushed to the stack
// 			curr = curr.Left
// 			// dont forget we want to be moving left
// 		}
// 		curr = stck[len(stck)-1]
// 		stck = stck[:len(stck)-1]
// 		// and when it doesnot have anything it is being popped ande appended to our res here
// 		res = append(res, curr.Val)
// 		curr = curr.Right
// 	}
// 	return res
/* I can practically use recursion here which append the left first and then the root with the right late
just like this (append(ourarr, inorderTraversal(root.Left)...), followed by append to root append(ourarr, inorderTraversal(root.Val)) and then append right) */
// }

// func isSameTree(p, q *TreeNode) bool {
// 	if p == nil && q == nil {
// 		return true
// 	}
//
// 	// check if one of them is empty or their values are not the same
// 	if (p == nil || q == nil) || (p.Val != q.Val) {
// 		return false
// 	}
//
// 	// now, this is the recursive step that checks the right and left
// 	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
// }

// func isSymmetry(root *TreeNode) bool {
// 	var dfs func(left *TreeNode, right *TreeNode) bool
//
// 	dfs = func(left, right *TreeNode) bool {
// 		if left == nil && right == nil {
// 			return true
// 		}
// 		if left == nil || right == nil {
// 			return false
// 		}
// 		// Corrected the order of comparison for opposite sides: left.Right with right.Left
// 		return left.Val == right.Val && dfs(left.Right, right.Left) && dfs(left.Left, right.Right)
// 	}
//
// 	// Check symmetry starting from the root
// 	return dfs(root.Left, root.Right)
// }

// func maximumDepth(root *TreeNode) int {
// 	max := func(a, b int) int {
// 		if a > b {
// 			return a
// 		}
// 		return b
// 	}
//
// 	if root == nil {
// 		return 0
// 	}
//
// 	leftHeight := maximumDepth(root.Left)
// 	rightHeight := maximumDepth(root.Right)
//
// 	return 1 + max(leftHeight, rightHeight)
// }

// func sortedArray2BST(nums []int) *TreeNode {
// 	if len(nums) == 0 {
// 		return nil
// 	}
//
// 	m := len(nums) / 2
// 	root := &TreeNode{Val: nums[m]}
// 	root.Left = sortedArray2BST(nums[:m])
// 	root.Right = sortedArray2BST(nums[m+1:])
// 	return root
// }

// func isBalanced(root *TreeNode) bool {
// 	var dfs func(root *TreeNode) (bool, int)
// 	var abs func(n int) int
// 	var max func(a, b int) int
//
// 	max = func(a, b int) int {
// 		if a > b {
// 			return a
// 		}
// 		return b
// 	}
//
// 	abs = func(n int) int {
// 		if n < 0 {
// 			return -n
// 		}
// 		return n
// 	}
//
// 	dfs = func(root *TreeNode) (bool, int) {
// 		if root == nil {
// 			return true, 0
// 		}
//
// 		isLeftBalanced, leftHeight := dfs(root.Left)
// 		isRightBalanced, rightHeight := dfs(root.Right)
// 		diff := abs(leftHeight - rightHeight)
// 		if isLeftBalanced && isRightBalanced && diff <= 1 {
// 			return true, 1 + max(leftHeight, rightHeight)
// 		}
// 		return false, -1
// 	}
// 	ans, _ := dfs(root)
// 	return ans
//
// 	// I can use 0(n2) time complexity that get the max depth first and then find the diff between them (if it is greater than 1 then we return false)
// }

// func minimumDepth(root *TreeNode) int {
// 	var min func(a, b int) int
//
// 	min = func(a, b int) int {
// 		if a == 0 {
// 			return b
// 		}
// 		if b == 0 {
// 			return a
// 		}
// 		if a < b {
// 			return a
// 		} else {
// 			return b
// 		}
// 	}
//
// 	if root == nil {
// 		return 0
// 	}
//
// 	leftHeight := minimumDepth(root.Left)
// 	rightHeight := minimumDepth(root.Right)
// 	return 1 + min(leftHeight, rightHeight)
// }

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum = targetSum - root.Val

	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}
