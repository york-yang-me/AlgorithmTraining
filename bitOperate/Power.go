package bitOperate

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param base double浮点型
 * @param exponent int整型
 * @return double浮点型
 */
func Power(base float64, exponent int) float64 {
	// write code here
	if exponent < 0 {
		base = 1 / base
		exponent = -exponent
	}
	res := 1.0
	for i := 0; i < exponent; i++ {
		res *= base
	}
	return res
}
