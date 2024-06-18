package main

import (
	"fmt"
	"math"
	"math/bits"
	"reflect"
	"slices"
	"sort"
	"strconv"
	"strings"
	"unicode"

	util "leetCode/struct"
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

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			return mid
		}
	}
	return l
	// Binary search using Olog(n)

	// for k := range nums {
	// 	if nums[k] == target {
	// 		return k
	// 	} else if nums[k] > target {
	// 		return k
	// 	}
	// }
	// return len(nums)
	// linear Search

	// we can also use map to solve this
}

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
// 	// create a stack of our datatype
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

func isSameTree(p, q *util.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	// check if one of them is empty or their values are not the same
	if (p == nil || q == nil) || (p.Val != q.Val) {
		return false
	}

	// now, this is the recursive step that checks the right and left
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

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

// func arrangeCoins(n int) int {
// 	l, r := 1, n
// 	res := 0
//
// 	for l <= r {
// 		m := l + (r-l)/n
// 		coins := m * (m + 1) / 2
// 		if coins > n {
// 			r = m - 1
// 		} else {
// 			res = m
// 			l = m + 1
// 		}
// 	}
// 	return res
// }

func arrangeCoins(n int) int {
	k := 1
	n--

	for n >= 0 {
		k++
		n -= k
	}
	return k - 1
}

func findAllDuplicates(nums []int) []int {
	res := []int{}

	for _, i := range nums {
		if nums[util.Abs(i)-1] < 0 {
			res = append(res, util.Abs(i))
		}
		nums[util.Abs(i)-1] *= -1
	}
	return res
}

func findDisapperedNumber(nums []int) []int {
	k := len(nums)
	res := make([]int, k)

	if k == 0 {
		return nil
	}

	for _, num := range nums {
		res[num-1] = num
	}

	j := 0
	for i := range res {
		if res[i] == 0 {
			res[j] = i + 1
			j++
		}
	}
	return res[:j]
}

func assignCookies(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	count := 0
	i, j := 0, 0

	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			count++
			i++
		}
		j++
	}
	return count
}

func hammingDistance(x int, y int) int {
	or := x ^ y
	count := 0

	for or > 0 {
		count += or & 1
		or >>= 1
	}
	return count
}

func reapeatedSubString(s string) bool {
	len := len(s)

	double := s[1:len] + s[0:len-1]
	return strings.Contains(double, s)
}

func findComplement(n int) int {
	temp := 1

	for temp < n {
		temp = (temp << 1) | 1
	}
	return n ^ temp
	// same as question 1009 bitwise complement
}

func licenseFormatting(s string, k int) string {
	s = strings.ToUpper(strings.Replace(s, "-", "", -1))

	mod := len(s) % k

	if mod == 0 {
		mod += k // new mod nos
	}

	for mod < len(s) {
		s = s[:mod] + "-" + s[mod:]
		mod = mod + k + 1
	}
	return s
}

func findMaxConsecutivesOne(nums []int) int {
	temp := 0
	max := 0

	for _, num := range nums {
		if num == 1 {
			temp++
			if temp > max {
				max = temp
			}
		} else {
			temp = 0
		}
	}
	return max
}

func construtRectangle(area int) []int {
	temp := int(math.Sqrt(float64(area)))

	for temp > 0 {
		if area%temp == 0 {
			return []int{area / temp, temp}
		}
		temp--
	}
	return nil
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	t := 0

	for i := 0; i < len(timeSeries)-1; i++ {
		if timeSeries[i]+duration-1 < timeSeries[i+1] {
			t += duration
		} else {
			t += timeSeries[i+1] - timeSeries[i]
		}
	}
	return t + duration
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	seen := make(map[int]int)
	res := make([]int, len(nums1))

	for i, num := range nums2 {
		seen[num] = i
	}

	for i, num := range nums1 {
		found := false
		for next := seen[num] + 1; next < len(nums2); next++ {
			if nums2[next] > num {
				res[i] = nums2[next]
				found = true
				break
			}
		}
		if !found {
			res[i] = -1
		}
	}
	return res
}

func findWords(words []string) []string {
	letters := make([]int, 26)
	res := make([]string, 0)

	for _, r := range "qwertyuiop" {
		letters[r-'a'] = 1
	}

	for _, r := range "asdfghjkl" {
		letters[r-'a'] = 2
	}

	for _, r := range "zxcvbnm" {
		letters[r-'a'] = 3
	}

	for _, each := range words {
		w := strings.ToLower(each)
		match := true
		for i := 1; i < len(w); i++ {
			if letters[w[i]-'a'] != letters[w[0]-'a'] {
				match = false
				break
			}
		}
		if match {
			res = append(res, each)
		}
	}
	return res
}

