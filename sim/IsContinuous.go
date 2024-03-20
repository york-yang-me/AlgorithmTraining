package sim

import "sort"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param numbers int整型一维数组
 * @return bool布尔型
 */
func IsContinuous(numbers []int) bool {
	// write code here
	sort.Ints(numbers)
	zeroNum := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == 0 {
			zeroNum++
		}
	}
	diffCnt := 0
	for i := zeroNum; i < len(numbers)-1; i++ {
		if numbers[i+1] == numbers[i] {
			return false
		}
		diffCnt += numbers[i+1] - numbers[i] - 1
	}

	return zeroNum >= diffCnt

}
