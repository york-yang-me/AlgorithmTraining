package writtentest

/**

解数独
一个数独的解法需遵循如下规则：
数字 1 - 9 在每一行只能出现一次。
数字 1 - 9 在每一列只能出现一次。
数字 1 - 9 在每一个以粗实线分隔的 3 x 3 宫内只能出现一次。

**/

import (
	"fmt"
)

// 定义数独的大小
const size = 9

// 定义一个数独的结构体
type Sudoku struct {
	grid [size][size]int
}

// 打印数独的当前状态
func (s *Sudoku) Print() {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(s.grid[i][j], " ")
		}
		fmt.Println()
	}
}

// 检查在指定位置(row, col)放置数字val是否合法
func (s *Sudoku) isValid(row, col, val int) bool {
	// 检查行和列
	for i := 0; i < size; i++ {
		if s.grid[row][i] == val || s.grid[i][col] == val {
			return false
		}
	}
	// 检查所在的3x3宫
	startRow, startCol := row-row%3, col-col%3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if s.grid[i][j] == val {
				return false
			}
		}
	}
	return true
}

// 解数独的递归函数
func solveSudoku(s *Sudoku, row, col int) bool {
	if row == size {
		return true // 已经遍历完所有行，数独解决
	}
	if col == size {
		return solveSudoku(s, row+1, 0) // 当前行已遍历完，进入下一行
	}
	if s.grid[row][col] != 0 {
		return solveSudoku(s, row, col+1) // 当前位置已经有数字，进入下一个位置
	}
	// 尝试放置数字1-9
	for num := 1; num <= size; num++ {
		if s.isValid(row, col, num) {
			s.grid[row][col] = num // 放置数字
			if solveSudoku(s, row, col+1) {
				return true // 递归求解下一位置
			}
			s.grid[row][col] = 0 // 回溯，重置当前位置为0
		}
	}
	return false
}

// 解数独的入口函数
func Sudu(sudoku *Sudoku) {
	if !solveSudoku(sudoku, 0, 0) {
		fmt.Println("No solution exists.")
	}
}

// func main() {
// 	// 示例数独，0表示空格
// 	sudoku := Sudoku{
// 		grid: [size][size]int{
// 			{5, 3, 0, 0, 7, 0, 0, 0, 0},
// 			{6, 0, 0, 1, 9, 5, 0, 0, 0},
// 			{0, 9, 8, 0, 0, 0, 0, 6, 0},
// 			{8, 0, 0, 0, 6, 0, 0, 0, 3},
// 			{4, 0, 0, 8, 0, 3, 0, 0, 1},
// 			{7, 0, 0, 0, 2, 0, 0, 0, 6},
// 			{0, 6, 0, 0, 0, 0, 2, 8, 0},
// 			{0, 0, 0, 4, 1, 9, 0, 0, 5},
// 			{0, 0, 0, 0, 8, 0, 0, 7, 9},
// 		},
// 	}

// 	fmt.Println("Original Sudoku:")
// 	sudoku.Print()

// 	fmt.Println("\nSolving Sudoku:")
// 	solve(&sudoku)

// 	fmt.Println("\nSolved Sudoku:")
// 	sudoku.Print()
// }