func findMode(root *util.TreeNode) []int {
	seen := make(map[int]int)
	var Inorder func(root *util.TreeNode)
	max := 0
	res := []int{}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	Inorder = func(root *util.TreeNode) {
		if root == nil {
			return
		}
		seen[root.Val]++
		if max < seen[root.Val] {
			max = root.Val
		}

		Inorder(root.Left)
		Inorder(root.Right)
	}
	Inorder(root)

	for i, val := range seen {
		if val == max {
			res = append(res, i)
		}
	}
	return res
}

func base7(n int) string {
	res := ""
	if n == 0 {
		return "0"
	}

	negative := n < 0
	if negative {
		n = -n
	}

	for n != 0 {
		rem := n % 7
		res = strconv.Itoa(rem) + res
		n /= 7
	}

	if negative {
		res = "-" + res
	}
	return res
}

func findRelativeRank(score []int) []string {
	seen := make(map[int]int)
	res := make([]string, len(score))

	for i, each := range score {
		seen[each] = i
	}

	sort.Sort(sort.Reverse(sort.IntSlice(score)))
	for i, val := range score {
		switch i {
		case 0:
			res[seen[val]] = "Gold Medal"
		case 1:
			res[seen[val]] = "Silver Medal"
		case 2:
			res[seen[val]] = "Bronze Medal"
		default:
			res[seen[val]] = strconv.Itoa(i + 1)
		}
	}
	return res
}

func perfectNum(num int) bool {
	// there are only 6 perfect nos (6,28,496,8128,3355036, 8589869056)
	// i can maybe range through it and boom

	if num <= 1 {
		return false
	}
	temp := 1
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			temp += i
			if i != num/i {
				temp += num / i
			}
		}
	}
	return temp == num
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	nMinus2 := 0
	nMinus1 := 1
	cur := 0
	for i := 2; i <= n; i++ {
		cur = nMinus2 + nMinus1
		nMinus2 = nMinus1
		nMinus1 = cur
	}
	return cur
}

// func detectCapitalUse(s string) bool {
// 	allCaps := strings.ToUpper(s)
// 	remainingPart := strings.ToLower(s)[1:]
//
// 	return s == allCaps || s == remainingPart
// }

func detectCapitalUse(S string) bool {
	var isCapital func(c byte) bool = func(c byte) bool {
		return c >= 'A' && c <= 'Z'
	}

	caps := 0
	for i := range S {
		if isCapital(S[i]) {
			caps++
		}
	}
	return (caps == len(S)) || (caps == 1 && isCapital(S[0])) || (caps == 0)
}

func reverseStr(s string, k int) string {
	ss := []byte(s)

	for i := 0; i < len(ss); i = i + 2*k {
		l := i
		r := util.Min(i+k-1, len(ss)-1)
		for l < r {
			ss[l], ss[r] = ss[r], ss[l]
			l++
			r--
		}
	}
	return string(ss)
}

func findLusLength(a string, b string) int {
	if a == b {
		return -1
	}
	if len(a) > len(b) {
		return len(a)
	}
	return len(b)
}

func getMinimumDifference(root *util.TreeNode) int {
	prev := -1
	minDiff := math.MinInt64

	var Inorder func(root *util.TreeNode)
	Inorder = func(root *util.TreeNode) {
		if root == nil {
			return
		}

		Inorder(root.Left)
		if prev != -1 && root.Val-prev < minDiff {
			minDiff = root.Val - prev
		}
		prev = root.Val
		Inorder(root.Right)
	}
	Inorder(root)
	return minDiff
}

func diameterOfBinaryTree(root *util.TreeNode) int {
	res := 0

	var dfs func(root *util.TreeNode) int
	dfs = func(root *util.TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)

		res = util.Max(res, left+right)
		return util.Max(left, right) + 1
	}
	dfs(root)
	return res
}

func checkRecords(s string) bool {
	absent := 0
	alwaysLate := 0

	for _, char := range s {
		if char == 'A' {
			absent++
			if absent > 1 {
				return false
			}
		}
		if char == 'L' {
			alwaysLate++
			if alwaysLate > 2 {
				return false
			}
		} else {
			alwaysLate = 0
		}
	}
	return true
}

func reverseWords(s string) string {
	char := []byte(s)
	l, r := 0, 0
	lastSpaceIndex := -1

	for i := 0; i <= len(char); i++ {
		if i == len(char) || char[i] == ' ' {
			l, r = lastSpaceIndex+1, i-1
			for l < r {
				char[l], char[r] = char[r], char[l]
				l++
				r--
			}
			lastSpaceIndex = i
		}
	}
	return string(char)
}

func maxDepth(root *util.Node) int {
	// this is on the nth-nary tree and not binary tree

	if root == nil {
		return 0
	}

	max := 0
	for _, v := range root.Children {
		max = util.Max(max, maxDepth(v))
	}
	return max + 1
}

func arrayPairSum(nums []int) int {
	sum := 0
	slices.Sort(nums)

	for i := 0; i < len(nums)-0; i += 2 {
		sum += nums[i]
	}
	return sum
}

