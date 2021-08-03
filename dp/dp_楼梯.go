package dp

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。
*/

// 状态: 楼梯剩余的级数
// 选择：爬一级or两级

// dp的定义：第n级台阶时有多少种方法

// 状态转移方程： dp[n] = dp[n-1] + dp[n-2]   dp[1]= 1  dp[0]= 0

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	// 空间可优化
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}

/*
数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost[i]（下标从 0 开始）。

每当你爬上一个阶梯你都要花费对应的体力值，一旦支付了相应的体力值，你就可以选择向上爬一个阶梯或者爬两个阶梯。

请你找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。
*/

// 状态: 楼梯剩余的级数
// 选择：爬一级or两级 从1or0开始

// dp的定义：第n级台阶时消耗的最小体力值

// 状态转移方程： dp[n] = min( dp[n-1] +cost[n-1]), cost[n-2] + dp[n-2])   dp[1]= 0 dp[0]= 0

// 最小花费爬楼梯
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n < 2 {
		return 0
	}
	f1, f2 := 0, 0
	minv := 0
	for i := 2; i <= n; i++ {
		// 如果不再使用到的内存使用省略
		minv = min(f2+cost[i-2], cost[i-1]+f1)
		f2 = f1
		f1 = minv
	}
	return minv
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
