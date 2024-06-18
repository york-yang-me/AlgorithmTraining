package tree

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param pRoot TreeNode类
 * @return TreeNode类
 */
func MirrorRecursive(pRoot *TreeNode) *TreeNode {
	// write code here
	if pRoot == nil {
		return pRoot
	}

	MirrorRecursive(pRoot.Left)
	MirrorRecursive(pRoot.Right)

	pRoot.Left, pRoot.Right = pRoot.Right, pRoot.Left

	return pRoot
}
