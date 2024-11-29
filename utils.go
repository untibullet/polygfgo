package polygfgo

import (
	"math"
)

const (
	ln2 = 0.693147180559945309417232121458176568075500134360255254120680009
)

func trimLeadingZeros(poly []int) []int {
	i := 0
	for i < len(poly) && poly[i] == 0 {
		i++
	}
	return poly[i:]
}

// Эта функция нужна чисто для операции деления, так как
// в ней затратно на каждом шаге вызывать Normalize() и получать
// пустой или нулевой многочлен для сравнения в Equals()
func isZero(poly []int) bool {
	for _, coeff := range poly {
		if coeff != 0 {
			return false
		}
	}
	return true
}

func expand(s []int, n int) []int {
	if n <= len(s) {
		return s
	}

	exp := make([]int, n)
	copy(exp, s)

	return exp
}

func reverse(s []int) []int {
	r := make([]int, len(s))
	copy(r, s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}

func nextPOT(n int) int {
	if isPOT(n) {
		return n
	}

	if n <= 0 {
		return 1
	}

	return int(math.Pow(2, math.Ceil(math.Log(float64(n))/ln2)))
}

func isPOT(n int) bool {
	return n&(n-1) == 0 && n >= 1
}

func toComplex128(s []int) []complex128 {
	c := make([]complex128, len(s))
	for i, v := range s {
		c[i] = complex(float64(v), 0)
	}

	return c
}

func toInt(s []complex128) []int {
	r := make([]int, len(s))
	for i, v := range s {
		r[i] = int(math.Round(real(v)))
	}

	return r
}

func modInverse(a, p int) int {
	a = a % p
	for x := 1; x < p; x++ {
		if (a*x)%p == 1 {
			return x
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
