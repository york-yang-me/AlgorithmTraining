package other

func Divide(nums []int, i int, n int, m int) int {
	if i > n {
		return (1 << 63) - 1
	}

	if len(nums[:i]) == 0 || len(nums[i:n]) == 0 {
		return (1 << 63) - 1
	}

	m = min(Divide(nums, i+1, n, m), sum(nums[:i])*sum(nums[i:n]))

	return m
}

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

func sum(array []int) int {
	res := 0
	for i := range array {
		res += array[i]
	}
	return res
}

// func main() {
// 	n := 0
// 	fmt.Scan(&n)
// 	scanner := bufio.NewScanner(os.Stdin)
// 	for scanner.Scan() {
// 		text := scanner.Text()
// 		str := strings.Split(text, " ")
// 		nums := make([]int, n)
// 		for i := range str {
// 			nums[i], _ = strconv.Atoi(str[i])
// 		}
// 		min := (1 << 63) - 1
// 		min = divide(nums, 1, n, min)
// 		fmt.Println(min)
// 	}
// }
