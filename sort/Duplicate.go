package sort

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param numbers int整型一维数组
 * @return int整型
 */
func Duplicate(numbers []int) int {
	// write code here
	var m = map[int]int{}
	for i := 0; i < len(numbers); i++ {
		_, exist := m[numbers[i]]
		if exist {
			return numbers[i]
		} else {
			m[numbers[i]] = numbers[i]
		}
	}

	return -1
}
