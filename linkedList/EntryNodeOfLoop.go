package linkedlist

/**
	快慢指针 or hash法（另开辟一个空间记录链表 但不满足空间复杂度要求O(1)）
	快指针走了2x步，慢指针走了x,环长为y
	相遇则有 2x = x + ny n为整数
	→ x = ny
	如果某个节点从起点出发，走到快慢指针交汇点走的是x步(n*y步)。
	此时，如果某个指针从快慢指针交汇点开始如果走环长的整数倍，那么它到时候还会在原位置
**/
func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return nil
	}

	fast, slow := pHead, pHead
	for fast != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// 找到快慢指针相遇的节点后，从相遇节点到入口节点的距离 == 从头节点到入口节点的距离
	if fast == slow {
		tmp := pHead
		for tmp != slow {
			tmp = tmp.Next
			slow = slow.Next
		}
		return tmp
	}
	return nil
}
