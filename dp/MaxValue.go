package dp

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 动态规划
 * @param grid int整型二维数组
 * @return int整型
 */
func MaxValue(grid [][]int) int {
	// write code here
	dp := make([][]int, len(grid)+1)
	for i := range dp {
		dp[i] = make([]int, len(grid[0])+1)
	}
	dp[1][1] = grid[0][0]
	for i := range grid {
		for j := range grid[i] {
			x, y := i+1, j+1
			dp[x][y] = max(dp[x-1][y], dp[x][y-1]) + grid[i][j]
		}
	}
	return dp[len(grid)][len(grid[0])]
}
