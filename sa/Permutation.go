package sa

import "sort"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 递归+回溯
 * step 1：先对字符串按照字典序排序，获取第一个排列情况。
 * step 2：准备一个空串暂存递归过程中组装的排列情况。使用额外的vis数组用于记录哪些位置的字符被加入了。
 * step 3：每次递归从头遍历字符串，获取字符加入：首先根据vis数组，已经加入的元素不能再次加入了；同时，如果当前的元素str[i]与同一层的前一个元素str[i-1]相同且str[i-1]已经用，也不需要将其纳入。
 * step 4：进入下一层递归前将vis数组当前位置标记为使用过。
 * step 5：回溯的时候需要修改vis数组当前位置标记，同时去掉刚刚加入字符串的元素，
 * step 6：临时字符串长度到达原串长度就是一种排列情况。
 * @param str string字符串
 * @return string字符串一维数组
 */
func Permutation(str string) []string {
	// write code here
	t := []byte(str)
	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})
	n := len(t)
	ans := []string{}
	cur := make([]byte, 0, n)
	vis := make([]bool, n)
	var backtrack func(i int)
	backtrack = func(i int) {
		if i == n {
			ans = append(ans, string(cur))
			return
		}
		for j, b := range vis {
			if b || j > 0 && !vis[j-1] && t[j-1] == t[j] {
				continue
			}
			vis[j] = true
			cur = append(cur, t[j])
			backtrack(i + 1)
			cur = cur[:len(cur)-1]
			vis[j] = false
		}
	}
	backtrack(0)
	return ans
}
