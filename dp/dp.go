package dp

// 看清题目给的初始值范围
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	f1 := 0
	f2 := 1
	ret := 0
	for i := 1; i < n; i++ {
		ret = f1 + f2
		f1 = f2
		f2 = ret
	}
	return ret
}

func tribonacci(n int) int {
	lth := 4
	fs := []int{
		0, 1, 1, 0,
	}

	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	for i := lth - 1; i <= n; i++ {
		fs[lth-1] = 0
		for j := 0; j < lth-1; j++ {
			fs[lth-1] += fs[j]
			fs[j] = fs[j+1]
		}
	}
	return fs[lth-1]
}

// 0 1背包问题
