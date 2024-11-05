package other

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * JZ58 左旋转字符串
 *
 * @param str string字符串
 * @param n int整型
 * @return string字符串
 */
func LeftRotateString(str string, n int) string {
	// write code here
	if len(str) == 0 {
		return ""
	}

	start := n % len(str)
	return str[start:] + str[:start]
}
