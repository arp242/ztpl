package tplfunc

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// JSON prints any object as JSON.
func JSON(v any) string {
	j, err := json.Marshal(v)
	if err != nil {
		panic(fmt.Errorf("json: %w", err))
	}
	return string(j)
}

// JSONPretty prints any object as indented JSON.
func JSONPretty(v any) string {
	j, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		panic(fmt.Errorf("json_pretty: %w", err))
	}
	return string(j)
}

// Number formats a number with thousand separators using the separator sep.
//
// For floats it will always use '.' as the digit separator, unless sep is set
// to '.' in which case it will use ','.
func Number(n any, sep ...rune) string {
	if len(sep) == 0 {
		sep = []rune{','}
	}

	s := strconv.FormatFloat(toFloat(n), 'f', -1, 64)
	if len(s) < 4 {
		return s
	}
	s, d, _ := strings.Cut(s, ".")

	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	var out []rune
	for i := range b {
		if i > 0 && i%3 == 0 && sep[0] > 1 {
			out = append(out, sep[0])
		}
		out = append(out, rune(b[i]))
	}

	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}

	s = string(out)
	if d != "" {
		if sep[0] == '.' {
			s += "," + d
		} else {
			s += "." + d
		}
	}
	return s
}

// LargeNumber formats a number, adding the suffix "M" for values larger than a
// million or "k" for values larger than 100,000.
//
// This loses some accuracy.
func LargeNumber(n any, sep ...rune) string {
	if len(sep) == 0 {
		sep = []rune{','}
	}

	num := toFloat(n)
	switch {
	case num > 100_000_000:
		return Number(int64(num/1_000_000), sep...) + "M"
	case num > 1_000_000:
		s := fmt.Sprintf("%.1fM", num/1_000_000)
		if sep[0] == '.' {
			s = strings.ReplaceAll(s, ".", ",")
		}
		return s
	case num > 10_000:
		return Number(int64(num/1000), sep...) + "k"
	default:
		return Number(num, sep...)
	}
}

// Time formats a time as the given format string.
//
// Return empty string if the time is nil or the zero value.
func Time(a any, format string) string {
	var t time.Time
	switch tt := a.(type) {
	case time.Time:
		t = tt
	case *time.Time:
		if tt == nil {
			return ""
		}
		t = *tt
	default:
		panic(fmt.Sprintf("time: unsupported type %T", tt))
	}

	if t.IsZero() {
		return ""
	}
	switch format {
	case "":
		format = "2006-01-02"
	case "rfc3339":
		format = time.RFC3339
	case "rfc3339nano":
		format = time.RFC3339Nano
	case "ansic":
		format = time.ANSIC
	}
	return t.Format(format)
}

func Duration(a any, format string) string {
	var d time.Duration
	switch dd := a.(type) {
	case time.Duration:
		d = dd
	case *time.Duration:
		if dd == nil {
			return ""
		}
		d = *dd
	default:
		panic(fmt.Sprintf("time: unsupported type %T", dd))
	}

	// TODO
	return d.String()
}

func fact(unit byte) uint64 {
	switch unit {
	case 'b':
		return 1
	case 'k':
		return 1024
	case 'm':
		return 1024 * 1024
	case 'g':
		return 1024 * 1024 * 1024
	case 't':
		return 1024 * 1024 * 1024 * 1024
	case 'p':
		return 1024 * 1024 * 1024 * 1024 * 1024
	default:
		panic(fmt.Sprintf("size:unknown unit value: %q", unit))
	}
}

// Size formats a file size.
//
// The optional parameter max gives the highest unit to format it as. Values for
// this can be 'b', 'k', 'm', 'g, 't', 'p'.
//
// The format string controls some formatting aspects, as key/value pairs:
//
//	min=n
//	max=n
//	from=n
func Size(n any, format ...string) string {
	var (
		min  = byte('b')
		max  = byte('t')
		from = byte('b')
	)
	for _, opt := range format {
		k, v, ok := strings.Cut(opt, "=")
		if !ok || len(v) != 1 {
			panic(fmt.Sprintf("size: invalid option %q", opt))
		}
		switch k {
		case "min":
			min = v[0]
		case "max":
			max = v[0]
		case "from":
			from = v[0]
		}
	}

	bytes := toFloat(n) * float64(fact(from))
	_, _ = min, max

	units := []string{"", "K", "M", "G", "T", "P"}
	i := 0
	for ; i < len(units); i++ {
		if bytes < 1024 {
			return fmt.Sprintf("%.1f%s", bytes, units[i])
		}
		bytes /= 1024
	}
	return fmt.Sprintf("%.1f%s", bytes*1024, units[i-1])
}

func Slug(s string) string {
	var n strings.Builder
	n.Grow(len(s))
	didDash := false
	for _, c := range s {
		// All ASCII punctuation and control characters
		if c <= '/' || (c >= ':' && c <= '@') || (c >= '[' && c <= '`') || (c >= '{' && c <= 0x7f) {
			if !didDash {
				didDash = true
				n.WriteByte('-')
			}
		} else {
			didDash = false
			n.WriteRune(c)
		}
	}
	if didDash {
		return n.String()[:n.Len()-1]
	}
	return n.String()
}
