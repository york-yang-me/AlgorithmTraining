package bitOperate

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param n int整型
 * @return int整型
 */
func NumberOf1(n int) int {
	// write code here
	if n == 0 {
		return 0
	}
	count := 0

	for i := 0; i < 32; i++ {
		if (1 << i & n) > 0 {
			count++
		}
	}

	return count
}
