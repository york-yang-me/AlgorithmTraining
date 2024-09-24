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
 * 递归
 *
 * @param root TreeNode类
 * @param o1 int整型
 * @param o2 int整型
 * @return int整型
 */
func ClosestCommonAncestor(root *TreeNode, o1 int, o2 int) int {
	// write code here
	if root == nil {
		return -1
	}
	if root.Val == o1 || root.Val == o2 { // 该节点是其中一个o
		return root.Val
	}

	left := ClosestCommonAncestor(root.Left, o1, o2)   // 寻找左子树
	right := ClosestCommonAncestor(root.Right, o1, o2) // 寻找右子树
	if left == -1 {                                    // 左子树没有，在右子树中
		return right
	}
	if right == -1 { // 右子树没有，左子树中
		return left
	}
	return root.Val // 两边都非空，则当前节点是祖先

}
