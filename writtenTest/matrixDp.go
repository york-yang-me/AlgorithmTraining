package writtentest

/**
矩阵动态规划
现有一个地图，由横线与竖线组成（参考围棋棋盘），且两点之间有行走距离起点为左上角，终点为右下角在地图上，
每次行走只能沿线移动到临近的点，并累加路径计算一个人从地图的起点走到终点的最小路径为多少。

输入描述：
m*n地图表示如下:

3
3
1 3 4
2 1 2
4 3 1

其中m=3，n=3 表示3*3的矩阵

行走路径为：下>右>右>下
输出描述：
路径总长：1+2+1+2+1=7
示例1
输入例子：
1
2
1 2
输出例子：
3
示例2
输入例子：
2
3
9 8 6
2 3 7
输出例子：
21

**/
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dfs(arr [][]int, rows, cols, i, j int, memo [][]int) int {
	if i >= rows || j >= cols {
		return int(^uint(0) >> 1)
	}

	if i == rows-1 && j == cols-1 {
		return arr[i][j]
	}

	if memo[i][j] != -1 {
		return memo[i][j]
	}

	resDown := dfs(arr, rows, cols, i+1, j, memo)
	resUp := dfs(arr, rows, cols, i, j+1, memo)

	memo[i][j] = arr[i][j] + min(resDown, resUp)

	return memo[i][j]
}

func MatrixDp() { // func main()
	m, n := 0, 0
	fmt.Scan(&m, &n)
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < m; i++ {
		scanner.Scan()
		text := scanner.Text()
		str := strings.Split(text, " ")
		for j := 0; j < n; j++ {
			matrix[i][j], _ = strconv.Atoi(str[j])
		}
	}

	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}

	result := dfs(matrix, m, n, 0, 0, memo)

	fmt.Println(result)
}
