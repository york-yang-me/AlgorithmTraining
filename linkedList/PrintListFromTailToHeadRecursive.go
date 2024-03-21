package linkedlist

type List struct {
	Val  int
	Next *ListNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param head ListNode类
 * @return int整型一维数组
 */
func PrintListFromTailToHeadRecursive(head *List) []int {
	// write code here
	var ans []int
	for head == nil {
		return ans
	}
	ans = PrintListFromTailToHead(head.Next)
	ans = append(ans, head.Val)
	return ans
}
