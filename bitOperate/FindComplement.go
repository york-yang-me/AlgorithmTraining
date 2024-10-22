package bitOperate

/**
求十进制整数反码 用go语言实现 输入 5 应返回 2， 输入7 应返回0，非负整数
**/

func FindComplement(num int) int {
	// 获取和num位数相同的全1数
	mask := ^0
	for num&mask != 0 {
		mask <<= 1
	}
	// 返回 num 的反码
	return ^num & ^mask
}
