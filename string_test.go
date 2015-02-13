package string

import "testing"

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
