package writtentest

/**
定义一个排列是好排列，当且仅当所有相邻两个数乘积为偶，长度为n的好排列一共多少 。第一行输入整数n，需要返回好排列的个数。
数唯一不重复，比如输入2，好排列有 1 2 和 2 1，输入3，好排列有1 2 3 和3 2 1 用go语言实现
**/

// 检查排列是否是好排列
func isGoodPermutation(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if (arr[i]*arr[i+1])%2 != 0 {
			return false
		}
	}
	return true
}

// 生成所有排列并检查是否为好排列
func permute(nums []int, l int, r int, result *int) {
	if l == r {
		if isGoodPermutation(nums) {
			*result++
		}
		return
	}

	for i := l; i <= r; i++ {
		nums[l], nums[i] = nums[i], nums[l]
		permute(nums, l+1, r, result)
		nums[l], nums[i] = nums[i], nums[l] // 还原顺序
	}
}

func CountGoodPermutations(n int) int {
	if n == 1 {
		return 1 // 长度为1的排列只有1种
	}

	// 构造1到n的数字数组
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}

	// 初始化计数
	result := 0
	// 计算好排列的个数
	permute(nums, 0, n-1, &result)
	return result
}
