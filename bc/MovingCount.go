package bc

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 路径问题 回溯
 * @param threshold int整型
 * @param rows int整型
 * @param cols int整型
 * @return int整型
 */
func MovingCount(threshold int, rows int, cols int) int {
	// write code here
	if threshold < 0 || rows <= 0 || cols <= 0 {
		return 0
	}

	isVisited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		isVisited[i] = make([]bool, cols)
	}

	return dfsCount(threshold, rows, cols, 0, 0, isVisited)
}

func dfsCount(threshold, rows, cols, x, y int, isVisited [][]bool) int {
	if x < 0 || y < 0 || x >= rows || y >= cols || isVisited[x][y] ||
		digitalSum(x)+digitalSum(y) > threshold {
		return 0
	}

	isVisited[x][y] = true
	return 1 + dfsCount(threshold, rows, cols, x-1, y, isVisited) +
		dfsCount(threshold, rows, cols, x+1, y, isVisited) +
		dfsCount(threshold, rows, cols, x, y-1, isVisited) +
		dfsCount(threshold, rows, cols, x, y+1, isVisited)
}

func digitalSum(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}
