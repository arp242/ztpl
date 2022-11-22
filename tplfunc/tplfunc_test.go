package tplfunc

import (
	"bytes"
	"testing"
	"text/template"
)

func mktpl(t *testing.T, text string, data any) string {
	t.Helper()

	tpl, err := template.New("").Option("missingkey=error").Funcs(FuncMap).Parse(text)
	if err != nil {
		t.Fatal(err)
	}

	have := new(bytes.Buffer)
	err = tpl.Execute(have, data)
	if err != nil {
		t.Fatal(err)
	}

	return have.String()
}

func TestMisc(t *testing.T) {
	s := "str"
	i := 42
	var n *int
	tests := []struct {
		tpl  string
		want string
		data any
	}{
		{`{{map "k" 5 "k2" "q"}}`, `map[k:5 k2:q]`, nil},
		{`{{map .k .v}}`, `map[key:val]`, map[string]any{"k": "key", "v": "val"}},

		{`{{deref "str"}}`, `str`, nil},
		{`{{deref .ptr}}`, `str`, map[string]any{"ptr": &s}},
		{`{{deref .ptr}}`, `42`, map[string]any{"ptr": &i}},
		{`{{deref .ptr}}`, `0`, map[string]any{"ptr": n}},

		{`{{if2 true "a" "b"}}`, `a`, nil},
		{`{{if2 false "a" "b"}}`, `b`, nil},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			have := mktpl(t, tt.tpl, tt.data)
			if have != tt.want {
				t.Errorf("\nhave: %s\nwant: %s\n", have, tt.want)
			}
		})
	}
}
