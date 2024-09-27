package sa

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 二分法
 *
 * @param nums int整型一维数组
 * @return int整型
 */
func MinNumberInRotateArray(rotateArray []int) int {
	// write code here
	size := len(rotateArray)
	if size == 1 {
		return rotateArray[0]
	}
	return min(MinNumberInRotateArray(rotateArray[0:(size/2)]), MinNumberInRotateArray(rotateArray[(size/2):size]))
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
