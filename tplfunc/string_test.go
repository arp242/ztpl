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

func TestStrings(t *testing.T) {
	tests := []struct {
		tpl  string
		want string
		data any
	}{
		{`{{join .l ","}}`, "a,b", map[string]any{"l": []string{"a", "b"}}},

		{`{{cat "a" "b"}}`, "ab", nil},
		{`{{cat "a" 23}}`, "a23", nil},
		{`{{cat "x" .l}}`, "xab", map[string]any{"l": []string{"a", "b"}}},
		{`{{cat "x" .l}}`, "xmap[XX:YY]", map[string]any{"l": map[string]string{"XX": "YY"}}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			have := mktpl(t, "", tt.tpl, tt.data)
			if have != tt.want {
				t.Errorf("\nhave: %s\nwant: %s\n", have, tt.want)
			}
		})
	}
}
