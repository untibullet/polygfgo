package polygfgo

import "fmt"

type Polynomial struct {
	coefs []int
	len   int
	deg   int
}

func newPolynomialNoReverse(coefs []int) Polynomial {
	return Polynomial{coefs, len(coefs), len(coefs) - 1}.Normalize()
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
	return Polynomial{p.coefs[:i+1], i + 1, i}
}

func (p Polynomial) Add(q Polynomial) (product Polynomial) {
	if p.len >= q.len {
		product = p
	} else {
		product = Polynomial{
			append(p.coefs, make([]int, q.len-p.len)...),
			q.len,
			p.deg,
		}
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
		product = p
	} else {
		product = Polynomial{
			append(p.coefs, make([]int, q.len-p.len)...),
			q.len,
			p.deg,
		}
	}

	for i := 0; i < product.len; i++ {
		if i < q.len {
			product.coefs[i] -= q.coefs[i]
		}
	}

	product = product.Normalize()

	return
}

// Пока что возвращает неформатированный список коэффициентов
func (p Polynomial) ToString() string {
	return fmt.Sprint(p.coefs)
}
