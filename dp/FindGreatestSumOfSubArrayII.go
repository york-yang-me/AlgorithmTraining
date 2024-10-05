package dp

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param array int整型一维数组
 * @return int整型
 */
func FindGreatestSumOfSubArrayII(array []int) []int {
	// write code here
	max := array[0]
	left, right, sum := 0, 0, 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
		if sum >= max {
			max = sum
			right = i
		}
		if sum < 0 {
			sum = 0
			if i < len(array)-1 {
				left = i + 1
			}
		}
	}
	if left > right {
		return array[right : right+1]
	} else {
		return array[left : right+1]
	}
}
