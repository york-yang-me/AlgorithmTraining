package writtentest

/**
 * 某企业在年终给优秀同学发放奖品时需要排队领取，由于男同学和女同学发放的奖品不同，为了提高发放效率，需要在发放奖品前先将队伍中的男同学排到一起，女同学排到一起。
 * 排队调换规则：只能挨个调换位置，不能跨越调换。
 * 我们用B字母代表男同学，用G字母代表女同学，输入一行只包含B和G的字符串（即男女同学排队领取奖品），使用你熟悉语言编程的方式计算出最少需要多少次调换才能将男同学和女同学分开排好吗？
 * @param peoples string字符串
 * @return int整型
 * 输入例子：
	"GGBBG"
	输出例子：
	2
	例子说明：
	调换过程是：第一次调换：GGBGB , 第二次调换：GGGBB
*/
func lineup(peoples string) int {
	// 计算酱所有男生（B)移动到左边的最少交换次数
	swapsB, countB := 0, 0
	for _, ch := range peoples {
		if ch == 'B' {
			// 每次遇到男生（B) , 计算器当前的偏移量
			swapsB += countB
		} else {
			// 每次遇到一个女生（G)， 表示在其左边应有一个男生，需要计数
			countB++
		}
	}

	// 计算将所有女生（G)移到左边的最少交换次数
	swapsG, countG := 0, 0
	for _, ch := range peoples {
		if ch == 'G' {
			swapsG += countG
		} else {
			countG++
		}
	}

	// 返回最小的交换次数
	if swapsB < swapsG {
		return swapsB
	}

	return swapsG

}
