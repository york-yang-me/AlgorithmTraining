package other

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * JZ57 和为S的两个数字
 * @param array int整型一维数组
 * @param sum int整型
 * @return int整型一维数组
 */
func FindNumbersWithSum(array []int, sum int) []int {
	// write code here
	cal := make(map[int]int)
	for _, v := range array {
		cal[v]++
	}

	for k, v := range cal {
		res := sum - k
		_, ok := cal[res]
		if res == k && v > 1 || ok {
			return []int{k, res}
		}
	}
	return nil
}
