package dp

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 动态规划
 *
 * @param number int整型
 * @return int整型
 */
func JumpFloor(number int) int {
	// write code here
	dp := make([]int, number+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i < len(dp); i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[number]
}
