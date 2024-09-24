package tree

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param root TreeNode类
 * @param sum int整型
 * @return int整型
 */
var res = 0

func dfs(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	if sum == root.Val {
		res++
	}
	dfs(root.Left, sum-root.Val)
	dfs(root.Right, sum-root.Val)
}

func FindPath(root *TreeNode, sum int) int {
	// write code here
	if root == nil {
		return res
	}
	dfs(root, sum)
	FindPath(root.Left, sum)
	FindPath(root.Right, sum)
	return res
}
