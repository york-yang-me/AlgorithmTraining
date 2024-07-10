package tree

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 分治法
 * 先找左子树再找右子树最后根节点
 * 再采用分治法检查每颗子树的左右节点是否合法
 * @param sequence int整型一维数组
 * @return bool布尔型
 */
var count = 0

func VerifySquenceOfBSTDivideConquer(sequence []int) bool {
	// write code here
	if count == 0 && len(sequence) == 0 {
		return false
	} else {
		count++
	}

	Len := len(sequence)
	if Len <= 1 {
		return true
	}

	key := sequence[Len-1]
	i := 0
	for i = 0; i < len(sequence); i++ {
		if sequence[i] > key {
			break
		}
	}

	if i == len(sequence) {
		return true
	}

	for j := i; j < Len; j++ {
		if sequence[j] < sequence[Len-1] {
			return false
		}
	}

	return VerifySquenceOfBSTDivideConquer(sequence[:i]) && VerifySquenceOfBSTDivideConquer(sequence[i:Len-1])
}