func findTilt(root *util.TreeNode) int {
	res := 0

	var temp func(root *util.TreeNode) int
	temp = func(root *util.TreeNode) int {
		if root == nil {
			return 0
		}

		l := temp(root.Left)
		r := temp(root.Right)
		res += util.Abs(l - r)
		return l + r + root.Val
	}
	temp(root)
	return res
}

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m := len(mat)
	n := len(mat[0])

	res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
	}

	newRow, newCol := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[newRow][newCol] = mat[i][j]
			newCol++
			if newCol == c {
				newCol = 0
				newRow++
			}
		}
	}
	return res
}

func isSubTree(root *util.TreeNode, subRoot *util.TreeNode) bool {
	if root == nil && subRoot == nil {
		return root == subRoot
	}

	if root == nil || subRoot == nil {
		return root != subRoot
	}

	if isSameTree(root, subRoot) {
		return true
	}

	return isSubTree(root.Left, subRoot) || isSubTree(root.Right, subRoot)
}

func distributeCandy(candyType []int) int {
	seen := make(map[int]int)
	n := len(candyType) / 2

	for _, i := range candyType {
		seen[i]++
	}
	if n > len(seen) {
		return len(seen)
	} else {
		return n
	}
}

func preOrder(root *util.Node) []int {
	// this is for the nth-nary tree and we can use recursion and iteration for this preOrder
	res := []int{}
	stack := []*util.Node{root}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return res
	//the recursive will be like this after the base case of root == 0
	// for _, child := range Children{
	//res = append(res, preOrder(child)...)
	//}
}

func postOrder(root *util.Node) []int {
	res := []int{}
	stack := []*util.Node{root}

	if root == nil {
		return res
	}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append([]int{node.Val}, res...)
		for _, child := range node.Children {
			stack = append(stack, child)
		}
	}
	return res
	//recursion on this
	// for i := range root.Children{
	// res = append(res, postOrder(root.Children[i])...)
	//}
	// res = append(res, root.Val)
}

func findLHS(nums []int) int {
	seen := make(map[int]int)
	res := 0

	for _, num := range nums {
		seen[num]++
	}

	for n, c1 := range seen {
		if c2, ok := seen[n+1]; ok {
			t := c1 + c2
			res = util.Max(res, t)
		}
	}
	return res
}

func maxCount(m int, n int, ops [][]int) int {
	x, y := m, n

	for _, op := range ops {
		x = util.Min(op[0], x)
		y = util.Min(op[1], y)
	}
	return x * y
}

func findRestaurant(list1 []string, list2 []string) []string {
	seen := make(map[string]int)
	var res []string
	min := math.MaxInt

	for i, each := range list1 {
		seen[each] = i
	}

	for i, each := range list2 {
		if j, ok := seen[each]; ok {
			temp := i + j
			if temp < min {
				min = temp
				res = []string{each}
			} else if temp == min {
				res = append(res, each)
			}
		}
	}
	return res
}

func canPlaceFlower(flowerBed []int, n int) bool {
	if n == 0 {
		return false
	}
	if len(flowerBed) == 1 {
		return flowerBed[0] == 0
	}
	if flowerBed[0] == 0 && flowerBed[1] == 0 {
		flowerBed[0] = 1
		n--
	}
	for i := 1; i < len(flowerBed)-1 && n != 0; i++ {
		if flowerBed[i] == 0 && flowerBed[i-1] == 0 && flowerBed[i+1] == 0 {
			flowerBed[i] = 1
			n--
		}
	}
	if n != 0 && flowerBed[len(flowerBed)-1] == 0 && flowerBed[len(flowerBed)-2] == 0 {
		flowerBed[len(flowerBed)-1] = 1
		n--
	}
	return n == 0
}

func mergeTrees(root1 *util.TreeNode, root2 *util.TreeNode) *util.TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	res := &util.TreeNode{Val: root1.Val + root2.Val}

	res.Left = mergeTrees(root1.Left, root2.Left)
	res.Right = mergeTrees(root1.Right, root2.Right)
	return res
}

func maximumProducts(nums []int) int {
	slices.Sort(nums)

	n := len(nums)

	max := nums[n-1] * nums[n-2] * nums[n-3]

	if nums[0] < 0 && nums[1] < 0 {
		temp := nums[0] * nums[1] * nums[n-1]
		if temp > n {
			max = temp
		}
	}
	return max
}

func averageOfLevels(root *util.TreeNode) []float64 {
	res := []float64{}
	if root == nil {
		return res
	}

	temp := []*util.TreeNode{}
	temp = append(temp, root)

	for len(temp) > 0 {
		n := len(temp)
		sum := 0.0

		for i := 0; i < n; i++ {
			j := temp[0]
			sum += float64(j.Val)
			temp = temp[1:]

			if j.Left != nil {
				temp = append(temp, j.Left)
			}
			if j.Right != nil {
				temp = append(temp, j.Right)
			}
		}

		res = append(res, float64(sum/float64(n)))
	}
	return res
}

