package tplfunc

import (
	"errors"
	"fmt"
	"html/template"
	"testing"
)

func TestOptionValue(t *testing.T) {
	tests := []struct {
		current, value, want string
	}{
		{"a", "a", `value="a" selected`},
		{"", "a", `value="a"`},
		{"x", "a", `value="a"`},
		{"a&'a", "a&'a", `value="a&amp;&#39;a" selected`},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			out := OptionValue(tt.current, tt.value)
			want := template.HTMLAttr(tt.want)
			if out != want {
				t.Errorf("\nout:  %#v\nwant: %#v\n", out, want)
			}
		})
	}
}

type str int

func (str) String() string { return "<&>" }

func TestUnsafe(t *testing.T) {
	var b []byte

	tests := []struct {
		in   interface{}
		want template.HTML
	}{
		{"", ""},
		{"<x>", "<x>"},
		{"<>", "<>"},
		{[]byte("<>"), "<>"},
		{b, ""},
		{template.HTML("<>"), "<>"},
		{str(0), "<&>"},
		{errors.New("<>"), "<>"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := Unsafe(tt.in)
			if got != tt.want {
				t.Errorf("\ngot:  %q\nwant: %q", got, tt.want)
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("\ngot:  %#v\nwant: %#v", got, tt.want)
			// }
			// if d := ztest.Diff(got, tt.want); d != "" {
			// 	t.Errorf(d)
			// }
		})
	}
}
