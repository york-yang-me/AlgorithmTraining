package linkedlist

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 *解题思路：
 *A+B = B+A
 *判断两个链表是否为空 如果是空 返回{}
 *如果一个链表等于另一个链表则返回任意一个
 *如果两个链表不相等，则遍历检索是否存在相同元素，输出第一个相同元素
 *如果两个链表没有相同元素，则返回空
 *
 * @param pHead1 ListNode类
 * @param pHead2 ListNode类
 * @return ListNode类
 */
func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	if pHead1 == nil || pHead2 == nil {
		return nil
	}

	cur1, cur2 := pHead1, pHead2
	for cur1 != cur2 {
		if cur1 == nil {
			cur1 = pHead2
		} else {
			cur1 = cur1.Next
		}

		if cur2 == nil {
			cur2 = pHead1
		} else {
			cur2 = cur2.Next
		}
	}
	return cur1
}
