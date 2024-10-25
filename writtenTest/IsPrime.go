package writtentest

func IsPrime(n int) []int {
	if n <= 2 {
		return []int{}
	}

	// 创建一个布尔切片 初始默认所有数字都是素数
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}

	for p := 2; p*p < n; p++ {
		// 如果 p 是素数
		if isPrime[p] {
			// 将所有 p 的倍数标记为非素数
			for i := p * p; i < n; i += p {
				isPrime[i] = false
			}
		}
	}

	// 收集所有素数
	var primes []int
	for i := 2; i < n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
