package medium

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGroupAnagram(t *testing.T) {
	tests := []struct {
		strs []string
		want [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.strs), func(t *testing.T) {
			got := GroupAnagram(tt.strs)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.nums, tt.k), func(t *testing.T) {
			got := topKFrequent(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := productExceptSelf(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := longestConsecutive(tt.nums)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoSumII(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.nums, tt.target), func(t *testing.T) {
			got := twoSumII(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, [][]int{}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := threeSum(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.height), func(t *testing.T) {
			got := maxArea(tt.height)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbb", 1},
		{"pwwkew", 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.s), func(t *testing.T) {
			got := lengthOfLongestSubString(tt.s)
			if got != tt.want {
				t.Fatalf("%v, %v", got, tt.want)
			}
		})
	}
}

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want int
	}{
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.s, tt.k), func(t *testing.T) {
			got := characterReplacement(tt.s, tt.k)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMIn(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := findMin(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("%v, %v", got, tt.want)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.nums, tt.target), func(t *testing.T) {
			got := search(tt.nums, tt.target)
			if got != tt.want {
				t.Fatalf("%v, %v", got, tt.want)
			}
		})
	}
}

// no test for reorderLIst

// no test for removeNthNode

// no test for lowest common ancestor

// no test for level order

// no test for kth smallest element

// no test for isValid BST

// no test for build tree

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		want       [][]int
	}{
		{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v, %v", tc.candidates, tc.target), func(t *testing.T) {
			got := combinationSum(tc.candidates, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("%v, %v", got, tc.want)
			}
		})
	}
}

