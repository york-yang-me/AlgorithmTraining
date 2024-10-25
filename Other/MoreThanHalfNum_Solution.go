package other

func MoreThanHalfNum_Solution(nums []int) int {
	t := make(map[int]int)
	for _, v := range nums {
		t[v]++
	}

	for v := range t {
		if t[v] > (len(nums) / 2) {
			return v
		}
	}

	return -1
}
