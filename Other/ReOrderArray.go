package other

func ReOrderArray(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}

	odd := []int{}
	even := []int{}

	for i := range array {
		if array[i]%2 == 0 {
			even = append(even, array[i])
		} else {
			odd = append(odd, array[i])
		}
	}

	return append(odd, even...)

}