func findMaxAverage(nums []int, k int) float64 {
	res := math.MinInt
	sum := 0

	for i, each := range nums {
		sum += each
		if i >= k-1 {
			if i != k-1 {
				sum = sum - nums[i-k]
			}
			if res < sum {
				res = sum
			}
		}
	}
	return float64(res) / float64(k)
}

func findErrorNums(nums []int) []int {
	n := len(nums)
	seen := make([]int, n+1)
	var double int

	for _, val := range nums {
		seen[val]++
		if seen[val] == 2 {
			double = val
		}
	}

	for i := 1; i <= n; i++ {
		if seen[i] == 0 {
			return []int{double, i}
		}
	}
	return nil
}

func findTarget(root *util.TreeNode, k int) bool {
	var dfs func(root *util.TreeNode, k int, m map[int]bool) bool

	dfs = func(root *util.TreeNode, k int, m map[int]bool) bool {
		if root == nil {
			return false
		}

		if m[k-root.Val] {
			return true
		}
		m[root.Val] = true
		return dfs(root.Left, k, m) || dfs(root.Right, k, m)
	}
	return dfs(root, k, make(map[int]bool))
}

func judgeCirce(moves string) bool {
	x, y := 0, 0

	for _, each := range moves {
		switch each {
		case 'L':
			x++
		case 'R':
			x--
		case 'U':
			y++
		case 'D':
			y--
		}
	}
	return x == 0 && y == 0
}

func findSecondMinValue(root *util.TreeNode) int {
	if root == nil {
		return -1
	}

	min := root.Val
	res := math.MaxInt
	queue := []*util.TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Val > min && node.Val < res {
			node.Val = res
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	if res != math.MaxInt {
		return res
	}
	return -1
}

func findLengthOfLCIS(nums []int) int {
	res, count := 1, 1

	for r := 1; r < len(nums); r++ {
		if nums[r] > nums[r-1] {
			count++
			if count > res {
				res = count
			}
		} else {
			count = 1
		}
	}
	return res
}

func validPalindrome(s string) bool {
	var IsPalindrome func(s string, l, r int) (bool, int, int)

	IsPalindrome = func(s string, l, r int) (bool, int, int) {
		for l < r {
			if s[l] != s[r] {
				return false, l, r
			}
			l++
			r--
		}
		return true, 0, 0
	}

	valid, l, r := IsPalindrome(s, 0, len(s)-1)
	if valid {
		return true
	}
	if valid, _, _ := IsPalindrome(s, l+1, r); valid {
		return true
	}
	if valid, _, _ := IsPalindrome(s, 0, r-1); valid {
		return true
	}
	return false
}

func calPoints(operations []string) int {
	temp := make([]int, 0)

	for _, each := range operations {
		switch each {
		case "+":
			temp = append(temp, temp[len(temp)-1]+temp[len(temp)-2])
		case "D":
			temp = append(temp, temp[len(temp)-1]*2)
		case "C":
			temp = temp[:len(temp)-1]
		default:
			val, _ := strconv.Atoi(each)
			temp = append(temp, val)
		}
	}

	sum := 0
	for _, val := range temp {
		sum += val
	}
	return sum
}

func hasAlternativesBit(n int) bool {
	prevr := -1

	for n > 0 {
		r := n % 2
		n /= 2

		if r == prevr {
			return false
		}
		prevr = r
	}
	return true
}

func countBinarySubstring(s string) int {
	p, q, res := 1, 0, 0

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			p++
		} else {
			p, q = 1, p
		}
		if q >= p {
			res++
		}
	}
	return res
}

func searchBST(root *util.TreeNode, val int) *util.TreeNode {
	temp := root

	for temp != nil {
		if val == temp.Val {
			return temp
		} else if val < temp.Val {
			temp = temp.Left
		} else if val > temp.Val {
			temp = temp.Right
		}
	}
	return nil
}

func findShortestSubarray(nums []int) int {
	res := 1
	starts := make(map[int]int)
	counters, maxCounters := make(map[int]int), 1

	for i, num := range nums {
		if _, ok := starts[num]; ok {
			counters[num]++
			if counters[num] > maxCounters {
				maxCounters = counters[num]
				res = i - starts[num] + 1
			} else if counters[num] == maxCounters && i-starts[num] < res {
				res = i - starts[num] + 1
			}
		} else {
			starts[num], counters[num] = i, 1
		}
	}
	return res
}

// func toLower(s string) string {
// 	return strings.ToLower(s)
// }

