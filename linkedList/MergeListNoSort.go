package linkedlist

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/*
*
* 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
* 迭代
把原来的两个链表作比较 值小的取出来放进dummy链表中 然后剩余的再往后加
注意引用实体，被引用的数据在内存中没有变化
*
* @param pHead1 ListNode类
* @param pHead2 ListNode类
* @return ListNode类
*/
func MergeListNoSort(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	dummy := &ListNode{}
	curr := dummy

	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val <= pHead2.Val {
			// curr.Next = pHead1
			// pHead1 = pHead1.Next
			// curr = curr.Next
			curr.Next, pHead1 = pHead1, pHead1.Next
		} else {
			// curr.Next = pHead2
			// pHead2 = pHead2.Next
			// curr = curr.Next
			curr.Next, pHead2 = pHead2, pHead2.Next
		}
		curr = curr.Next
	}

	if pHead1 != nil {
		curr.Next = pHead1
	}

	if pHead2 != nil {
		curr.Next = pHead2
	}

	return dummy.Next
}
