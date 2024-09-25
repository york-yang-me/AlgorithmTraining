package queue

import "strings"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param str string字符串
 * @return string字符串
 */
func ReverseSentence(str string) string {
	// write code here
	ss := strings.Split(str, " ")
	for i := 0; i < len(ss)/2; i++ {
		tmp := ss[len(ss)-1-i]
		ss[len(ss)-1-i] = ss[i]
		ss[i] = tmp
	}
	return strings.Join(ss, " ")
}
