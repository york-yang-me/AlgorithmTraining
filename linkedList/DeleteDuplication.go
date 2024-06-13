package linkedlist

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 注意删除头节点的问题
 * @param pHead ListNode类
 * @return ListNode类
 */
func DeleteDuplication(pHead *ListNode) *ListNode {
	// write code here
	if pHead == nil {
		return pHead
	}

	sentry := new(ListNode)
	sentry.Next = pHead
	prev, cur := sentry, pHead
	for cur != nil {
		// 如果有相同的 就移动指针至下一位
		if cur.Next != nil && cur.Val == cur.Next.Val {
			for cur.Next != nil && cur.Val == cur.Next.Val {
				cur = cur.Next
			}
			prev.Next = cur.Next
			cur = cur.Next
			continue
		}
		cur = cur.Next
		prev = prev.Next
	}
	return sentry.Next
}
