package medium

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
