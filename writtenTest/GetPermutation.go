package writtentest

import "strconv"

/**
力扣 leetcode 59
给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。按大小顺序列出所有排列情况，并一一标记，当 n = 3 时，所有排列如下：

1."123"

2."132"

3."213"

4."231"

5."312"

6."321"

给定 n 和 k，返回第 k 个排列。

注意：
给定 n 的范围是 [1, 9]。

给定 k 的范围是 [1, n!]。

示例:
输入:
n = 3, k = 3
输出:

"213"
**/
/**
1.计算出每个位置上数字的排列个数，并根据 k 的值确定应该选取哪个数字作为当前位置的值。

2.从左到右依次确定每个位置的数字，直到确定完所有位置上的数字。

3.将确定好的数字按顺序组合成字符串，即为第 k 个排列。
**/
func GetPermutation(n int, k int) string {
	nums := make([]int, n)
	factorial := make([]int, n+1)
	result := ""

	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	factorial[0] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i-1] * i
	}

	k--
	for i := n; i >= 1; i-- {
		index := k / factorial[i-1]
		result += strconv.Itoa(nums[index])
		nums = append(nums[:index], nums[index+1:]...)
		k %= factorial[i-1]
	}

	return result
}
