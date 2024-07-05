package tree

var index int

func KthNodeRecursive(proot *TreeNode, k int) int {
	// write code here
	res := midTravel(proot, k)
	if res == nil {
		return -1
	}
	return res.Val
}

func midTravel(node *TreeNode, targetPos int) *TreeNode {
	if node == nil {
		return nil
	}
	if res := midTravel(node.Left, targetPos); res != nil {
		return res
	}
	index += 1
	if targetPos == index {
		return node
	}
	if res := midTravel(node.Right, targetPos); res != nil {
		return res
	}
	return nil
}
