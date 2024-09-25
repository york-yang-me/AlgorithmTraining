package queue

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 双端队列
 * 遍历次数超过k时 就把上一个区间里的最大值删掉
 * i从k-1开始 就可以把队里的最大值加入结果数组
 * @param num int整型一维数组
 * @param size int整型
 * @return int整型一维数组
 */
func MaxInWindows(num []int, k int) []int {
	// write code here 处理边界
	if k == 0 || len(num) == 0 || k > len(num) {
		return []int{}
	}

	queue := []int{}
	res := []int{}
	for i := 0; i < len(num); i++ {
		for len(queue) > 0 && queue[len(queue)-1] < num[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, num[i])
		if i >= k && queue[0] == num[i-k] {
			queue = queue[1:]
		}
		if i >= k-1 {
			res = append(res, queue[0])
		}
	}
	return res
}
