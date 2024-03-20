package dp

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 动态规划
 * @param array int整型一维数组
 * @return int整型
 */
func FindGreatestSumOfSubArray(array []int) int {
	// write code here
	if len(array) == 0 {
		return 0
	}

	dp := make([]int, len(array))
	dp[0] = array[0]

	maxNum := dp[0]
	for i := 1; i < len(array); i++ {
		dp[i] = max(array[i], array[i]+dp[i-1])
		if dp[i] > maxNum {
			maxNum = dp[i]
		}
	}

	return maxNum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
