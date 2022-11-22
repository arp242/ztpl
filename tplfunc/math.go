package tplfunc

import (
	"fmt"
	"reflect"
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

func Round(n any, to int) any {
	// TODO
	return 0
}