func toLower(s string) string {
	temp := []byte(s)

	for i := 0; i < len(temp); i++ {
		if temp[i] <= 'Z' && temp[i] >= 'A' {
			temp[i] = temp[i] + 32
		}
	}
	return string(temp)
}

func isOneBitsCharacter(bits []int) bool {
	if len(bits) == 1 {
		return true
	}

	for i, j := 0, 1; j < len(bits) && i < len(bits)-1; {
		if i == len(bits)-2 {
			if bits[i] == 0 {
				return true
			}
			return false
		}
		if bits[i] == 0 {
			i++
			j++
		} else {
			i = i + 2
			j = j + 2
		}
	}
	return true
}

func pivotIndex(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	leftSum := 0
	for i, num := range nums {
		if leftSum == sum-leftSum-num {
			return i
		}
		leftSum += num
	}
	return -1
}

func selfDividingNumbers(left int, right int) []int {
	res := []int{}
	var isSelfDivising func(k int) bool

	isSelfDivising = func(k int) bool {
		for n := k; n > 0; n /= 10 {
			d := n % 10
			if d == 0 || k%d != 0 {
				return false
			}
		}
		return true
	}
	for left <= right {
		if isSelfDivising(left) {
			res = append(res, left)
		}
		left++
	}
	return res
}

func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)-1

	for l < r {
		m := l + (r-l)/2
		if letters[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	if letters[l] <= target {
		return letters[0]
	}
	return letters[l]
}

func minCostClimbingStairs(cost []int) int {
	n1, n2 := 0, 0

	for _, each := range cost {
		n1, n2 = n2, util.Min(n1, n2)+each
	}
	return util.Min(n1, n2)
}

func dominantIndex(nums []int) int {
	max := math.MinInt32
	secondMax := math.MinInt32
	maxIndex := 0

	for i, num := range nums {
		if num > max {
			secondMax = max
			max = num
			maxIndex = i
		} else if num > secondMax {
			secondMax = num
		}
	}
	if max >= secondMax*2 {
		return maxIndex
	}
	return -1
}

func countPrimeSetBits(left int, right int) int {
	var countBits func(n int) int

	countBits = func(n int) int {
		cnt := 0
		for n != 0 {
			n = n & (n - 1)
			cnt++
		}
		return cnt
	}

	primes := map[int]bool{
		2:  true,
		3:  true,
		5:  true,
		7:  true,
		11: true,
		13: true,
		17: true,
		19: true,
	}
	c := 0
	for i := left; i <= right; i++ {
		s := countBits(i)
		if primes[s] {
			c++
		}
	}
	return c
}

func numJewelInStone(jewels string, stones string) int {
	var cnt int
	// for _, stone := range stones {
	// 	if strings.Contains(jewels, string(stone)) {
	// 		cnt++
	// 	}
	// }

	for _, jewel := range jewels {
		for _, stone := range stones {
			if jewel == stone {
				cnt++
			}
		}
	}
	return cnt
}

func rotateString(s string, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
	// arr := []rune(s)
	//
	// for _, arr1 := range arr {
	// 	arr = arr[1:]
	// 	arr = append(arr, arr1)
	// 	if string(arr) == goal {
	// 		return true
	// 	}
	// }
	// return false
}

func uniqueMorseRepresentation(words []string) int {
	morseMap := map[byte]string{
		'a': ".-",
		'b': "-...",
		'c': "-.-.",
		'd': "-..",
		'e': ".",
		'f': "..-.",
		'g': "--.",
		'h': "....",
		'i': "..",
		'j': ".---",
		'k': "-.-",
		'l': ".-..",
		'm': "--",
		'n': "-.",
		'o': "---",
		'p': ".--.",
		'q': "--.-",
		'r': ".-.",
		's': "...",
		't': "-",
		'u': "..-",
		'v': "...-",
		'w': ".--",
		'x': "-..-",
		'y': "-.--",
		'z': "--..",
	}
	morseSet := make(map[string]int)
	for _, word := range words {
		morseWord := ""
		for _, letter := range word {
			morseWord += morseMap[byte(letter)]
		}
		morseSet[morseWord] = 0
	}
	return len(morseSet)
}

func numberOfLines(width []int, s string) []int {
	pixel := 0
	lines := 1
	for _, each := range s {
		pixel += width[each-'a']
		// to get the index that is why we are using each-'a', if the letter is a then a - 'a' = 0 which is index 0
		if pixel > 100 {
			lines++
			pixel = width[each-'a']
		}
	}
	return []int{lines, pixel}
}

func mostCommonWord(paragraph string, banned []string) string {
	m := make(map[string]int)
	m2 := make(map[string]int)

	words := strings.FieldsFunc(paragraph, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	for _, ban := range banned {
		m[ban] = 1
	}

	for _, word := range words {
		s := strings.ToLower(word)
		m2[s]++
	}

	res := ""
	max := 0

	for i, l := range m2 {
		if l > max && m[i] == 0 && i != "" {
			max = l
			res = i
		}
	}
	return res
}

func shortestToChar(s string, c byte) []int {
	res := make([]int, len(s))
	c_pos := -len(s)

	for i := 0; i < len(s); i++ {
		if s[i] == c {
			c_pos = i
		}
		res[i] = i - c_pos
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			c_pos = i
		}
		res[i] = util.Min(res[i], util.Abs(i-c_pos))
	}
	return res
}

func toGoatLatin(sentence string) string {
	res := ""
	words := strings.Split(sentence, " ")
	vowels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
		"A": true,
		"E": true,
		"O": true,
		"I": true,
		"U": true,
	}

	for i, word := range words {
		if vowels[string(word[0])] {
			word += "ma"
		} else {
			word = word[1:] + string(word[0]) + "ma"
		}
		res += word + strings.Repeat("a", i+1) + " "
	}
	return strings.TrimSpace(res)
}

