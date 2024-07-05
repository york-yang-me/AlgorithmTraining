package tree

// 前序确定根节点，再中序找到根节点，确定左右子树节点长度 进行分冶 还原节点
func ReconstructBinaryTreeRecursive(pre []int, vin []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	r, idx := pre[0], 0

	for _, v := range vin {
		if v == r {
			break
		}
		idx++
	}

	return &TreeNode{
		r,
		// 左子树
		ReconstructBinaryTreeRecursive(pre[1:idx+1], vin[:idx]),
		// 右子树
		ReconstructBinaryTreeRecursive(pre[idx+1:], vin[idx+1:]),
	}
}
