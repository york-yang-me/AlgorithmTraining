package bitOperate

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 求1+2+3+...+n，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）
 * 闭包函数
 * @param n int整型
 * @return int整型
 */
func Sum_Solution(n int) int {
	// write code here
	var res int

	_ = (n > 0) && func() bool {
		res = n + Sum_Solution(n-1)
		return true
	}()

	return res
}
