package algo

/*
"12"
"223"
"1234"
"1"
"0"
"01"
"0123"
"1020"
"102030"
"123456789"
"2101"
"1123"

期望：
2
3
3
1
0
0
0
1
0
3
1
bad case
"2101"
期望：1
输出：2
*/

func numDecodings(s string) int {
	n := len(s)
	if s[0] == '0' {
		return 0
	}
	if n == 1 {
		return 1
	}
	var dp = make([]int, n+1) // dp[0]=0; res=dp[n]
	dp[0] = 1                 // 没写出来的重要点 空字符串也算一种？？！满足"12"的2的需求罢了！

	if s[0] != '0' {
		dp[1] = 1
	}

	for i := 2; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}
		num := (int(s[i-2])-'0')*10 + int(s[i-1]-'0')
		if s[i-2] != '0' && num <= 26 {
			dp[i] = dp[i] + dp[i-2]
		}
	}

	return dp[n]
}
