package dp

// 斐波那契数列
// 暴力递归 return f(n-1)+f(n-2)
// 备忘录模式 if 有记录 return record;else return return f(n-1)+f(n-2) （自顶向下分解，自底向上回溯求解）
// 直接自底向上求解 dp[i] = dp[i-1]+ dp[i-2]
// 空间优化dp[i]只依赖dp[i-1]、dp[i-2]

// 解决动态规划的一般解题思路：明确 base case -> 明确「状态」-> 明确「选择」 -> 定义 dp 数组/函数的含义。
/*
# 初始化 base case
dp[0][0][...] = base
# 进行状态转移
for 状态1 in 状态1的所有取值：
    for 状态2 in 状态2的所有取值：
        for ...
            dp[状态1][状态2][...] = 求最值(选择1，选择2...)
*/

// 优化 状态压缩
