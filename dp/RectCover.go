package dp

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param number int整型
 * @return int整型
 */
func RectCover(number int) int {
	// write code here
	if number <= 2 {
		return number
	}

	return RectCover(number-1) + RectCover(number-2)
}