func largeGroupPosition(s string) [][]int {
	res := [][]int{}
	s = s + " "
	l := 0

	for r := 1; r < len(s); r++ {
		if s[r] != s[r-1] {
			if r-l >= 3 {
				res = append(res, []int{l, r - 1})
			}
			l = r
		}
	}
	return res
}

func flipAndInvertImage(image [][]int) [][]int {
	flip := func(a int) int {
		if a == 0 {
			return 1
		}
		return 0
	}

	for r := 0; r < len(image); r++ {
		l, i := 0, len(image[r])-1

		for l < i {
			image[r][l], image[r][i] = flip(image[r][i]), flip(image[r][l])
			l++
			i--
		}
		if l == i {
			image[r][i] = flip(image[r][i])
		}
	}
	return image
}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	l, r := util.Max(rec1[0], rec2[0]), util.Min(rec1[2], rec2[2])
	up, down := util.Min(rec1[3], rec2[3]), util.Max(rec1[1], rec2[1])

	if r > l && up > down {
		return true
	}
	return false
}

func computeArea(a, b, c, d, e, f, g, h int) int {
	ABCDrec := (d - b) * (c - a)
	EFGHrec := (h - f) * (g - e)
	sum := ABCDrec + EFGHrec

	l, r := util.Max(a, e), util.Min(g, c)
	up, down := util.Min(d, h), util.Max(f, b)

	if r > l && up > down {
		sum -= (r - l) * (up - down)
	}
	return sum
}

func backspaceCompare(s string, t string) bool {
	processString := func(char []rune) int {
		k := 0
		for _, c := range char {
			if c != '#' {
				char[k] = c
				k++
			} else if k > 0 {
				k--
			}
		}
		return k
	}

	sRunes := []rune(s)
	tRunes := []rune(t)
	k := processString(sRunes)
	p := processString(tRunes)

	if k != p {
		return false
	}

	for i := 0; i < k; i++ {
		if sRunes[i] != tRunes[i] {
			return false
		}
	}
	return true
}

