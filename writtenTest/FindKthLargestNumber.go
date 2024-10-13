package writtentest

import (
	"regexp"
	"sort"
	"strconv"
)

/***
	混合字符串 按非负整数降序 找第K个数
***/

func FindKthLargestNumber(s string, k int) int {
	// 使用正则表达式提取字符串中的所有数字{
	re := regexp.MustCompile(`\d+`)
	numberString := re.FindAllString(s, -1)

	// 将提取出的字符串转换为整数切片
	nums := make([]int, 0, len(numberString))
	for _, numStr := range numberString {
		num, _ := strconv.Atoi(numStr)
		nums = append(nums, num)
	}

	// 按降序排序
	sort.Ints(nums)

	if k > len(nums) {
		return 0
	}

	return nums[k-1]
}
