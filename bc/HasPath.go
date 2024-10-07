package bc

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *  回溯 dfs
 * @param matrix char字符型二维数组
 * @param word string字符串
 * @return bool布尔型
 */
func HasPath(matrix [][]byte, word string) bool {
	// write code here
	words := []byte(word)
	for i := range matrix {
		for j := range matrix[0] {
			if dfs(matrix, words, i, j, 0) {
				return true
			}
		}
	}

	return false
}

func dfs(matrix [][]byte, word []byte, i, j, index int) bool {
	// 边界判断, 如果越界或者字符不匹配 返回false
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) || matrix[i][j] != word[index] {
		return false
	}

	// 如果已经查找到最后一个字符 说明路径成立
	if index == len(word)-1 {
		return true
	}

	// 保存当前坐标的值 用于后续复原
	tmp := matrix[i][j]
	// 标记当前坐标已经访问过
	matrix[i][j] = '.'

	// 沿着当前坐标的上下左右4个方向查找
	res := dfs(matrix, word, i+1, j, index+1) ||
		dfs(matrix, word, i-1, j, index+1) ||
		dfs(matrix, word, i, j+1, index+1) ||
		dfs(matrix, word, i, j-1, index+1)

	// 递归之后恢复当前坐标值
	matrix[i][j] = tmp
	return res
}
