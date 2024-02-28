package main

import (
	//	"strings"
	//	"unicode"
	"fmt"
	"leetCode/struct"
	"math"
	"math/bits"
	"strconv"
	"strings"
	"unicode"
	// "math/big"
)

// duplicate num in a array

func DupNum(nums []int) bool {
	s := make(map[int]bool)

	for _, num := range nums {
		if s[num] {
			return true
			// because it ask us to return true if it has the same nos of occurrence, check unique occur(question 1207 for the reason)
		}
		s[num] = true
	}
	return false
}

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
// 		// and when it doesnot have anything it is being popped and appended to our res here
// 		res = append(res, curr.Val)
// 		curr = curr.Right
// 	}
// 	return res
/* I can practically use recursion here which append the left first and then the root with the right late
just like this (append(ourarr, inorderTraversal(root.Left)...), followed by append to root append(ourarr, inorderTraversal(root.Val)) and then append right) */
// }

// func preOrderTraversal(root *TreeNode) []int {
// 	res := make([]int, 0)
// 	stck := make([]*TreeNode, 0)
// 	curr := root
//
// 	for curr != nil || len(stck) > 0 {
// 		for curr != nil {
// 			res = append(res, curr.Val)
// 			stck = append(stck, curr)
// 			curr = curr.Left
// 		}
// 		curr = stck[len(stck)-1]
// 		stck = stck[:len(stck)-1]
// 		curr = curr.Right
// 	}
// 	return res
// 	// the recursion method and also trat the base case if where root == 0
// 	// return append([]int{root.Val}, append(preOrderTraversal(root.Left), preOrderTraversal(root.Right)...)...)
// }

// func postOrderTraversal(root *TreeNode) []int {
// 	res := make([]int, 0)
// 	stck := make([]*TreeNode, 0)
// 	curr := root
//
// 	for curr != nil || len(stck) > 0 {
// 		for curr != nil {
// 			stck = append(stck, curr)
// 			res = append([]int{curr.Val}, res...)
// 			curr = curr.Right
// 		}
// 		curr = stck[len(stck)-1]
// 		stck = stck[:len(stck)-1]
// 		curr = curr.Left
// 	}
// 	return res
// 	// the recursion method after you have treated the base case finish
// 	// return append(
// 	// 	append(postOrderTraversal(root.Left), postOrderTraversal(root.Right)...),
// 	// 	root.Val,
// 	// )
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

// func hasPathSum(root *TreeNode, targetSum int) bool {
// 	if root == nil {
// 		return false
// 	}
//
// 	targetSum = targetSum - root.Val
//
// 	if root.Left == nil && root.Right == nil {
// 		return targetSum == 0
// 	}
// 	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
// }

// func generatePascal(numsRows int) [][]int {
// 	res := make([][]int, numsRows)
//
// 	for i := 0; i < numsRows; i++ {
// 		// just understand here that for the first 2, which i = 0 and 1 j doesnot get to execute. J start executing when i = 2 that means it have 2 nos
// 		rows := make([]int, i+1)
// 		rows[0], rows[i] = 1, 1
// 		for j := 1; j < i; j++ {
// 			rows[j] = res[i-1][j-1] + res[i-1][j]
// 		}
// 		res[i] = rows
// 	}
// 	return res
// }
//
// // I love this solution it is straight forward and easily understandable
// func getRow(rowIndex int) []int {
// 	result := []int{1}
//
// 	var value int = 1
//
// 	for i := 1; i <= rowIndex; i++ {
// 		value = value * (rowIndex - i + 1) / i
//
// 		result = append(result, value)
// 	}
//
// 	return result
// }

// This another era and we are using a defined max function here
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
//
// func bestTimeStock(prices []int) int {
// 	buy := 0
// 	profit := 0
//
// 	for sell := 1; sell < len(prices); sell++ {
// 		if prices[sell] > prices[buy] {
// 			profit = max(profit, prices[sell]-prices[buy])
// 		} else {
// 			buy = sell
// 		}
// 	}
// 	return profit
// }
//
// func maxProfit(prices []int) int {
// 	buy := 0
// 	profit := 0
//
// 	for sell := 1; sell < len(prices); sell++ {
// 		if prices[sell] > prices[buy] {
// 			profit += prices[sell] - prices[buy]
// 		}
// 		buy = sell
// 	}
// 	return profit
// }

