package tree

import "container/list"

/***
把二叉树打印成多行

step 1：如果树为空，则返回空数组，没有任何打印结果。
step 2：使用队列辅助层次遍历，优先加入二叉树的根节点。
step 3：从根节点开始，每次进入一层的时候，记录队列的长度即本层的节点数，然后访问相应节点数全在同一个数组中，子节点加入队列中继续排队。
step 4：每次访问完一层将数组加入二维数组中再进入下一层。
***/
func Print(pRoot *TreeNode) [][]int {
	// write code here
	var result [][]int
	if pRoot == nil {
		return result
	}

	queue := list.New()
	queue.PushBack(pRoot)

	for queue.Len() > 0 {
		curlen := queue.Len()
		var curList []int
		for i := 0; i < curlen; i++ {
			curTree := queue.Remove(queue.Front()).(*TreeNode)
			curList = append(curList, curTree.Val)
			if curTree.Left != nil {
				queue.PushBack(curTree.Left)
			}

			if curTree.Right != nil {
				queue.PushBack(curTree.Right)
			}
		}
		result = append(result, curList)
	}
	return result
}
