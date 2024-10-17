package writtentest

/**
* 删除重复元素
* 给定一个数组，你需要删除其中重复出现的元素，只保留最后一次出现的重复元素，
  使得每个元素只出现一次，返回新数组，并保证新数组中的元素顺序与原数组一致。
* @param array int整型一维数组
* @return int整型一维数组
* 示例1
  输入例子：
  [3,5,8,2,3,8]
  输出例子：
  [5,2,3,8]
  示例2
  输入例子：
  [1,1,1,2,1]
  输出例子：
  [2,1]
*/
func RemoveDuplicate(array []int) []int {
	// 创建一个哈希表来记录每个元素最后一次出现的索引
	lastOccur := make(map[int]int)
	for i, v := range array {
		lastOccur[v] = i
	}

	// 创建结果数组，并且使用一个哈希表来记录哪些元素已添加
	result := []int{}
	checkMap := make(map[int]bool)

	for i, v := range array {
		// 遍历原数组，按顺序添加最后一次出现的元素
		if lastOccur[v] == i && !checkMap[v] {
			result = append(result, v)
			checkMap[v] = true
		}
	}

	return result
}
