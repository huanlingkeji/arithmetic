package locde

import "math"

// 获取两字符串的编辑距离

func getMinLen(s1, s2 string) int {
	// dp[i][j] s1[0..i]和s2[0..j]的最小编辑距离
	m, n := len(s1), len(s2)
	dp := make([][]int, len(s1)+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, len(s2)+1)
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i-1][j-1]+1, //替换
					dp[i-1][j]+1,   //添加
					dp[i][j-1]+1,   //删除
				)
			}
		}
	}
	return dp[len(s1)][len(s2)]
}

func min(items ...int) int {
	minV := math.MaxInt
	for _, v := range items {
		if v < minV {
			minV = v
		}
	}
	return minV
}
