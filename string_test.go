package string

import "testing"

var result string

func BenchmarkReverse(b *testing.B) {
	s := "Hello, WǾrld™!"
	var r string
	for i := 0; i < b.N; i++ {
		r = Reverse(s)
	}
	result = r
}

func TestReverse(t *testing.T) {
	pairs := []struct {
		s, want string
	}{
		{"Backward", "drawkcaB"},
		{"WǾrld", "dlrǾW"},
	}

	for _, c := range pairs {
		got := Reverse(c.s)
		if got != c.want {
			t.Errorf("Reverse(%g) == %g", c.s, got)
		}
	}
}
