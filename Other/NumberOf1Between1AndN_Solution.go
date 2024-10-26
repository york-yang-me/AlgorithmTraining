package other

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 数学逻辑 找规律
 * @param n int整型
 * @return int整型
 */
func NumberOf1Between1AndN_Solution(n int) int {
	// write code here
	bitNum, count := 1, 0
	high, cur, low := n/10, n%10, 0
	for high != 0 || cur != 0 {
		if cur < 1 {
			// case 1: cur == 0
			// cur = 0时，高位需要减去一位用于低位进行计算
			count += high * bitNum
		} else if cur == 1 {
			// case 2: cur == 1
			// 相当于高位+低位计算结果 即（high * bitNum）+ (low + 1)
			count += high*bitNum + low + 1
		} else {
			// case 3: cur > 1
			// 相对于cur=0的情况 就不需要高位减去一位来计算低位的结果数了
			// 相当于（high * bitNum）+ (低位数结果数)
			count += (high + 1) * bitNum
		}
		// low, cur, high都向左偏移一个位
		// bitNum表示cur的数位
		low = low + cur*bitNum
		high, cur = high/10, high%10
		bitNum = bitNum * 10
	}
	return count
}
