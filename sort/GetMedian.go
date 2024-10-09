package sort

import "sort"

var array []int

func Insert(num int) {
	array = append(array, num)
}

func GetMedian() float64 {
	if len(array) < 2 {
		return float64(array[0])
	}
	var result float64
	sort.Ints(array)
	if len(array)%2 == 0 {
		result = float64(array[len(array)/2-1]+array[len(array)/2]) / 2
	} else {
		result = float64(array[len(array)/2])
	}
	return result
}
