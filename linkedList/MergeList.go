package linkedlist

import (
	"sort"
)

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
创建一个新数组 然后原来的两个链表作比较 把p1 p2值取出来放进数组 然后对数组进行排序 最后将数组转成链表
注意引用实体，被引用的数据在内存中没有变化
*
* @param pHead1 ListNode类
* @param pHead2 ListNode类
* @return ListNode类
*/
func MergeList(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	var list []int
	for pHead1 != nil {
		list = append(list, pHead1.Val)
		pHead1 = pHead1.Next
	}
	for pHead2 != nil {
		list = append(list, pHead2.Val)
		pHead2 = pHead2.Next
	}
	sort.Ints(list)

	dummy := &ListNode{}
	curr := dummy
	for _, v := range list {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}
