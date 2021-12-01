package locde

import (
	"testing"
)

// dp问题
// 1 明确base case 初始化 max = nums[0]
// 2 明确状态 选择奇数/选择偶数 转为表达
// 3 明确转移方程  dp[i] = max (dp[i-1]+num[i],dp[i-2]) (i%2==0)
// dp[i] = max (dp[i-1]-num[i],dp[i-2]) (i%2==1)

func maxAlternatingSumC(nums []int) {
	// 	int n = nums.length;
	// 	long[][] f = new long[n][2];
	// 	//f[i][0] 表示至下标i的长度为奇数的子序列的最大交替和
	// 	//f[i][1] 表示至下标i的长度为偶数的子序列的最大交替和
	// 	f[0][0] = nums[0];
	// 	f[0][1] = 0;
	// 	for (int i = 1; i < n; i++) {
	// 	f[i][0] = Math.max(f[i - 1][0], f[i - 1][1] + nums[i]);
	// 	f[i][1] = Math.max(f[i - 1][1], f[i - 1][0] - nums[i]);
	// 	}
	// 	return f[n - 1][0];
}

func maxAlternatingSum(nums []int) int64 {
	t1, t2 := nums[0], 0
	for i := 1; i < len(nums); i++ {
		t1 = max(t1, t2+nums[i])
		t2 = max(t2, t1-nums[i])
	}
	return int64(t1)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func TestName(t *testing.T) {

}
