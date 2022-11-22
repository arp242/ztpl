package tplfunc

import "testing"

func TestUCFirst(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"hello", "Hello"},
		{"helLo", "Hello"},
		{"€elLo", "€ello"},
		{"łelLo", "Łello"},
		{"语elLo", "语ello"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := UCFirst(tt.in)
			if got != tt.want {
				t.Errorf("\ngot:  %q\nwant: %q", got, tt.want)
			}
		})
	}
}
