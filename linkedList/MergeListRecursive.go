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
*
* 递归 困难的是如何判断递归终止的条件
如果一个链表为空 则返回特殊情况
如果pHead1的值比pHead2的小，那么需return pHead1，但是在return之前需要指定下次递归的节点是pHead1.Next
如果pHead2的值比pHead1的小，那么需return pHead2，但是在return之前需要指定下次递归的节点是pHead2.Next
*
* @param pHead1 ListNode类
* @param pHead2 ListNode类
* @return ListNode类
*/
func MergeListRecursive(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}

	if pHead1.Val < pHead2.Val {
		pHead1.Next = MergeListRecursive(pHead1.Next, pHead2)
		return pHead1
	} else {
		pHead2.Next = MergeListRecursive(pHead1, pHead2.Next)
		return pHead2
	}

}
