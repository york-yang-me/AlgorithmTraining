package other

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 动态规划
 * 如果是在递归中，1或2或3或4时，因为前面已经有其他段了，因此直接*1或2或3或4就行了
 * @param n int整型
 * @return int整型
 */
func CutRope(n int) int {
	// write code here
	if n <= 3 {
		return n - 1
	}

	res := make([]int, n+1)
	res[1] = 1
	res[2] = 2
	res[3] = 3
	res[4] = 4

	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			res[i] = max(j*res[i-j], res[i])
		}
	}

	return res[n]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
