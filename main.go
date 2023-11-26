package main

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
func isAnagram(s, t string) bool {
	// make our map
	c_s := make(map[rune]int)
	c_t := make(map[rune]int)

	// add all the character individually to the map
	for _, char := range s {
		c_s[char]++
	}

	for _, char := range t {
		c_t[char]++
	}

	// range through our mao and check if the count of c_s matches the one of c_t
	for char, count := range c_s {
		if count != c_t[char] {
			return false
		}
	}
	return true
}
