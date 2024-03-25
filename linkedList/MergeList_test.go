package linkedlist

import (
	"reflect"
	"testing"
)

func TestMergeList(t *testing.T) {
	// 创建测试用例链表1: 1 -> 3 -> 5
	pHead1 := &ListNode{Val: 1}
	pHead1.Next = &ListNode{Val: 3}
	pHead1.Next.Next = &ListNode{Val: 5}

	// 创建测试用例链表2: 2 -> 4 -> 6
	pHead2 := &ListNode{Val: 2}
	pHead2.Next = &ListNode{Val: 4}
	pHead2.Next.Next = &ListNode{Val: 6}

	// 创建期望结果链表: 1 -> 2 -> 3 -> 4 -> 5 -> 6
	wantList := createLinkedList([]int{1, 2, 3, 4, 5, 6})

	// result := MergeList(pHead1, pHead2)
	result := MergeListNoSort(pHead1, pHead2)

	if !reflect.DeepEqual(wantList, result) {
		t.Errorf("test failed")
	}
}

// 创建链表的辅助函数
func createLinkedList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, num := range nums {
		curr.Next = &ListNode{Val: num}
		curr = curr.Next
	}
	return dummy.Next
}
