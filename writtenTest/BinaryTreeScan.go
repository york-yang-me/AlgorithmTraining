package writtentest

/**
 * 二叉树遍历
 * 对给定的二叉树依次完成前序，中序，后序遍历，并输出遍历结果。
 * 每行输入为一个二叉树，一维数组形式。其中-1表示Nil节点，例如：1,7,2,6,-1,4,8 构成的二叉树如图BinaryTreeScanImage
 * 对给定的二叉树依次完成前序，中序，后序遍历，并输出遍历结果
 * @param input int整型一维数组 -1表示Nil节点
 * @return int整型二维数组
 * 输入例子：
	[1,7,2,6,-1,4,8]
	输出例子：
	[[1,7,6,2,4,8],[6,7,1,4,2,8],[6,7,4,8,2,1]]
	例子说明：
	注意二维数组中的结果依次为：前序，中序，后序遍历的结果，Nil（-1）节点不用输出
 * 先通过层序构建二叉树。可以通过队列的方式逐层构建节点和它们的左右孩子。。
*/

/**
 * 对给定的二叉树依次完成前序，中序，后序遍历，并输出遍历结果
 * @param input int整型一维数组 -1表示Nil节点
 * @return int整型二维数组
 * 定义二叉树节点
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(arr []int) *TreeNode {
	if len(arr) == 0 || arr[0] == -1 {
		return nil
	}

	root := &TreeNode{Val: arr[0]}
	queue := []*TreeNode{root}
	i := 1

	var node *TreeNode
	// 按层次构造二叉树
	for i < len(arr) {
		if len(queue) != 0 {
			node = queue[0]
			queue = queue[1:]
		}
		// 构建左孩子
		if i < len(arr) && arr[i] != -1 {
			node.Left = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Left)
		}
		i++

		// 构建右孩子
		if i < len(arr) && arr[i] != -1 {
			node.Right = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func preOrderTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	preOrderTraversal(root.Left, res)
	preOrderTraversal(root.Right, res)
}

func inOrderTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	inOrderTraversal(root.Left, res)
	*res = append(*res, root.Val)
	inOrderTraversal(root.Right, res)
}

func postOrderTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	postOrderTraversal(root.Left, res)
	postOrderTraversal(root.Right, res)
	*res = append(*res, root.Val)
}

func BinaryTreeScan(arr []int) [][]int {
	// 构建二叉树
	root := buildTree(arr)

	// 保存遍历结果
	var perOrderRes, inOrderRes, postOrderRes []int

	// 执行前序 中序 后序遍历
	preOrderTraversal(root, &perOrderRes)
	inOrderTraversal(root, &inOrderRes)
	postOrderTraversal(root, &postOrderRes)

	result := [][]int{perOrderRes, inOrderRes, postOrderRes}

	return result
}

// func main() {
// 	// 输入例子
// 	arr := []int{1, -1, 7, -1, -1, -1, 4}

// 	fmt.Println(binaryTreeScan(arr))
// }
