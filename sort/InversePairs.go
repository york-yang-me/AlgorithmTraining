package sort

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 归并排序 对等划分子列
 * 两两归并 四四归并 一直下去直到归并了整个数组
 * @param nums int整型一维数组
 * @return int整型
 */
var counter int

func InversePairs(nums []int) int {
	// write code here
	cmod := 1000000007
	mergeSort(nums)
	return counter % cmod
}

func mergeSort(in []int) []int {
	d := len(in) / 2
	if d == 0 { // 输入数组元素只有一个
		return in
	}

	return merge(mergeSort(in[:d]), mergeSort(in[d:]))
}

func merge(a, b []int) []int {
	tmp := make([]int, 0, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			tmp = append(tmp, a[i])
			i++
		} else {
			// 出现逆序对 计算逆序对数量
			counter += len(a) - i // 个数等于a数组中剩余元素的数量
			tmp = append(tmp, b[j])
			j++
		}
	}

	// 将剩余元素直接追加到结果数组
	if i < len(a) {
		tmp = append(tmp, a[i:]...)
	}

	if j < len(b) {
		tmp = append(tmp, b[j:]...)
	}

	return tmp
}
