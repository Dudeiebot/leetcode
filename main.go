package main

// import (
// 	"strings"
// 	"unicode"
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