func lemonadeChange(bills []int) bool {
	five := 0
	ten := 0
	twenty := 0

	for _, bill := range bills {
		switch bill {
		case 5:
			five++
		case 10:
			if five >= 1 {
				ten++
				five--
			} else {
				return false
			}
		case 20:
			if five >= 1 && ten >= 1 {
				twenty++
				five--
				ten--
			} else if five >= 3 {
				twenty++
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

func transposeMatrix(matrix [][]int) [][]int {
	rows, cols := len(matrix), len(matrix[0])
	res := make([][]int, cols)

	for r := 0; r < cols; r++ {
		res[r] = make([]int, rows)
		for c := 0; c < rows; c++ {
			res[r][c] = matrix[c][r]
		}
	}
	return res
}

func binaryGap(n int) int {
	bins := fmt.Sprintf("%b", n)
	count, res := 0, 0

	for _, bin := range bins {
		if bin == '0' {
			count++
		} else {
			if res < count {
				res = count
			}
			count = 1
		}
	}
	return res
}

func leafSimilarTrees(root1 *util.TreeNode, root2 *util.TreeNode) bool {
	var dfs func(root *util.TreeNode) []int

	dfs = func(root *util.TreeNode) []int {
		if root == nil {
			return []int{}
		}
		if root.Left == nil && root.Right == nil {
			return []int{root.Val}
		}
		return append(dfs(root.Left), dfs(root.Right)...)
	}

	return reflect.DeepEqual(dfs(root1), dfs(root2))
}

func buddyStrings(s, goal string) bool {
	sn, gn := len(s), len(goal)

	if sn != gn {
		return false
	}

	diff := []int{}
	for i := range s {
		if s[i] != goal[i] {
			diff = append(diff, i)

			if len(diff) > 2 {
				return false
			}
		}
	}

	if len(diff) == 2 {
		return s[diff[0]] == goal[diff[1]] && s[diff[1]] == goal[diff[0]]
	}

	if len(diff) == 1 {
		return false
	}

	sameMap := make(map[byte]int)
	for i := range s {
		_, exists := sameMap[s[i]]
		if exists {
			return true
		}
		sameMap[s[i]]++
	}
	return false
}

func projectionArea(grid [][]int) int {
	res := 0

	for i := range grid {
		maxRow, maxCol := 0, 0
		// this col part was able to work because they are a 3d array size instead it is meant to be range in grid[i]
		for j := range grid {
			if grid[j][i] > maxRow {
				maxRow = grid[j][i]
			}
			if grid[i][j] > maxCol {
				maxCol = grid[i][j]
			}

			if grid[i][j] != 0 {
				res++
			}
		}
		res += maxCol + maxRow
	}
	return res
}

func middleNode(head *util.ListNode) *util.ListNode {
	s, f := head, head

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}
	return s
}

func uncommonFromSentence(s1 string, s2 string) []string {
	word := strings.Split(s1, " ")
	word = append(word, strings.Split(s2, " ")...)
	res := []string{}
	sameMap := make(map[string]int)

	for _, each := range word {
		sameMap[each]++
	}

	for k, v := range sameMap {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	aliceSum := 0
	bobSum := 0

	for _, n := range aliceSizes {
		aliceSum += n
	}

	for _, n := range bobSizes {
		bobSum += n
	}

	targetTotal := (aliceSum + bobSum) / 2
	aliceDiff := targetTotal - aliceSum

	for _, aliceBox := range aliceSizes {
		for _, bobBox := range bobSizes {
			if aliceBox+aliceDiff == bobBox {
				return []int{aliceBox, bobBox}
			}
		}
	}
	return []int{}
}

func isMonotonic(nums []int) bool {
	increasing := true
	decreasing := true

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			decreasing = false
		} else if nums[i] < nums[i-1] {
			increasing = false
		}
		if !increasing && !decreasing {
			return false
		}
	}
	return true
}

func increasingBst(root *util.TreeNode) *util.TreeNode {
	var dfs func(root, tail *util.TreeNode) *util.TreeNode

	dfs = func(root, tail *util.TreeNode) *util.TreeNode {
		if root == nil {
			return tail
		}
		res := dfs(root.Left, root)
		root.Left = nil
		root.Right = dfs(root.Right, tail)
		return res
	}
	return dfs(root, nil)
}

func sortArrayByParity(nums []int) []int {
	odds := []int{}
	evens := []int{}

	for _, num := range nums {
		if num%2 == 0 {
			evens = append(evens, num)
		} else {
			odds = append(odds, num)
		}
	}
	evens = append(evens, odds...)
	return evens
}

func reverseOnlyLetters(s string) string {
	res := []rune(s)
	l, r := 0, len(s)-1

	for l < r {
		if !unicode.IsLetter(res[l]) {
			l++
		} else if !unicode.IsLetter(res[r]) {
			r--
		} else {
			res[l], res[r] = res[r], res[l]
			l++
			r--
		}
	}
	return string(res)
}

func sortByParityII(nums []int) []int {
	l, r := 0, 1
	res := make([]int, len(nums))

	for _, n := range nums {
		if n%2 == 0 {
			res[l] = n
			l += 2
		} else {
			res[r] = n
			r += 2
		}
	}
	return res
}

// func numInUniqueEmails(emails []string) int {
// 	// this is the first metghod to do it and the efficient method using standard lib
// 	seen := make(map[string]int)
// 	for _, email := range emails {
// 		nameAndDomain := strings.Split(email, "@")
// 		splitName := strings.Split(nameAndDomain[0], "+")
// 		trimmedName := strings.ReplaceAll(splitName[0], ".", "")
// 		fullName := trimmedName + "@" + nameAndDomain[1]
// 		seen[fullName] = 1
// 	}
// 	return len(seen)
// }

func numInUniqueEmails(emails []string) int {
	seen := make(map[string]int)

	for i := 0; i < len(emails); i++ {
		var localName []byte
		j := 0
		for emails[i][j] != '@' && emails[i][j] != '+' {
			if emails[i][j] != '.' {
				localName = append(localName, emails[i][j])
			}
			j++
		}
		for emails[i][j] != '@' {
			j++
		}
		address := string(localName) + string(emails[i][j:])
		if _, exists := seen[address]; !exists {
			seen[address] = 1
		}
	}
	return len(seen)
}

func rangeSumBST(root *util.TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	} else if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	} else {
		return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
	}
}

func DIStringMatch(s string) []int {
	res := make([]int, len(s)+1)
	low, high := 0, len(s)

	for i, v := range s {
		if v == 'I' {
			res[i] = low
			low++
		} else {
			res[i] = high
			high--
		}
	}
	res[len(s)] = low
	return res
}

