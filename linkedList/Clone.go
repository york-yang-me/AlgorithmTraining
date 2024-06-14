package linkedlist

type RandomListNode struct {
	Label  int
	Next   *RandomListNode
	Random *RandomListNode
}

/**
 *
 * @param pHead RandomListNode类
 * @return RandomListNode类
 */
func Clone(head *RandomListNode) *RandomListNode {
	//write your code here
	if head == nil {
		return head
	}

	p1, p2 := head, head
	m := map[*RandomListNode]*RandomListNode{}

	// 遍历旧链表 建立节点map
	for p1 != nil {
		m[p1] = &RandomListNode{Label: p1.Label}
		p1 = p1.Next
	}

	// 遍历旧链表 对map进行节点链接
	for p2 != nil {
		m[p2].Next = m[p2.Next]
		m[p2].Random = m[p2.Random]
		p2 = p2.Next
	}

	// 指向链表头
	return m[head]
}
