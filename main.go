package main

// duplicate num in a array

func DupNum(nums []int) bool {
	s := make(map[int]bool)

	for _, num := range nums {
		if s[num] {
			return true
		}
		s[num] = true
	}
	return false
}
