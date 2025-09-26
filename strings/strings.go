package strings

func reverse(in string) string {
	return reverseRunes([]rune(in))
}

func reverseRunes(r []rune) string {
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}