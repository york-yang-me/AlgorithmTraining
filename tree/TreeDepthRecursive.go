package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param pRoot TreeNode类
 * @return int整型
 */
func TreeDepthRecursive(pRoot *TreeNode) int {
	// write code here
	if pRoot == nil {
		return 0
	}
	leftDepth := TreeDepthRecursive(pRoot.Left)
	rightDepth := TreeDepthRecursive(pRoot.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}
