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
 *
 * @param head ListNode类
 * @param val int整型
 * @return ListNode类
 */
func DeleteNode(head *ListNode, val int) *ListNode {
	// write code here
	res := &ListNode{}
	res.Next = head
	p := res
	for cur := head; nil != cur.Next; cur = cur.Next {
		if val == p.Next.Val {
			p.Next = cur.Next
			break
		}
		p = cur
	}
	return res.Next
}
