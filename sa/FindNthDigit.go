package sa

import "strconv"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param n int整型
 * @return int整型
 */
func FindNthDigit(n int) int {
	// write code here
	// write code here
	//记录n是几位数
	digit := 1
	//记录当前位数区间的起始数字：1,10,100...
	start := 1
	//记录当前区间之前总共有多少位数字
	sum := 9
	//将n定位在某个位数的区间中
	for n > sum {
		n -= sum
		start *= 10
		digit++
		//该区间的总共位数
		sum = 9 * start * digit
	}
	//定位n在哪个数字上
	num := start + (n-1)/digit
	//定位n在数字的哪一位上
	index := (n - 1) % digit
	numstr := strconv.Itoa(num)
	ans := numstr[index] - '0'
	return int(ans)
}
