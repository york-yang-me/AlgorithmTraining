package other

import (
	"sort"
	"strconv"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 重载排序
 * @param numbers int整型一维数组
 * @return string字符串
 */
func PrintMinNumber(numbers []int) string {
	// write code here
	sort.Slice(numbers, func(i, j int) bool {
		s1 := strconv.Itoa(numbers[i])
		s2 := strconv.Itoa(numbers[j])
		return s1+s2 < s2+s1
	})
	ans := ""
	for _, x := range numbers {
		ans += strconv.Itoa(x)
	}
	return ans
}
