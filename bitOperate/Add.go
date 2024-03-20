package bitOperate

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *  位运算
 * @param num1 int整型
 * @param num2 int整型
 * @return int整型
 */
func Add(num1 int, num2 int) int {
	// write code here
	res := 0
	carry := 0
	for num2 != 0 {
		res = num1 ^ num2
		carry = (num1 & num2) << 1
		num1 = res
		num2 = carry
	}
	return num1
}
