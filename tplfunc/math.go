package tplfunc

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

// Base the arithmetic on float64; they're safe for natural numbers up to 2^53,
// which should be enough, and using floats means we can use any numeric type.
func toFloat(n any) float64 {
	v := reflect.ValueOf(n)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return 0
		}
		n = v.Elem().Interface()
	}

	switch nn := n.(type) {
	case float32:
		return float64(nn)
	case float64:
		return float64(nn)
	case int:
		return float64(nn)
	case int8:
		return float64(nn)
	case int16:
		return float64(nn)
	case int32:
		return float64(nn)
	case int64:
		return float64(nn)
	case uint:
		return float64(nn)
	case uint8:
		return float64(nn)
	case uint16:
		return float64(nn)
	case uint32:
		return float64(nn)
	case uint64:
		return float64(nn)
	default:
		panic(fmt.Sprintf("tplfunc: unsupported type: %T", nn))
	}
}

// Int converts any int, float, or string to an integer.
//
// Floats are always rounded down; strings and []byte is parsed as a base-10
// number. Any other type will panic.
func Int(n any) int64 {
	switch nn := (n).(type) {
	case float32:
		return int64(nn)
	case float64:
		return int64(nn)
	case int:
		return int64(nn)
	case int8:
		return int64(nn)
	case int16:
		return int64(nn)
	case int32:
		return int64(nn)
	case int64:
		return int64(nn)
	case uint:
		return int64(nn)
	case uint8:
		return int64(nn)
	case uint16:
		return int64(nn)
	case uint32:
		return int64(nn)
	case uint64:
		return int64(nn)
	case []byte:
		return Int(string(nn))
	case string:
		num, err := strconv.ParseInt(nn, 10, 64)
		if err != nil {
			panic(fmt.Errorf("ztpl.Int: converting string: %w", err))
		}
		return num
	default:
		panic(fmt.Errorf("ztpl.Int: unsupported type %T", n))
	}
}

// Sum all the given numbers.
func Sum(n, n2 any, n3 ...any) any {
	r := toFloat(n) + toFloat(n2)
	for _, nn := range n3 {
		r += toFloat(nn)
	}
	return r
}

// Sub subtracts all the given numbers.
func Sub(n, n2 any, n3 ...any) any {
	r := toFloat(n) - toFloat(n2)
	for _, nn := range n3 {
		r -= toFloat(nn)
	}
	return r
}

// Mult multiplies all the given numbers.
func Mult(n, n2 any, n3 ...any) any {
	r := toFloat(n) * toFloat(n2)
	for _, nn := range n3 {
		r *= toFloat(nn)
	}
	return r
}

// Div all the given numbers.
func Div(n, n2 any, n3 ...any) any {
	r := toFloat(n) / toFloat(n2)
	for _, nn := range n3 {
		r /= toFloat(nn)
	}
	return r
}

// Abs gets the absolute value of n.
func Abs(n any) float64 {
	return math.Abs(toFloat(n))
}

// IsInf reports if n is Inf.
func IsInf(n float64) bool {
	return math.IsInf(n, 0)
}

// Round n; if to is 0 it will round to nearest, <1 to floor, >1 to ceil.
func Round(n any, to int) float64 {
	r := toFloat(n)
	switch {
	case to > 0:
		return math.Ceil(r)
	case to < 0:
		return math.Floor(r)
	default:
		return math.Round(r)
	}
	return 0
}

func Min(a, b any) float64 {
	aa, bb := toFloat(a), toFloat(b)
	return math.Min(aa, bb)
}

func Max(a, b any) float64 {
	aa, bb := toFloat(a), toFloat(b)
	return math.Max(aa, bb)
}
