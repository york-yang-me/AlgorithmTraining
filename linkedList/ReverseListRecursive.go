package linkedlist

// 递归求解
/**
步骤：
	判断终止条件: 链表为空，或者是链表没有尾结点的时候
	递归调用：从当前节点的下个节点开始递归
	逻辑处理：当前节点挂到递归之后的链表的末尾
注意：时间空间复杂度都为O(N)
**/

type ReverseListRecursiveNode struct {
	Val  int
	Next *ReverseListRecursiveNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 当前反转节点
 * 已反转的链表
 * 未反转的链表
 * @param pHead ListNode类
 * @return ListNode类
 */
func ReverseListRecursive(pHead *ReverseListRecursiveNode) *ReverseListRecursiveNode {
	// write code here
	if pHead == nil || pHead.Next == nil {
		return pHead
	}

	res := ReverseListRecursive(pHead.Next)
	pHead.Next.Next = pHead // 把当前节点的下一个节点的下一个节点，置为当前节点
	pHead.Next = nil        //设置尾节点为空 避免跳不出循环

	return res
}
