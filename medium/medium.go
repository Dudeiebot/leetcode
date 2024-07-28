package medium

import (
	"slices"

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

// chip in characterReplacement here
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
