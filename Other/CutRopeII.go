package other

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 原动态规划计算量太大 考虑使用快速幂（借助贪心思想）
 * 快速幂 相当于 二分 复杂度O(log2(n))
 * @param number long长整型
 * @return long长整型
 */

func CutRopeII(n int64) int64 {
	// write code here
	if n <= 3 {
		return n - 1
	}

	mod := int64(998244353)
	if n%3 == 0 {
		return fastPow(3, n/3, mod) // 余数为0 所有段长为3
	} else if n%3 == 1 { // 余数为1 最后一个段长为1 将长度为3的段+1 保留一个长度为4的段
		return (fastPow(3, (n/3)-1, mod) << 2) % mod
	} else { // 余数为2
		return (fastPow(3, n/3, mod) << 1) % mod
	}
}

// base^exp mod Mod
func fastPow(base, exp, mod int64) int64 {
	res := int64(1)
	for exp > 0 {
		if exp%2 == 1 {
			res = (res * base) % mod
		}
		base = (base * base) % mod
		exp /= 2
	}
	return res
}
