package dp

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param number int整型
 * @return int整型
 */
func JumpFloorII(number int) int {
	// write code here
	ans := 1
	for i := 1; i < number; i++ {
		ans += JumpFloorII(number - i)
	}
	return ans
}
