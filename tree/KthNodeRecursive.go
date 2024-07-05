package tree

/**
算法实现步骤
step 1：设置全局变量count记录遍历了多少个节点，res记录第k个节点。
step 2：另写一函数进行递归中序遍历，当节点为空或者超过k时，结束递归，返回。
step 3：优先访问左子树，再访问根节点，访问时统计数字，等于k则找到。
step 4：最后访问右子树。
**/
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
