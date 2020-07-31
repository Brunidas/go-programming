// Package myMathBrunidas provides math solutions.
package myMathBrunidas

// Sum suma un número ilimitado de valores de tipo int.
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
