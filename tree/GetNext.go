package tree

type TreeLinkNode struct {
	Val   int
	Left  *TreeLinkNode
	Right *TreeLinkNode
	Next  *TreeLinkNode
}

func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
	if pNode == nil {
		return nil
	}
	if pNode.Right != nil {
		pNode = pNode.Right
		for pNode.Left != nil {
			pNode = pNode.Left
		}
		return pNode
	}
	for pNode.Next != nil {
		root := pNode.Next
		if root.Left == pNode {
			return root
		}
		pNode = pNode.Next
	}
	return nil
}