// This check among 3 nos
// O(n2)
//
//	func singleNum(nums []int) int {
//		seen := make(map[int]int)
//
//		for _, num := range nums {
//			seen[num]++
//		}
//
//		for num, count := range seen {
//			if count == 1 {
//				return num
//			}
//		}
//		return 0
//	}
//
// O(n)
// func singleNum(nums []int) int {
// 	ones := 0
// 	twos := 0
//
// 	for _, num := range nums {
// 		ones ^= (num & ^twos)
// 		twos ^= (num & ^ones)
// 	}
// 	return ones
// }

// func hasCycle(head *ListNode) bool {
// 	slow, fast := head, head
//
// 	for fast != nil && fast.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 		if slow == fast {
// 			return true
// 		}
// 	}
// 	return false
// }

// func detectCycle(head *ListNode) *ListNode {
// 	slow, fast := head, head
//
// 	for fast != nil && fast.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 		if slow == fast {
// 			slow = head
// 			for slow != fast {
// 				slow = slow.Next
// 				fast = fast.Next
// 			}
// 			return slow
// 		}
// 	}
// 	return nil
// }
//
// // Using Hashmap
// func detectCycle(head *ListNode) *ListNode {
// 	seen := make(map[*TreeNode]bool)
//
// 	for i := head; i != nil; i = i.Next {
// 		if _, ok := seen[i]; ok {
// 			return i
// 		}
// 		seen[i] = true
// 	}
// 	return nil
// }

// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	seen := make(map[*ListNode]bool)
//
// 	for i := headA; i != nil; i = i.Next {
// 		seen[i] = true
// 	}
//
// 	for j := headB; j != nil; j = j.Next {
// 		if seen[j] {
// 			return j
// 		}
// 	}
// 	return nil
// }

// For more optimization we can use the 2 pointer notation
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	a, b := headA, headB // create our 2 pointers
// 	for a != b {
// 		if a != nil {
// 			a = a.Next
// 		} else {
// 			a = headB
// 		} // anytime a becomes nil, it move to the other ListNode
//
// 		if b != nil {
// 			b = b.Next
// 		} else {
// 			b = headA
// 		} // vice Versa here also
// 	}
// 	return a // We return one of them
// }

// func excelSheetConv(colunmnNumber int) string {
// 	res := ""
// 	for colunmnNumber > 0 {
// 		rem := (colunmnNumber - 1) % 26
// 		res = string(rune('A'+rem)) + res
// 		colunmnNumber = (colunmnNumber - 1) / 26
// 	}
// 	return res
// }
//
// func excelSheetConvNum(columnTitle string) int {
// 	res := 0
// 	mul := 1 // multiplier
// 	for i := len(columnTitle) - 1; i >= 0; i-- {
// 		res = int(columnTitle[i]-'A'+1)*mul + res
// 		mul *= 26
// 	}
// 	return res
// }

// func majorityNumber(nums []int) int {
// 	// We are going to be using 2 method to solve this. 1. Hashmap, 2. pointer solution (boyer moore algorithm)
//
// 	// Using Hashmap
// 	seen := make(map[int]int)
// 	m := len(nums) / 2
// 	for _, num := range nums {
// 		seen[num]++
// 	}
//
// 	for i, count := range seen {
// 		if count > m {
// 			return i
// 		}
// 	}
// 	return 0
// }

// func majorityNumber(nums []int) int {
// 	// Using boyer moore
// 	res, count := 0, 0
//
// 	for _, num := range nums {
// 		if count == 0 {
// 			res = num
// 			count = 1
// 		} else if num == res {
// 			count++
// 		} else {
// 			count--
// 		}
// 	}
// 	return res
// }

