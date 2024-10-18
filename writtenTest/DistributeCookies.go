package writtentest

import "fmt"

/**
分饼干
幼儿园老师想给她班上的孩子分饼干。所有的孩子都坐在一条线上，每个孩子都根据在课堂上的表现得到评分。
老师必须给每个孩子至少1个饼干。如果两个孩子坐在一起，那么评分较高的孩子必须得到更多的饼干(孩子必须左右都比较)。
输出老师购买的饼干总数的最小值。例如，假设她的学生的评分为[3,6,3,5,6,2]。
她给学生饼干的数量如下：[1,2,1,2,3,1]。她必须购买至少10个饼干。


时间限制：C/C++ 1秒，其他语言2秒
空间限制：C/C++ 32M，其他语言64M
输入描述：
假设学生评分为：[1,2,3],则输入应为：
3
1
2
3
第1行为数组元素大小，2至n+1行为数组元素（每行一个）
输出描述：
可分配最小值
示例1
输入例子：
6
3
6
3
5
6
2
输出例子：
10
例子说明：
孩子数量为6（第一个数），所以饼干数位  1+2+1+2+3+1=10
示例2
输入例子：
10
2
4
2
6
1
7
8
9
2
1
输出例子：
19
例子说明：
孩子数量为10,
所以饼干数为  1+2+1+2+1+2+3+4+2+1=19 (最后一个孩子获得一个饼干，但是他比前一个孩子评分小，所以前一个孩子饼干数加一)

**/

func DistributeCookies() { // func main()
	var n int
	fmt.Scan(&n)

	ratings := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ratings[i])
	}

	// 初始化饼干数组，所有孩子至少分到1个饼干
	cookies := make([]int, n)
	for i := range cookies {
		cookies[i] = 1
	}

	// 第一次遍历：从左到右，如果右边孩子的评分更高，则饼干数比左边孩子多
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			cookies[i] = cookies[i-1] + 1
		}
	}

	// 第二次遍历：从右到左，如果左边孩子的评分更高，则饼干数比右边孩子多
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			cookies[i] = max(cookies[i], cookies[i+1]+1)
		}
	}

	// 计算饼干总数
	totalCookies := 0
	for i := 0; i < n; i++ {
		totalCookies += cookies[i]
	}

	fmt.Println(totalCookies)
}
