package other

/**
JZ75 字符流中第一个不重复的字符
**/
var str string
var cnt map[byte]int = map[byte]int{}

func Insert(ch byte) {
	str += string(ch)
	cnt[ch]++
}

func FirstAppearingOnce() byte {
	for _, ch := range []byte(str) {
		if cnt[ch] == 1 {
			return ch
		}
	}
	return '#'
}