func minDeletionSize(strs []string) int {
	res := 0

	for c := 0; c < len(strs[0]); c++ {
		prev := -1
		for r := 0; r < len(strs); r++ {
			if int(strs[r][c]) < prev {
				res++
				break
			}
			prev = int(strs[r][c])
		}
	}
	return res
}

func isAlienSorted(words []string, order string) bool {
	m := make(map[rune]int)

	for i, v := range order {
		m[v] = i
	}

	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		for j := 0; j < len(words[i]); j++ {
			if j >= len(w2) {
				return false
			}
			l := rune(w1[j])
			r := rune(w2[j])
			if l != r {
				if m[l] > m[r] {
					return false
				} else {
					break
				}
			}
		}
	}
	return true
}

func repeatedNtimes(nums []int) int {
	m := make(map[int]int)

	for _, num := range nums {
		if m[num] == 1 {
			return num
		}
		m[num]++
	}
	return -1
}

func largestPerimeter(nums []int) int {
	sort.Ints(nums)

	for i := len(nums) - 3; i >= 0; i-- {
		if nums[i]+nums[i+1] > nums[i+2] {
			return nums[i] + nums[i+1] + nums[i+2]
		}
	}
	return 0
}

func isUnivalTree(root *util.TreeNode) bool {
	q := []*util.TreeNode{root}

	for len(q) > 0 {
		for i := 0; i < len(q); i++ {
			n := q[0]
			q := q[1:]
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
			if n.Val != root.Val {
				return false
			}
		}
	}
	return true
}

func addToArrayform(nums []int, k int) []int {
	i := 1

	for k > 0 {
		if len(nums) >= i {
			k += nums[len(nums)-i]
			nums[len(nums)-i] = k % 10
		} else {
			nums = append([]int{k % 10}, nums...)
		}
		i++
		k = k / 10
	}
	return nums
}

func sortedSquares(nums []int) []int {
	i, l, r := len(nums)-1, 0, len(nums)-1
	res := make([]int, len(nums))

	for l <= r {
		lsq := nums[l] * nums[l]
		rsq := nums[r] * nums[r]

		if lsq > rsq {
			res[i] = lsq
			l++
		} else {
			res[i] = rsq
			r--
		}
		i--
	}
	return res
}

func findJudge(trust [][]int, n int) int {
	trustedcount := make([]int, n+1)

	for _, trustInfo := range trust {
		truster := trustInfo[0]
		trusted := trustInfo[1]

		trustedcount[truster]--
		trustedcount[trusted]++
	}
	for p := 1; p <= n; p++ {
		if trustedcount[p] == n-1 {
			return p
		}
	}
	return -1
}

func isCousins(root *util.TreeNode, x, y int) bool {
	var xHeight, yHeight, xParent, yParent int

	var dfs func(r *util.TreeNode, h int)
	dfs = func(r *util.TreeNode, h int) {
		if r == nil {
			return
		}

		dfs(r.Left, h+1)
		dfs(r.Right, h+1)

		if r.Left != nil {
			if r.Left.Val == x {
				xHeight = h + 1
				xParent = r.Val
			}
			if r.Left.Val == y {
				yHeight = h + 1
				yParent = r.Val
			}

			if r.Right != nil {
				if r.Right.Val == x {
					xHeight = h + 1
					xParent = r.Val
				}
				if r.Right.Val == y {
					yHeight = h + 1
					yParent = r.Val
				}
			}
		}
	}

	dfs(root, 0)
	if xHeight == yHeight {
		if xHeight > 1 && yHeight > 1 {
			if xParent != yParent {
				return true
			}
		}
	}
	return false
}

func commonChars(words []string) []string {
	res := []string{}
	seen := make(map[rune]int)

	for _, each := range words[0] {
		seen[each]++
	}

	for i := 1; i <= len(words)-1; i++ {
		seenN := make(map[rune]int)

		for _, char := range words[i] {
			seenN[char]++
		}

		for k, v := range seen {
			if _, exists := seenN[k]; !exists {
				seen[k] = 0
			} else {
				seen[k] = util.Min(v, seenN[k])
			}
		}
	}
	for k, v := range seen {
		if v != 0 {
			for i := 0; i < v; i++ {
				res = append(res, string(k))
			}
		}
	}
	return res
}

func canThreePartsEqualSum(arr []int) bool {
	sum := 0

	for _, num := range arr {
		sum += num
	}

	if sum%3 != 0 {
		return false
	}

	partSum, currentSum, count := sum/3, 0, 0

	for i := 0; i < len(arr); i++ {
		currentSum += arr[i]
		if currentSum == partSum {
			currentSum = 0
			count++
		}
	}
	return count >= 3
}

func prefixesDivBy5(nums []int) []bool {
	res := make([]bool, len(nums))
	rem := 0

	for i, num := range nums {
		rem = (rem*2 + num) % 5
		if rem == 0 {
			res[i] = true
		}
	}
	return res
}
