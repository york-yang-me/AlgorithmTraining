package dp

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 动态规划
 * @param s string字符串
 * @return int整型
 */
func LengthOfLongestSubstring(s string) int {
	// write code here
	if len(s) == 0 {
		return 0
	}

	// dp[i] 表示以 s[i] 结尾的最长无重复子串长度
	dp := make([]int, len(s))

	// lastIndex 记录字符上一次出现的位置
	lastIndex := make(map[byte]int)

	dp[0] = 1           // 第一个字符子串长度为1
	lastIndex[s[0]] = 0 // 记录第一个字符的位置
	maxLen := 1         // 记录全局最长子串长度

	for i := 1; i < len(s); i++ {
		lastPos, found := lastIndex[s[i]]

		if found && lastPos >= i-dp[i-1] {
			// 如果字符在 dp[i-1] 子串内出现过
			dp[i] = i - lastPos
		} else {
			// 如果没有出现过或者不在 dp[i-1] 子串内
			dp[i] = dp[i-1] + 1
		}

		// 更新字符位置
		lastIndex[s[i]] = i

		// 更新最大长度
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}

	return maxLen
}
