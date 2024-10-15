package sim

import "regexp"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 正则
 * @param str string字符串
 * @return bool布尔型
 */
func IsNumeric(str string) bool {
	// write code here
	reg := regexp.MustCompile(`^\s*[+-]?((\d+(\.\d*)?)|(\.\d+))([eE][+-]?\d+)?\s*$`)
	return reg.MatchString(str)
}
