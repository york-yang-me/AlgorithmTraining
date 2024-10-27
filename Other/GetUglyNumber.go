package other

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param index int整型
 * @return int整型
 */
func GetUglyNumber_Solution(index int) int {
	// 1 2 3 4 5 6
	if index <= 6 {
		return index
	}

	i2, i3, i5 := 0, 0, 0
	res := make([]int, index)
	res[0] = 1 // 设置第一个丑数为 1

	for i := 1; i < index; i++ {
		// 获取下一个丑数 三者中最小的
		res[i] = min(res[i2]*2, min(res[i3]*3, res[i5]*5))
		// 第一次是 2 3 5比较 得到最小是2
		// 第二次是 4 3 5比较 得到最小是3
		if res[i] == res[i2]*2 {
			i2++
		}
		if res[i] == res[i3]*3 {
			i3++
		}
		if res[i] == res[i5]*5 {
			i5++
		}
	}
	return res[index-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
