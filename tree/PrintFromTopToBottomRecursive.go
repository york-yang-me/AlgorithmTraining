package tree

/**
 * 递归
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param root TreeNode类
 * @return int整型一维数组
 */
func PrintFromTopToBottomRecursive(root *TreeNode) []int {
	// write code here
	var res [][]int
	var result []int

	if root == nil {
		return result
	}

	traverse(root, &res, 1)

	for _, num := range res {
		result = append(result, num...)
	}
	return result
}

func traverse(root *TreeNode, res *[][]int, depth int) {
	if root != nil {
		// 新的一层
		if len(*res) < depth {
			*res = append(*res, []int{root.Val})
		} else {
			// 读取该层的一维数组，将元素加入末尾
			(*res)[depth-1] = append((*res)[depth-1], root.Val)
		}
		// 递归左右时深度+1
		traverse(root.Left, res, depth+1)
		traverse(root.Right, res, depth+1)
	}
}
