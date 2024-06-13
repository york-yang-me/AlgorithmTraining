package linkedlist

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 类似于截取绳子，先从头切取K长度，之后在相同的绳子后面找到k长度的绳子
 *  双指针 倒数k个 正数n-k个
 * @param pHead ListNode类
 * @param k int整型
 * @return ListNode类
 */
func FindKthToTail(pHead *ListNode, k int) *ListNode {
	if pHead == nil {
		return nil
	}
	fast := pHead
	for i := 0; i < k; i++ {
		if fast == nil {
			return fast
		}
		if fast != nil {
			fast = fast.Next
		}
	}
	last := pHead
	for fast != nil {
		fast = fast.Next
		last = last.Next
	}
	return last
}
