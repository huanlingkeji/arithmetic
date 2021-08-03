package dp

/*
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

*/

// 状态: 打劫的家
// 选择：打or不打

// dp的定义：走到第i家时打劫的最高金额

/*
由于不可以在相邻的房屋闯入，所以在当前位置 n 房屋可盗窃的最大值，要么就是 n-1 房屋可盗窃的最大值，
要么就是 n-2 房屋可盗窃的最大值加上当前房屋的值，二者之间取最大值

*/

// 状态转移方程： dp[n] = max(num[n-2]+dp[n-2],dp[n-1] )   dp[1]= nums[1] dp[0]= 0

func rob(nums []int) int {
	lth := len(nums)
	if lth < 1 {
		return 0
	}
	if lth == 1 {
		return nums[0]
	}
	maxv := max(nums[0], nums[1])
	f2 := 0
	f1 := nums[0]
	for i := 2; i <= lth; i++ {
		maxv = max(f1, f2+nums[i-1])
		f2 = f1
		f1 = maxv
	}
	return maxv
}

/*
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋
和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

题解：环状排列意味着第一个房子和最后一个房子中只能选择一个偷窃，因此可以把此环状排列房间问题约化为两个单排排列房间子问题

*/

func rob2(nums []int) int {
	lth := len(nums)
	if lth == 1 {
		return nums[0]
	}
	fct := func(nums []int) int {
		lth := len(nums)
		if lth < 1 {
			return 0
		}
		if lth == 1 {
			return nums[0]
		}
		f2 := 0
		f1 := 0
		for i := 0; i < lth; i++ {
			f1, f2 = max(f1, f2+nums[i]), f1
		}
		return f1
	}
	return max(fct(nums[:lth-1]), fct(nums[1:]))
}