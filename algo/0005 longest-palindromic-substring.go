package algo

/*
5. 最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。

示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

示例 2：
输入：s = "cbbd"
输出："bb"

提示：
1 <= s.length <= 1000
s 仅由数字和英文字母组成

test:"aacabdkacaa"
*/

func longestPalindrome(s string) string {
	strLen := len(s)
	if strLen <= 2 {
		if s[0] == s[strLen-1] {
			return s
		} else {
			return s[0:1]
		}
	}
	var dp [][]bool
	dp = make([][]bool, strLen)
	for i := 0; i < strLen; i++ {
		dp[i] = make([]bool, strLen)
	}

	// init dp
	left, right, maxLen := 0, 0, 1
	for i := 0; i < strLen; i++ {
		dp[i][i] = true
		if i+1 < strLen {
			if s[i] == s[i+1] {
				dp[i][i+1] = true
				left, right, maxLen = i, i+1, 2
			} else {
				dp[i][i+1] = false
			}
		}
	}

	for p := 2; p < strLen; p++ {
		for i := 0; i+p < strLen; i++ {
			j := i + p
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				if j-i+1 > maxLen {
					left, right, maxLen = i, j, j-i+1
				}
			} else {
				dp[i][j] = false
			}
		}
	}
	return s[left : right+1]
}