func TestCombinationSumII(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		want       [][]int
	}{
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		{[]int{2, 5, 2, 1, 2}, 5, [][]int{{1, 2, 2}, {5}}},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v, %v", tc.candidates, tc.target), func(t *testing.T) {
			got := combinationSumII(tc.candidates, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("%v, %v", got, tc.want)
			}
		})
	}
}

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]byte
		expected int
	}{
		{
			name: "Single island",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			name: "Multiple islands",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
		{
			name: "No islands",
			grid: [][]byte{
				{'0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 0,
		},
		{
			name: "All islands",
			grid: [][]byte{
				{'1', '1'},
				{'1', '1'},
			},
			expected: 1,
		},
		{
			name:     "Empty grid",
			grid:     [][]byte{},
			expected: 0,
		},
		{
			name: "Single cell island",
			grid: [][]byte{
				{'1'},
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numIslands(tt.grid)
			if result != tt.expected {
				t.Errorf("numIslands() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// no test for cloned Graph

func TestPacificAtlanctic(t *testing.T) {
	tests := []struct {
		heights [][]int
		want    [][]int
	}{
		{
			[][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			[][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}},
		},
		{[][]int{{1}}, [][]int{{0, 0}}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.heights), func(t *testing.T) {
			got := pacificAtlantic(tt.heights)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestCanFinish(t *testing.T) {
	tests := []struct {
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{2, [][]int{{1, 0}}, true},
		{2, [][]int{{1, 0}, {0, 1}}, false},
		{4, [][]int{{1, 0}, {2, 1}, {3, 2}}, true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.numCourses, tt.prerequisites), func(t *testing.T) {
			got := canFinish(tt.numCourses, tt.prerequisites)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRob(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{1, 2, 3, 1}, 4},
		{[]int{1, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := rob(tt.nums)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRobII(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 2}, 3},
		{[]int{1, 2, 3, 1}, 4},
		{[]int{1, 2, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := robII(tt.nums)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.s), func(t *testing.T) {
			got := longestPalindrome(tt.s)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPalindromicString(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abc", 3},
		{"aaa", 6},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.s), func(t *testing.T) {
			got := palindromicSubstring(tt.s)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumDecoding(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"12", 2},
		{"226", 3},
		{"06", 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.s), func(t *testing.T) {
			got := numDecodings(tt.s)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins  []int
		amount int
		want   int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.coins, tt.amount), func(t *testing.T) {
			got := coinChange(tt.coins, tt.amount)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoinChangeII(t *testing.T) {
	tests := []struct {
		amount int
		coins  []int
		want   int
	}{
		{5, []int{1, 2, 5}, 4},
		{3, []int{2}, 0},
		{10, []int{10}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.coins, tt.amount), func(t *testing.T) {
			got := coinChangeII(tt.amount, tt.coins)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := maxProduct(tt.nums)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordBreak(t *testing.T) {
	tests := []struct {
		s        string
		wordDict []string
		want     bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.s, tt.wordDict), func(t *testing.T) {
			got := wordBreak(tt.s, tt.wordDict)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLengthOfLIS(t *testing.T) {
	test := []struct {
		nums []int
		want int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
	}
	for _, tt := range test {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := lengthOfLIS(tt.nums)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		m, n, want int
	}{
		{3, 7, 28},
		{3, 2, 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v and %v", tt.m, tt.n), func(t *testing.T) {
			got := uniquePaths(tt.m, tt.n)
			if got != tt.want {
				t.Fatalf("Got %v, Want %v", got, tt.want)
			}
		})
	}
}

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		text1, text2 string
		want         int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v and %v", tt.text1, tt.text2), func(t *testing.T) {
			got := longestCommonSubsequence(tt.text1, tt.text2)
			if got != tt.want {
				t.Fatalf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxSubArr(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{-1}, -1},
		{[]int{5, 4, -1, 7, 8}, 23},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("got %v", tt.nums), func(t *testing.T) {
			got := maxSubArray(tt.nums)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanJump(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := canJump(tt.nums)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{[][]int{{1, 3}, {6, 9}, {11, 15}}, []int{2, 7}, [][]int{{1, 9}, {11, 15}}},
		{
			[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			[]int{4, 8},
			[][]int{{1, 2}, {3, 10}, {12, 16}},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.intervals, tt.newInterval), func(t *testing.T) {
			got := insert(tt.intervals, tt.newInterval)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.intervals), func(t *testing.T) {
			got := merge(tt.intervals)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEraseOverLapIntervals(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      int
	}{
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
		{[][]int{{1, 2}, {1, 2}, {1, 2}}, 2},
		{[][]int{{1, 2}, {2, 3}}, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.intervals), func(t *testing.T) {
			got := eraseOverLapIntervals(tt.intervals)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   [][]int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		{
			[][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 1, 16}},
			[][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.matrix), func(t *testing.T) {
			rotate(tt.matrix)
			if !reflect.DeepEqual(tt.matrix, tt.want) {
				t.Fatalf("got %v, want %v", tt.matrix, tt.want)
			}
		})
	}
}

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   []int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{
			[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.matrix), func(t *testing.T) {
			got := spiralOrder(tt.matrix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetZeroes(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   [][]int
	}{
		{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{
			[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			[][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.matrix), func(t *testing.T) {
			setZeroes(tt.matrix)
			if !reflect.DeepEqual(tt.matrix, tt.want) {
				t.Fatalf("got %v, want %v", tt.matrix, tt.want)
			}
		})
	}
}

func TestGetSum(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{2, 3, 5},
		{4, 8, 12},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.a, tt.b), func(t *testing.T) {
			got := getSum(tt.a, tt.b)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1}, 1, []int{1}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.nums, tt.k), func(t *testing.T) {
			got := maxSlidingWindow(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinWindows(t *testing.T) {
	tests := []struct {
		s, t, want string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.s, tt.t), func(t *testing.T) {
			got := minWindow(tt.s, tt.t)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		// {1, []string{"()"}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.n), func(t *testing.T) {
			got := generateParenthesis(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidSudoku(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		want  bool
	}{
		{
			name: "Valid board 1",
			board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			want: true,
		},
		{
			name: "Invalid board - duplicate in row",
			board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'5', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidSudoku(tt.board)
			if got != tt.want {
				t.Fatalf("%s: got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestTrap(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.height), func(t *testing.T) {
			got := trap(tt.height)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		temperatures []int
		want         []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{[]int{30, 60, 90}, []int{1, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.temperatures), func(t *testing.T) {
			got := dailyTemperatures(tt.temperatures)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindsDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
		{[]int{3, 3, 3, 3, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := FindsDuplicates(tt.nums)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeastInterval(t *testing.T) {
	tests := []struct {
		tasks []byte
		n     int
		want  int
	}{
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2, 8},
		{[]byte{'A', 'C', 'A', 'B', 'D', 'B'}, 1, 6},
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 3, 10},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.tasks, tt.n), func(t *testing.T) {
			got := leastInterval(tt.tasks, tt.n)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubsets(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 2}, {1, 3}, {1}, {2, 3}, {2}, {3}, {}}},
		{[]int{0}, [][]int{{0}, {}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := subsets(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermute(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{[]int{0, 1}, [][]int{{0, 1}, {1, 0}}},
		{[]int{1}, [][]int{{1}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := permute(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubsetsII(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{[]int{1, 2, 2}, [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}},
		{[]int{0}, [][]int{{}, {0}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := subsetsII(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExist(t *testing.T) {
	tests := []struct {
		board [][]byte
		word  string
		want  bool
	}{
		{
			[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			"ABCCED",
			true,
		},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE", true},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCD", false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.board, tt.word), func(t *testing.T) {
			got := exist(tt.board, tt.word)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}
