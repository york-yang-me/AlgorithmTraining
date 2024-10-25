package other

func FirstNotRepeatingChar(str string) int {
	m := map[string]int{}
	for i := 0; i < len(str); i++ {
		tmp := string(str[i])
		m[tmp] += 1
	}
	for i := 0; i < len(str); i++ {
		tmp := string(str[i])
		if m[tmp] == 1 {
			return i
		}
	}
	return -1
}
