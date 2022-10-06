package algo

/*
70. 爬楼梯
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

示例 1：
输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶

示例 2：
输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶
*/

// ClimbStairs 简单动态规划
// 转移方程 f(1)=1, f(2)=2, f(n) = f(n-1)+f(n-2)
func ClimbStairs(n int) int {
	if n < 1 {
		return -1
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		return ClimbStairs(n-1) + ClimbStairs(n-2)
	}
}

//ClimbStairsArr 动态规划+存储数组结果
// n>=44 提交会超时
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.8 MB, 在所有 Go 提交中击败了65.26%的用户
*/
func ClimbStairsArr(n int) int {
	arr := make([]int, n)
	if n < 1 {
		return -1
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		arr[0], arr[1] = 1, 2
		return ClimbStairsArr_(n-1, &arr) + ClimbStairsArr_(n-2, &arr)
	}
}

func ClimbStairsArr_(n int, arr *[]int) int {
	if (*arr)[n-1] == 0 {
		(*arr)[n-1] = ClimbStairsArr_(n-1, arr) + ClimbStairsArr_(n-2, arr)
	}
	return (*arr)[n-1]
}

// ClimbStairsDp 取消数组空间, 最优DP
func ClimbStairsDp(n int) int {
	if n <= 2 {
		return n
	}
	num1, num2 := 1, 2
	for i := 2; i < n; i++ {
		num1, num2 = num2, num1+num2
	}
	return num2
}