func reverseBit(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res <<= 1             // for every loop we shift result to the left by 1
		res = res | (num & 1) // performs a bitwise operation
		num >>= 1             // shifts the num to the right by 1
	}
	return res
}

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		count += int(num % 2)
		num >>= 1
	}
	return count
	// for more optimization, we can do num = num & (num - 1) and then increment the count
	// We minus 1 from our num and then & it
}

func isHappy(n int) bool {
	seen := make(map[int]bool)

	for n != 1 && !seen[n] {
		seen[n] = true
		sum := 0

		for n > 0 {
			rem := n % 10
			sum += rem * rem
			n /= 10
		}
		n = sum
	}
	return n == 1
}

func removeLinkedList(head *util.ListNode, val int) *util.ListNode {
	dummyNode := &util.ListNode{Next: head}
	prev, curr := dummyNode, head

	for curr != nil {
		nxt := curr.Next
		if curr.Val == val {
			prev.Next = nxt
		} else {
			prev = curr
		}
		curr = nxt
	}
	return dummyNode.Next
}

func isIsomorphics(s, t string) bool {
	seenS, seenT := make(map[byte]byte), make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		s1, t1 := s[i], t[i]
		if value, ok := seenS[t1]; ok && value != s1 {
			return false
		}
		seenS[t1] = s1

		if value, ok := seenT[s1]; ok && value != t1 {
			return false
		}
		seenT[s1] = t1
	}
	return true
}

func reverseList(head *util.ListNode) *util.ListNode {
	// iterative solution
	// We can use iteration method also but i am not taking it necessary
	var prev *util.ListNode
	curr := head

	for curr != nil {
		nxt := curr.Next
		curr.Next = prev
		prev = curr
		curr = nxt
	}
	return prev
}

func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]int)

	for i, num := range nums {
		if j, exists := seen[num]; exists && i-j <= k {
			return true
		}
		seen[num] = i
	}
	return false
}

func countNodes(root *util.TreeNode) int {
	// guess what we can use both the inorder, preorder and postorder to traverse around the treenode and return the count but let just use recursion here to get the countNodes

	if root == nil {
		return 0
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func inverseNode(root *util.TreeNode) *util.TreeNode {
	if root == nil {
		return nil
	}
	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	inverseNode(root.Left)
	inverseNode(root.Right)
	return root
}

func summaryRanges(nums []int) []string {
	var res []string

	for i := 0; i < len(nums); {
		k := strconv.Itoa(nums[i])
		j := i + 1

		for ; j < len(nums) && nums[j]-nums[j-1] == 1; j++ {
		}
		if j-1 > i {
			k += "->" + strconv.Itoa(nums[j-1])
		}
		res = append(res, k)
		i = j
	}
	return res
}

func isPowerOfTwos(n int) bool {
	temp := n

	for temp > 1 {
		if temp%2 != 0 {
			return false
		}
		temp /= 2
	}
	return temp == 1

	// you can use bit manipulation also
	// func isPowerOfTwo(n int) bool {
	//     return n != 0 && (n & (n - 1) == 0)
	// }
}

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}

	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}

func isPowerOfFour(n int) bool {
	if n < 1 {
		return false
	}

	for n%4 == 0 {
		n /= 4
	}
	return n == 1
}

// func isLinkPalindrome(head *util.ListNode) bool {
// 	slow := head
// 	fast := head
//
// 	// Find the middle node
// 	for fast != nil && fast.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 	}
//
// 	// reverse the node
// 	var prev *util.ListNode
// 	for slow != nil {
// 		temp := slow.Next
// 		slow.Next = prev
// 		prev = slow
// 		slow = temp
// 	}
//
// 	// compare the first node and the new node prev
// 	l, r := head, prev
// 	for r != nil {
// 		if l.Val != r.Val {
// 			return false
// 		}
// 		l = l.Next
// 		r = r.Next
// 	}
// 	return true
// }

