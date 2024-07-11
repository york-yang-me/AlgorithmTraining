package tree

var nodeList []*TreeNode

func TreeConvertDoubleChains(pRootOfTree *TreeNode) *TreeNode {
	if pRootOfTree == nil {
		return nil
	}
	nodeList = []*TreeNode{}
	inOrder(pRootOfTree)

	head := nodeList[0]
	head.Left = nil
	preNode := head

	for i := 1; i < len(nodeList); i++ {
		preNode.Right = nodeList[i]
		nodeList[i].Left = preNode
		preNode = preNode.Right
	}

	return head
}

func inOrder(curNode *TreeNode) {
	if curNode == nil {
		return
	}

	inOrder(curNode.Left)
	nodeList = append(nodeList, curNode)
	inOrder(curNode.Right)
}
