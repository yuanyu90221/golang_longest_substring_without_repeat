# longest substring without repeated character

Given a string s, find the length of the longest substring without repeating characters.

 
## Examples

```
Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
 

Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
```
## 解析

要找出最長不重複字元字串

首先可以透過 hashmap 紀錄已經出現過的字元以及位置

以 slide window 方式來做

就是先設定 左界為 0 ， 然後透過右界去擴張 當右界遇到 重複時 ， 更新左界到道上一個遇到重複值的位置+1

每次計算目前不重複的右界 - 左界

直到右界碰到邊界為止

## 程式碼

```golang
package longest_substring

func lengthOfLongestSubstring(s string) int {
	hitMap := make(map[byte]int)
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
	}
  return result
}
```