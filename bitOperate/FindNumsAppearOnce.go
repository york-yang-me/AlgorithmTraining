package bitOperate

import (
	"sort"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 哈希 位运算
 * @param nums int整型一维数组
 * @return int整型一维数组
 */
func FindNumsAppearOnce(nums []int) []int {
	// write code here
	sort.Ints(nums)
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if len(res) != 0 && res[len(res)-1] == nums[i] {
			res = res[:len(res)-1]
			continue
		}
		res = append(res, nums[i])
	}
	return res
}