func isLinkPalindrome(head *util.ListNode) bool {
	var nums []int

	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l] != nums[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func binaryTreePaths(root *util.TreeNode) []string {
	var res []string
	var rec func(*util.TreeNode, string)
	rec = func(node *util.TreeNode, path string) {
		if node.Left == nil && node.Right == nil {
			res = append(res, path)
			return
		}

		if node.Left != nil {
			rec(node.Left, fmt.Sprintf("%s->%d", path, node.Left.Val))
		}
		if node.Right != nil {
			rec(node.Right, fmt.Sprintf("%s->%d", path, node.Right.Val))
		}
	}

	rec(root, fmt.Sprintf("%d", root.Val))
	return res
}

func addDigits(num int) int {
	if num < 10 {
		return num
	}

	if num%9 == 0 {
		return 9
	}
	return num % 9
}

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	for _, i := range []int{2, 3, 5} {
		if n%i == 0 {
			n /= i
		}
	}
	return n == 1
}

func missingNum(nums []int) int {
	n := len(nums)
	total := (n + 1) * n / 2 // sum of num from 0 to n
	for _, num := range nums {
		total -= num // subtract from num and then return reminder
	}
	return total
}

// first bad version is solve using binary search from the middle to the left and right

func moveZeroes(nums []int) {
	r := 0
	for l := 0; l < len(nums); l++ {
		if nums[l] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			r++
		}
	}
}

func wordPattern(pattern string, s string) bool {
	realS := strings.Split(s, " ")

	if len(pattern) != len(realS) {
		return false
	}

	seen1 := make(map[string]string)
	seen2 := make(map[string]byte)

	for i, str := range realS {
		realP := string(pattern[i]) // convert all letter each in pattern to string
		if val, ok := seen1[str]; ok && val != realP {
			return false
		}

		if val, ok := seen2[realP]; ok && val != str[0] {
			return false
		}
		seen1[str] = realP
		seen2[realP] = str[0]
	}
	return true
}

func countBits(n int) []int {
	ans := make([]int, n+1)
	count := 1

	for i := 1; i <= n; i++ {
		if count*2 == i {
			count = i
		}
		ans[i] = 1 + ans[i-count]
	}
	return ans
}

func reverseStrings(s []byte) {
	r := len(s) - 1

	for l := 0; l < r; l++ {
		s[l], s[r] = s[r], s[l]
		r--
	}
}

func reverseVowels(s string) string {
	var isVowel func(s rune) bool
	temp := []rune(s)

	isVowel = func(s rune) bool {
		return s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' || s == 'A' || s == 'E' ||
			s == 'I' ||
			s == 'O' ||
			s == 'U'
	}

	for i, j := 0, len(temp)-1; i < j; {
		if isVowel(temp[i]) && isVowel(temp[j]) {
			temp[i], temp[j] = temp[j], temp[i]
			i++
			j--
		} else if !isVowel(temp[i]) {
			i++
		} else if !isVowel(temp[j]) {
			j--
		}
	}
	return string(temp)
}

func intersection(nums1 []int, nums2 []int) []int {
	seen := make(map[int]bool)
	res := []int{}

	for _, num := range nums1 {
		seen[num] = true
	}

	for _, num := range nums2 {
		if seen[num] {
			res = append(res, num)
			seen[num] = false
		}
	}
	return res
}

func intersect(nums1, nums2 []int) []int {
	if len(nums1) < len(nums2) {
		return intersect(nums2, nums1)
	}

	seen := make(map[int]int)
	res := make([]int, 0)

	for _, num := range nums1 {
		seen[num]++
	}

	for _, num := range nums2 {
		if count, exists := seen[num]; exists && count > 0 {
			res = append(res, num)
			seen[num]--
		}
	}
	return res
}

func validSquare(n int) bool {
	l, r := 0, n+1

	for l < r {
		m := l + (r-l)/2

		if m*m >= n {
			r = m // we have just equate right to mid cause if the mid * mid is > n then we can have a new right
		} else {
			l = m + 1
		}
	}
	return l*l == n
}

