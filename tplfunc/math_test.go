package tplfunc

import (
	"strings"
	"testing"
)

func TestInt(t *testing.T) {
	tests := []struct {
		in      any
		want    int64
		wantErr string
	}{
		{"5", 5, ""},
		{[]byte("05"), 5, ""},
		{5, 5, ""},
		{5.2, 5, ""},
		{5.9, 5, ""},
		{[]int64{5}, 0, "unsupported type []int64"},
		{"a5", 0, `strconv.ParseInt: parsing "a5": invalid syntax`},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if tt.wantErr != "" {
				defer func() {
					r := recover()
					if r == nil {
						t.Fatal("no error")
						return
					}
					have := r.(error).Error()
					if !strings.Contains(have, tt.wantErr) {
						t.Errorf("wrong error\nhave: %q\nwant: %q", have, tt.want)
					}
				}()
			}

			have := Int(tt.in)
			if have != tt.want {
				t.Errorf("\nhave: %q\nwant: %q", have, tt.want)
			}
		})
	}
}

func TestArithmetic(t *testing.T) {
	i := 3
	var n *int
	tests := []struct {
		tpl  string
		want string
		data any
	}{
		{`{{sum 3 5}}`, `8`, nil},
		{`{{sum 3.5 5}}`, `8.5`, nil},
		{`{{sum .A .B}}`, `3`, map[string]any{"A": 1, "B": 2}},
		{`{{sum .A .B 10 1}}`, `14`, map[string]any{"A": int64(1), "B": 2}},
		{`{{sum .A .B}}`, `3.299999952316284`, map[string]any{"A": int64(1), "B": float32(2.3)}},

		{`{{sub 3 5}}`, `-2`, nil},
		{`{{sub 3.5 5}}`, `-1.5`, nil},
		{`{{sub .A .B}}`, `-1`, map[string]any{"A": 1, "B": 2}},
		{`{{sub .A .B 10 1}}`, `-12`, map[string]any{"A": int64(1), "B": 2}},
		{`{{sub .A .B}}`, `-1.2999999523162842`, map[string]any{"A": int64(1), "B": float32(2.3)}},

		{`{{mult 3 5}}`, `15`, nil},
		{`{{mult 3.5 5}}`, `17.5`, nil},
		{`{{mult .A .B}}`, `2`, map[string]any{"A": 1, "B": 2}},
		{`{{mult .A .B 10 1}}`, `20`, map[string]any{"A": int64(1), "B": 2}},
		{`{{mult .A .B}}`, `2.299999952316284`, map[string]any{"A": int64(1), "B": float32(2.3)}},

		{`{{div 3 5}}`, `0.6`, nil},
		{`{{div 3.5 5}}`, `0.7`, nil},
		{`{{div .A .B}}`, `0.5`, map[string]any{"A": 1, "B": 2}},
		{`{{div .A .B 10 1}}`, `0.05`, map[string]any{"A": int64(1), "B": 2}},
		{`{{div .A .B}}`, `0.4347826177095873`, map[string]any{"A": int64(1), "B": float32(2.3)}},

		{`{{sum .A .B .C}}`, `6`, map[string]any{"A": &i, "B": &i, "C": n}},
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
