package other

func Multiply(A []int) []int {
	var B = make([]int, len(A))
	for i := range A {
		tmp := A[i]
		B[i] = 1
		A[i] = 1
		for k := range A {
			B[i] *= A[k]
		}
		A[i] = tmp
	}
	return B
}
