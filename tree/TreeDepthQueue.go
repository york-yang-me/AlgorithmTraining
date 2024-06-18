package tree

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param pRoot TreeNode类
 * @return int整型
 */
func TreeDepthQueue(pRoot *TreeNode) int {
	// write code here
	if pRoot == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, pRoot)
	var depth int
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			size--
		}
		depth++
	}
	return depth
}
