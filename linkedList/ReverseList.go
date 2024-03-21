package linkedlist

// 双链表求解
/**
步骤：
	保存 pHead 的下一个节点，防止后续链表丢失
	将 pHead 指向 newHead，完成节点反转
	更新 newHead 为 pHead，将完成反转的节点加入链表
	更新 pHead，继续反转还未反转的链表
**/

type ReverseListNode struct {
	Val  int
	Next *ReverseListNode
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
func ReverseList(pHead *ReverseListNode) *ReverseListNode {
	// write code here
	if pHead == nil || pHead.Next == nil {
		return pHead
	}
	var newHead *ReverseListNode
	// for pHead != nil {
	// 	pNext := pHead.Next
	// 	pHead.Next = newHead
	// 	newHead = pHead
	// 	pHead.Next = pNext
	// }
	//	平行赋值语法
	for pHead != nil {
		pHead, pHead.Next, newHead = pHead.Next, newHead, pHead
	}
	return newHead
}
