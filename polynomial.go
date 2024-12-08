package polygfgo

import (
	"fmt"

	"github.com/mjibson/go-dsp/fft"
)

type Polynomial struct {
	coefs []int
	len   int
	deg   int
}

func (p Polynomial) GetDegree() int {
	return p.deg
}

func NewPolynomial(coefs []int) Polynomial {
	coefsCopy := make([]int, len(coefs))
	copy(coefsCopy, reverse(coefs))

	return Polynomial{coefsCopy, len(coefsCopy), len(coefsCopy) - 1}.Normalize()
}

func newPolynomialNoReverse(coefs []int) Polynomial {
	result := make([]int, len(coefs))
	copy(result, coefs)

	return Polynomial{result, len(result), len(result) - 1}.Normalize()
}

func newZeroPolynomial() Polynomial {
	return Polynomial{[]int{}, 0, -1}
}

func (p Polynomial) isZeroPolynomial() bool {
	return p.Equals(newZeroPolynomial())
}

// Сравнение не учитывает нормализацию полиномов (нулевые слагаемые учавствуют в сравнении)
func (p Polynomial) Equals(q Polynomial) (result bool) {
	if p.len != q.len || p.deg != q.deg {
		return
	}
	for i := 0; i < p.len; i++ {
		if p.coefs[i] != q.coefs[i] {
			return
		}
	}

	result = true
	return
}

func (p Polynomial) Normalize() Polynomial {
	i := p.len - 1
	for ; i >= 0; i-- {
		if p.coefs[i] != 0 {
			break
		}
	}

	result := make([]int, i+1)
	copy(result, p.coefs[:i+1])

	return Polynomial{result, i + 1, i}
}

func (p Polynomial) Add(q Polynomial) (product Polynomial) {
	if p.len >= q.len {
		product = Polynomial{make([]int, p.len), p.len, p.deg}
		copy(product.coefs, p.coefs)
	} else {
		product = Polynomial{make([]int, q.len), q.len, p.deg}
		copy(product.coefs, p.coefs)
	}

	for i := 0; i < product.len; i++ {
		if i < q.len {
			product.coefs[i] += q.coefs[i]
		}
	}

	product = product.Normalize()

	return
}

func (p Polynomial) Sub(q Polynomial) (product Polynomial) {
	if p.len >= q.len {
		product = Polynomial{make([]int, p.len), p.len, p.deg}
		copy(product.coefs, p.coefs)
	} else {
		product = Polynomial{make([]int, q.len), q.len, p.deg}
		copy(product.coefs, p.coefs)
	}

	for i := 0; i < product.len; i++ {
		if i < q.len {
			product.coefs[i] -= q.coefs[i]
		}
	}

	product = product.Normalize()

	return
}

func (p Polynomial) Mul(q Polynomial) Polynomial {
	if p.deg == -1 || q.deg == -1 {
		return newZeroPolynomial()
	}

	if p.deg == 0 {
		return q.MulScalar(p.coefs[0])
	}
	if q.deg == 0 {
		return p.MulScalar(q.coefs[0])
	}

	prodlen := p.deg + q.deg + 1
	potlen := nextPOT(prodlen)

	a := fft.FFT(toComplex128(expand(p.coefs, potlen)))
	b := fft.FFT(toComplex128(expand(q.coefs, potlen)))

	// Pointwise multiplication.
	c := make([]complex128, potlen)
	for i := 0; i < potlen; i++ {
		c[i] = a[i] * b[i]
	}

	return newPolynomialNoReverse(toInt(fft.IFFT(c))[:prodlen])
}

func (p Polynomial) MulScalar(alpha int) Polynomial {
	result := newPolynomialNoReverse(p.coefs)
	for i := 0; i < result.len; i++ {
		result.coefs[i] *= alpha
	}
	return result.Normalize()
}

func (p Polynomial) Sprint() string {
	return fmt.Sprint(p.coefs) + fmt.Sprintf("%d:%d", p.len, p.deg)
}

func (p Polynomial) ToString() string {
	return fmt.Sprint(reverse(p.coefs))
}
