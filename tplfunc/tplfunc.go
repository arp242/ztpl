package tplfunc

import (
	"fmt"
	"reflect"
	"strings"
	"text/template"
)

// Add a new template function.
func Add(name string, f any) { FuncMap[name] = f }

// FuncMap contains all the template functions.
var FuncMap = template.FuncMap{
	// Math
	"int":    Int,
	"sum":    Sum,
	"sub":    Sub,
	"mult":   Mult,
	"div":    Div,
	"round":  Round,
	"abs":    Abs,
	"is_inf": IsInf,
	"min":    Min,
	"max":    Max,

	// Strings
	"substr":     Substr,
	"elide":      Elide,
	"has_prefix": HasPrefix,
	"has_suffix": HasSuffix,
	"join":       strings.Join,
	"ucfirst":    UCFirst,
	"cat":        Cat,

	// Misc
	"deref":    Deref,
	"if2":      If2,
	"map":      Map,
	"contains": Contains,

	// Formatting
	"json":         JSON,
	"json_pretty":  JSONPretty,
	"number":       Number,
	"large_number": LargeNumber,
	"time":         Time,
	"duration":     Duration,
	"size":         Size,
	"slug":         Slug,
}

// Deref dereferences a pointer.
func Deref(s any) any {
	v := reflect.ValueOf(s)
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return reflect.New(reflect.TypeOf(s).Elem()).Interface()
		}
		v = v.Elem()
	}
	return v.Interface()
}

// If2 returns yes if cond is true, and no otherwise.
func If2(cond bool, yes any, no ...any) any {
	if cond {
		return yes
	}
	switch len(no) {
	case 0:
		return ""
	case 1:
		return no[0]
	default:
		panic("if2: too many parameters")
	}
}

// Map creates a map
func Map(values ...any) map[string]any {
	if len(values)%2 != 0 {
		panic("map: need key/value")
	}
	dict := make(map[string]any, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			panic(fmt.Sprintf("map: key must be a string: %T: %#[1]v", key))
		}
		dict[key] = values[i+1]
	}
	return dict
}

// Contains reports if the slice contains the element value find.
func Contains(slice any, find any) bool {
	// Can't use type parameters with text/template :-/
	sliceV, findV := reflect.ValueOf(slice), reflect.ValueOf(find)

	if !sliceV.IsValid() {
		return false
	}
	if findV.Type() != sliceV.Type().Elem() {
		panic(fmt.Sprintf("mismatched types: %s and %s", sliceV.Type(), findV.Type()))
	}

	for i, l := 0, sliceV.Len(); i < l; i++ {
		if sliceV.Index(i).Equal(findV) {
			return true
		}
	}
	return false
}
