// This package contains various string transformation functions
package string

// Reverse returns reverses the input string and returns it.
// s will not be altered.
func Reverse(s string) string {
	c := []rune(s)
	for i := 0; i < len(c)/2; i++ {
		j := len(c) - i - 1
		c[i], c[j] = c[j], c[i]
	}
	return string(c)

}
