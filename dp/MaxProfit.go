package dp

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 动态规划
 * @param prices int整型一维数组
 * @return int整型
 */
func MaxProfit(prices []int) int {
	// write code here
	min := prices[0]
	max := 0
	for _, d := range prices {
		min = Min(min, d)
		max = Max(max, d-min)
	}
	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
