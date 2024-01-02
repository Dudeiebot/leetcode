package main

// import "strings"
//
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

// We are finding the sum of 2 number in an array that formes a target and we return the indices of the numbers
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

func longestPrefix(s []string) string {
	if len(s) == 0 {
		return ""
	}

	firstLet := s[0]

	for i := 0; i < len(s); i++ {
		for _, words := range s[1:] {
			if i == len(words) || words[i] != firstLet[i] {
				return firstLet[0:i]
			}
		}
	}
	return firstLet
}
