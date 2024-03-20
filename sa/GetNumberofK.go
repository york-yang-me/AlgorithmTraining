package sa

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 二分法
 * @param nums int整型一维数组
 * @param k int整型
 * @return int整型
 */
func GetNumberOfK(nums []int, k int) int {
	// write code here
	if len(nums) == 0 {
		return 0
	}

	right := rightBound(nums, k)
	if right == -1 {
		return 0
	}

	left := leftBound(nums, k)

	return right - left + 1
}

func leftBound(arr []int, k int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		middle := left + (right-left)>>1

		if arr[middle] == k {
			right = middle - 1
		} else if arr[middle] < k {
			left = middle + 1
		} else if arr[middle] > k {
			right = middle - 1
		}
	}

	if left >= len(arr) || arr[left] != k {
		return -1
	}
	return left
}

func rightBound(arr []int, k int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		middle := left + (right-left)>>1

		if arr[middle] == k {
			left = middle + 1
		} else if arr[middle] > k {
			right = middle - 1
		} else if arr[middle] < k {
			left = middle + 1
		}
	}

	if right < 0 || arr[right] != k {
		return -1
	}
	return right
}
