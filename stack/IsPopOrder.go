package stack

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param pushV int整型一维数组
 * @param popV int整型一维数组
 * @return bool布尔型
 */
func IsPopOrder(pushV []int, popV []int) bool {
	// write code here
	index, pushS := 0, []int{}
	for i := 0; i < len(pushV); i++ {
		pushS = append(pushS, pushV[i])

		if popV[index] == pushS[len(pushS)-1] {
			for len(pushS) != 0 && popV[index] == pushS[len(pushS)-1] {
				index++
				if len(pushS) == 1 {
					pushS = []int{}
					continue
				}
				pushS = pushS[:len(pushS)-1]
			}
		}
	}
	return index >= len(popV)
}