func canConstruct(ransomNote, magazine string) bool {
	seenR := make(map[rune]int)
	seenM := make(map[rune]int)

	for _, each := range ransomNote {
		seenR[each]++
	}

	for _, each := range magazine {
		seenM[each]++
	}

	for char, count := range seenR {
		count2, exists := seenM[char]
		if !exists || count > count2 {
			return false
		}
	}
	return true
}

// func firstUniqueChar(s string) int {
// 	seen := make(map[rune]int)
//
// 	for _, each := range s {
// 		seen[each]++
// 	}
//
// 	for i, char := range s {
// 		if seen[char] == 1 {
// 			return i
// 		}
// 	}
// 	return -1
// }

func firstUniqueChar(s string) int {
	// the array method
	seen := make([]int, 26)

	for _, each := range s {
		seen[each-'a']++
	}

	for i, char := range s {
		if seen[char-'a'] == 1 {
			return i
		}
	}
	return -1
}

func findTheDifference(s, t string) byte {
	// i can use hashmap to solve this but let use use the bit manipulation
	// for the hashmap we creat 2 hashmap and then we caheck for count diff and return the byte(count)

	var result byte = 0

	for _, char := range s {
		result ^= byte(char)
	}

	for _, char := range t {
		result ^= byte(char)
	}

	return result
}

func isSubsequence(s, t string) bool {
	l, r := 0, 0

	for l < len(s) && r < len(t) {
		if s[l] == t[r] {
			l++
		}
		r++
	}
	return l == len(s)
}

func readBinaryWatch(turnedOn int) []string {
	time := []string{}

	for h := 0; h < 12; h++ {
		i := bits.OnesCount(uint(h))
		for m := 0; m < 60; m++ {
			j := bits.OnesCount(uint(m))
			if turnedOn == i+j {
				time = append(time, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return time
}

func sumOfLeftLeaves(root *util.TreeNode) int {
	res := 0

	if root == nil {
		return 0
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		res = root.Left.Val
	}

	return res + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

func toHex(n int) string {
	res := ""
	const hexChar = "0123456789abcdef"

	if n == 0 {
		return "0"
	}

	if n < 1 {
		n += 0x100000000
	}

	for n > 0 {
		rem := n % 16
		res = string(hexChar[rem]) + res
		n /= 16
	}
	return res
}

func longestPalindrome(s string) int {
	charCount := make(map[rune]int)
	ans := 0
	oddCount := false

	for _, char := range s {
		charCount[char]++
	}

	for _, count := range charCount {
		if count%2 == 0 {
			ans += count
		} else {
			ans += count - 1
			oddCount = true
		}
	}

	if oddCount {
		ans++
	}
	return ans
}

func fizzbuzz(n int) []string {
	res := make([]string, 0)

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else if i%5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, fmt.Sprintf("%d", i))
		}
	}
	return res
}

func thirdMax(nums []int) int {
	first, second, third := math.MinInt64, math.MinInt64, math.MinInt64

	for _, num := range nums {
		if num == first || num == second || num == third {
			// skipping duplicates here, if the number is seen we just skip and continue
			continue
		}

		if num > first {
			third, second, first = second, first, num
		} else if num > second {
			third, second = second, num
		} else if num > third {
			third = num
		}
	}

	if third == math.MinInt64 {
		return first
	}

	return third
}

func addStrings(num1, num2 string) string {
	res := ""
	carry := 0

	i, j := len(num1)-1, len(num2)-1

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}

		carry = sum / 10
		digit := sum % 10

		res = fmt.Sprintf("%d", digit) + res
	}
	return res
}

func countSegments(s string) int {
	str := strings.FieldsFunc(s, unicode.IsSpace)
	count := 0

	for _, each := range str {
		if each != "" {
			count++
		}
	}
	return count
}

func arrangeCoins(n int) int {
	l, r := 1, n
	res := 0

	for l <= r {
		m := l + (r-l)/n
		coins := m * (m + 1) / 2
		if coins > n {
			r = m - 1
		} else {
			res = m
			l = m + 1
		}
	}
	return res
}
