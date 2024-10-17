package writtentest

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * 特殊去重
    给出一个无序的数组，删除数组中出现偶数次的元素，剩余的元素需要保留原来的顺序。
	时间限制：C/C++ 1秒，其他语言2秒
	空间限制：C/C++ 256M，其他语言512M
	示例1
	输入例子：
	[1, 2, 1, 2, 0, 2]
	输出例子：
	[2,2,0,2]
	例子说明：
	0、1、2出现的次数分别为1次、2次、3次，因此保留0、2，保留原来顺序后输出为 [2, 2, 0, 2]
 * @param arr int整型一维数组
 * @return int整型一维数组
*/
func RemoveDuplicate2(arr []int) []int {
	// write code here
	// 统计元素出现的次数
	freqMap := make(map[int]int)
	for _, v := range arr {
		freqMap[v]++
	}

	// 构建结果数组，保留出现奇数次的元素
	result := []int{}
	for _, v := range arr {
		if freqMap[v]%2 != 0 {
			result = append(result, v)
		}
	}

	return result
}
