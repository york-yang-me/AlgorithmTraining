package sim

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param matrix int整型二维数组
 * @return int整型一维数组
 */
func PrintMatrix(matrix [][]int) []int {
	// write code here
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	res := []int{}
	return Print(res, matrix, 0, 0, len(matrix)-1, len(matrix[0])-1)
}

func Print(res []int, matrix [][]int, row1, col1, row2, col2 int) []int {
	curR, curC := row1, col1
	if (col1-col2) > 0 || (row1-row2) > 0 {
		return res
	}
	if row1 == row2 {
		for i := col1; i <= col2; i++ {
			res = append(res, matrix[row2][i])
		}
		return res
	} else if col1 == col2 {
		for i := row1; i <= row2; i++ {
			res = append(res, matrix[i][col1])
		}
		return res
	}
	for curC < col2 {
		res = append(res, matrix[curR][curC])
		curC++
	}
	for curR < row2 {
		res = append(res, matrix[curR][curC])
		curR++
	}
	for curC > col1 {
		res = append(res, matrix[curR][curC])
		curC--
	}
	for curR > row1 {
		res = append(res, matrix[curR][curC])
		curR--
	}
	return Print(res, matrix, row1+1, col1+1, row2-1, col2-1)
}
