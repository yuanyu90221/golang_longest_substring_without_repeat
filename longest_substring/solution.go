package longest_substring

func lengthOfLongestSubstring(s string) int {
	hitMap := make(map[byte]int)
	if len(s) == 0 {
		return 0
	}
	result := 0
	left := 0
	right := 0
	for right < len(s) {
		if idx, ok := hitMap[s[right]]; ok && idx >= left { // 當遇到重覆值， 把左界移動到上一個重複的值位
			left = idx + 1
		}
		hitMap[s[right]] = right
		right++
		if right-left > result {
			result = right - left
		}
		if left+result >= len(s) {
			break
		}
	}

	return result
}
