package dp

import (
	"strconv"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 动态规划
 * @param nums string字符串 数字串
 * @return int整型
 */
func NumToStr(nums string) int {
	// write code here
	if len(nums) == 0 || nums[0] == '0' {
		return 0
	}

	dp := make([]int, len(nums)+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= len(nums); i++ {

		if nums[i-1] != '0' { // 单字节解码
			dp[i] += dp[i-1]
		}

		if j, _ := strconv.Atoi(nums[i-2 : i]); j >= 10 && j <= 26 { // 10-26 双字节解码
			dp[i] += dp[i-2]
		}
	}

	return dp[len(nums)]
}
