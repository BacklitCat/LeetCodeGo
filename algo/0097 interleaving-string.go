package algo

// 边界调了很久 真的要小心！
func isInterleave(s1 string, s2 string, s3 string) bool {
	s1Len, s2Len, s3Len := len(s1), len(s2), len(s3)
	if s1Len+s2Len != s3Len {
		return false
	}
	var dp = make([][]bool, s1Len+1)
	for i := 0; i < s1Len+1; i++ {
		dp[i] = make([]bool, s2Len+1)
	}
	dp[0][0] = true

	for i := 0; i < s1Len && s1[i] == s3[i]; i++ {
		dp[i+1][0] = true
	}
	for j := 0; j < s2Len && s2[j] == s3[j]; j++ {
		dp[0][j+1] = true
	}

	for i := 0; i < s1Len; i++ {
		for j := 0; j < s2Len; j++ {
			if s3[i+j+1] == s1[i] {
				dp[i+1][j+1] = dp[i+1][j+1] || dp[i][j+1]
			}
			if s3[i+j+1] == s2[j] {
				dp[i+1][j+1] = dp[i+1][j+1] || dp[i+1][j]
			}
		}
	}
	return dp[s1Len][s2Len]
}
