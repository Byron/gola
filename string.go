package string

func Reverse(s string) string {
	c := []byte(s)
	for i := 0; i < len(c) / 2; i++ {
		j := len(c) - i - 1
		c[i], c[j] = c[j], c[i]
	}
	return string(c)

}
