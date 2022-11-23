package tplfunc

import (
	"fmt"
	"strings"

	"zgo.at/zstd/zstring"
)

// Cat/join any arguments.
func Cat(args ...any) string {
	b := new(strings.Builder)
	for _, a := range args {
		switch aa := a.(type) {
		case string:
			b.WriteString(aa)
		case []byte:
			b.WriteString(string(aa))

		case []string:
			b.WriteString(strings.Join(aa, ""))
		default:
			fmt.Fprintf(b, "%v", aa)
		}
	}
	return b.String()
}

// String converts anything to a string.
func String(v any) string { return fmt.Sprintf("%v", v) }

// HasPrefix tests whether the string s begins with prefix.
func HasPrefix(s, prefix string) bool { return strings.HasPrefix(s, prefix) }

// HasSuffix tests whether the string s ends with suffix.
func HasSuffix(s, suffix string) bool { return strings.HasSuffix(s, suffix) }

// UCFirst converts the first letter to uppercase, and the rest to lowercase.
func UCFirst(s string) string {
	f := ""
	for _, c := range s {
		f = string(c)
		break
	}
	return strings.ToUpper(f) + strings.ToLower(s[len(f):])
}

// Substr returns part of a string.
func Substr(s string, i, j int) string {
	if i == -1 {
		return s[:j]
	}
	if j == -1 {
		return s[i:]
	}
	return s[i:j]
}

// Elide a string to at most n characters.
func Elide(s string, n int) string {
	return zstring.ElideLeft(s, n)
}
