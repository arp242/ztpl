package tplfunc

import (
	"fmt"
	"html/template"
)

var FuncMapHTML = map[string]any{
	"unsafe":       Unsafe,
	"unsafe_js":    UnsafeJS,
	"checkbox":     Checkbox,
	"checked":      Checked,
	"option_value": OptionValue,
}

// Unsafe converts a string to template.HTML, preventing any escaping.
//
// Can be dangerous if used on untrusted input!
func Unsafe(s any) template.HTML {
	switch ss := s.(type) {
	default:
		panic(fmt.Sprintf("unsafe: unsupported type: %T", s))
	case string:
		return template.HTML(ss)
	case []byte:
		return template.HTML(ss)
	case template.HTML:
		return template.HTML(ss)
	case fmt.Stringer:
		return template.HTML(ss.String())
	case error:
		return template.HTML(ss.Error())
	}
}

// UnsafeJS converts a string to template.JS, preventing any escaping.
//
// Can be dangerous if used on untrusted input!
func UnsafeJS(s string) template.JS { return template.JS(s) }

// Checkbox adds a checkbox; if current is true then it's checked.
//
// It also adds a hidden input with the value "off" so that's sent to the server
// when the checkbox isn't sent, which greatly simplifies backend handling.
func Checkbox(current any, name string) template.HTML {
	var c bool
	switch cc := current.(type) {
	case bool:
		c = cc
	case interface{ Bool() bool }:
		c = cc.Bool()
	default:
		panic(fmt.Sprintf("checkbox: unknown type %T", cc))
	}

	if c {
		return template.HTML(fmt.Sprintf(`
			<input type="checkbox" name="%s" id="%[1]s" checked>
			<input type="hidden" name="%[1]s" value="off">
		`, template.HTMLEscapeString(name)))
	}
	return template.HTML(fmt.Sprintf(`
		<input type="checkbox" name="%s" id="%[1]s">
		<input type="hidden" name="%[1]s" value="off">
	`, template.HTMLEscapeString(name)))
}

// Checked returns a 'checked="checked"' attribute if id is in vals.
func Checked(vals []int64, id int64) template.HTMLAttr {
	for _, v := range vals {
		if id == v {
			return template.HTMLAttr(` checked="checked"`)
		}
	}
	return ""
}

// OptionValue inserts the value attribute, and selected attribute if the value
// is the same as current.
func OptionValue(current, value string) template.HTMLAttr {
	if value == current {
		return template.HTMLAttr(fmt.Sprintf(`value="%s" selected`,
			template.HTMLEscapeString(value)))
	}
	return template.HTMLAttr(fmt.Sprintf(`value="%s"`,
		template.HTMLEscapeString(value)))
}
