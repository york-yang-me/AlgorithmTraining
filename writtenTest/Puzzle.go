package writtentest

import (
	"strings"
)

/**
* 拼图问题/八数问题
* BFS-广度优先搜索
*
**/
// 构造邻接
var neighbors = [9][]int{{2, 4}, {1, 3, 5}, {2, 6}, {1, 5, 7}, {2, 4, 6, 8}, {3, 5, 0}, {4, 8}, {5, 7, 0}, {6, 8}}

func SlidingPuzzle(board [][]int) int {
	const target = "123456780"

	// 二维数组转换成start字符串
	s := make([]byte, 0, 9)
	for _, r := range board {
		for _, v := range r {
			s = append(s, '0'+byte(v))
		}
	}
	start := string(s)

	// 边界值 目标状态直接返回0
	if start == target {
		return 0
	}

	// 枚举 status 通过一次交换操作得到的状态
	// status 是当前拼图状态
	// x 是空格 0 的位置
	// 通过查找 neighbors 数组，x 的相邻位置上的数字与空格交换，生成新状态，并将其加入返回值 ret
	// 每次交换后，还原状态，以便生成所有可能的状态
	get := func(status string) (ret []string) {
		s := []byte(status)
		x := strings.Index(status, "0")
		for _, y := range neighbors[x] {
			s[x], s[y] = s[y], s[x]
			ret = append(ret, string(s))
			s[x], s[y] = s[y], s[x]
		}
		return
	}

	type pair struct {
		status string
		step   int
	}
	// 使用q进行BFS
	q := []pair{{start, 0}}
	// 哈希表 用来记录访问的状态 避免重复搜索
	seen := map[string]bool{start: true}
	// 每次从队列中取出一个状态 p，通过 get 函数获取所有可能的下一个状态
	// 如果状态还未访问过，且等于目标状态 target，返回所需步数
	// 如果没有找到目标状态，则将该新状态加入队列继续搜索
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, nxt := range get(p.status) {
			if !seen[nxt] {
				if nxt == target {
					return p.step + 1
				}
				seen[nxt] = true
				q = append(q, pair{nxt, p.step + 1})
			}
		}
	}
	// 遍历所有状态都无法找到目标状态
	return -1
}
