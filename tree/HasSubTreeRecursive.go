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
* 递归实现 深度优先遍历 DFS
  1. 判断有无子树
  2. 判断是否相等
* @param pRoot1 TreeNode类
* @param pRoot2 TreeNode类
* @return bool布尔型
*/
func HasSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	// write code here
	if pRoot1 == nil || pRoot2 == nil {
		return false
	}

	res := false
	if pRoot1.Val == pRoot2.Val {
		res = isSubTree(pRoot1, pRoot2)
	}
	if !res {
		res = HasSubtree(pRoot1.Left, pRoot2)
	}
	if !res {
		res = HasSubtree(pRoot1.Right, pRoot2)
	}

	return res

}

func isSubTree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	if pRoot2 == nil {
		return true
	}

	if pRoot1 == nil {
		return false
	}

	if pRoot1.Val != pRoot2.Val {
		return false
	}

	return isSubTree(pRoot1.Left, pRoot2.Left) && isSubTree(pRoot1.Right, pRoot2.Right)
}
