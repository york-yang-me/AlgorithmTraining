package other

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * JZ62 孩子们的游戏(圆圈中最后剩下的数)
 * @param n int整型
 * @param m int整型
 * @return int整型
 */
func LastRemaining_Solution(n int, m int) int {
	// write code here
	if n == 0 || m == 0 {
		return -1
	}
	last := 0
	for i := 2; i < n+1; i++ {
		last = (m + last) % i
	}
	return last
}
