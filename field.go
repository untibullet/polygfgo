package polygfgo

import (
	"fmt"
	"math"
)

const UNIT_DEGREE = 1

type FieldInterface interface {
	AddPolynomials(p1, p2 Polynomial) Polynomial
	SubPolynomials(p1, p2 Polynomial) Polynomial
	MulPolynomials(p1, p2 Polynomial) Polynomial
	DivPolynomials(p1, p2 Polynomial) (Polynomial, Polynomial, error)
	RandomIrreducible(deg int) Polynomial
	ToString() string
}

func FieldFactory(p, m int, generator Polynomial, enableLogging bool) (field FieldInterface, err error) {
	if generator.deg > m {
		err = fmt.Errorf("the degree of the generator must be lower than or equal to %d", m)
		return
	}
	if m == 1 {
		field = SimpleField{p, enableLogging}
		return
	}
	field = ExtendedField{SimpleField{p, enableLogging}, p, m, generator, enableLogging}
	return
}

// Представление конечного поля GF(p)
type SimpleField struct {
	p             int
	enableLogging bool
}

func (f SimpleField) Normalize(poly Polynomial) (product Polynomial) {
	product = newPolynomialNoReverse(poly.coefs)
	for i := 0; i < product.len; i++ {
		product.coefs[i] %= f.p
		if product.coefs[i] < 0 {
			product.coefs[i] += f.p
		}
	}

	product = product.Normalize()

	return
}

func (f SimpleField) AddPolynomials(p1, p2 Polynomial) (product Polynomial) {
	product = f.Normalize(p1.Add(p2))
	return
}

func (f SimpleField) SubPolynomials(p1, p2 Polynomial) (product Polynomial) {
	product = f.Normalize(p1.Sub(p2))
	return
}

func (f SimpleField) MulPolynomials(p1, p2 Polynomial) (product Polynomial) {
	product = f.Normalize(p1.Mul(p2))
	return
}

func (f SimpleField) DivPolynomials(p1, p2 Polynomial) (quot, rem Polynomial, err error) {
	q := []int{}
	r := make([]int, p1.len)
	copy(r, reverse(p1.coefs))
	d := make([]int, p2.len)
	copy(d, reverse(p2.coefs))

	if p2.isZeroPolynomial() {
		return newZeroPolynomial(), newZeroPolynomial(), fmt.Errorf("division by zero is not supported")
	}

	if p1.deg < p2.deg {
		return newZeroPolynomial(), f.Normalize(p1), nil
	}

	// Находим коэффициент для вычитания
	inv := modInverse(d[0], f.p)
	if inv == -1 {
		return newZeroPolynomial(), newZeroPolynomial(), fmt.Errorf("there is no reverse element")
	}

	for len(r) >= len(d) && !isZero(r) {
		leadCoeff := (r[0] * inv) % f.p
		if leadCoeff < 0 {
			leadCoeff += f.p
		}
		q = append(q, leadCoeff)

		// Вычитаем (leadCoeff * b) из r
		for i := 0; i < len(d); i++ {
			r[i] = (r[i] - leadCoeff*d[i]) % f.p
			if r[i] < 0 {
				r[i] += f.p
			}
		}
		// Удаляем старший член
		r = r[1:]
	}

	quot, rem = NewPolynomial(q), NewPolynomial(r)

	return
}

// PowModPolynomial выполняет быстрое возведение в степень многочлена base в степени exp по модулю mod.
// Параметры:
// - base: Многочлен, который нужно возводить в степень.
// - exp: Степень возведения.
// - mod: Многочлен, по модулю которого вычисляется результат.
// Возвращает:
// - Результат вычисления base^exp % mod.
func (f SimpleField) PowModPolynomial(base Polynomial, exp int, mod Polynomial) Polynomial {
	// Результат инициализируем как единичный многочлен [1]
	result := newPolynomialNoReverse([]int{1})

	// Копируем base для работы, чтобы не изменять исходный многочлен
	_, currentBase, _ := f.DivPolynomials(base, mod)

	for exp > 0 {
		if exp%2 == 1 {
			// Если текущий бит степени exp равен 1, умножаем результат на currentBase
			result = f.MulPolynomials(result, currentBase)
			_, result, _ = f.DivPolynomials(result, mod) // Берем остаток от деления
		}

		// Возводим currentBase в квадрат
		currentBase = f.MulPolynomials(currentBase, currentBase)
		_, currentBase, _ = f.DivPolynomials(currentBase, mod) // Берем остаток от деления

		// Переходим к следующему биту
		exp /= 2
	}

	return f.Normalize(result)
}

func (f SimpleField) RandomIrreducible(deg int) (irreducible Polynomial) {
	return
}

func (f SimpleField) gorutineRandomIrreducible(start, stop int, deg int) (mass []Polynomial) {
	return
}

func (f SimpleField) isIrreducible(poly Polynomial) bool {
	if poly.isZeroPolynomial() {
		return false
	}
	n := poly.deg

	if n == 0 || (poly.coefs[0] == 0 && n > 1) {
		return false
	}
	if n == 1 {
		return true
	}

	x := newPolynomialNoReverse([]int{0, 1}) // x
	for m, i := n/2, 1; i <= m; i++ {
		tmp := f.PowModPolynomial(x, int(math.Pow(float64(f.p), float64(i))), poly)
		tmp = f.SubPolynomials(tmp, x)
		if f.GCD(poly, tmp).deg > 0 {
			return false
		}
	}
	return true
}

func (f SimpleField) GCD(p1, p2 Polynomial) Polynomial {
	for !(p2.isZeroPolynomial()) {
		_, mod, _ := f.DivPolynomials(p1, p2)
		p1, p2 = p2, mod
	}
	return p1
}

func (f SimpleField) ToString() string {
	return fmt.Sprintf("GF(%d)", f.p)
}

// Представление конечного поля GF(q), q = p^m
// Вохможно стоит хранить в атрибутах простое поле
type ExtendedField struct {
	simple        SimpleField
	p, m          int
	generator     Polynomial
	enableLogging bool
}

// Возращает poly(x) mod g(x)
func (f ExtendedField) Normalize(poly Polynomial) (product Polynomial) {
	_, product, _ = SimpleField{f.p, f.enableLogging}.DivPolynomials(poly, f.generator)
	return
}

func (f ExtendedField) AddPolynomials(p1, p2 Polynomial) (product Polynomial) {
	product = f.Normalize(p1.Add(p2))
	return
}

func (f ExtendedField) SubPolynomials(p1, p2 Polynomial) (product Polynomial) {
	product = f.Normalize(p1.Sub(p2))
	return
}

func (f ExtendedField) MulPolynomials(p1, p2 Polynomial) (product Polynomial) {
	product = f.Normalize(p1.Mul(p2))
	return
}

func (f ExtendedField) DivPolynomials(p1, p2 Polynomial) (Polynomial, Polynomial, error) {
	inverse, err := f.modInverse(p2)
	if err != nil {
		return newZeroPolynomial(), newZeroPolynomial(), err
	}

	return newZeroPolynomial(), f.MulPolynomials(p1, inverse), nil
}

func (f ExtendedField) modInverse(poly Polynomial) (Polynomial, error) {
	if poly.deg >= f.generator.deg {
		_, poly, _ = f.simple.DivPolynomials(poly, f.generator)
	}
	if poly.isZeroPolynomial() {
		return newZeroPolynomial(), fmt.Errorf("polinomial cannot be zero")
	}

	q := int(math.Pow(float64(f.p), float64(f.generator.deg)))

	return f.simple.PowModPolynomial(poly, q-2, f.generator), nil
}

func (f ExtendedField) RandomIrreducible(deg int) (irreducible Polynomial) {
	return
}

func (f ExtendedField) ToString() string {
	return fmt.Sprintf("GF(%d^%d) mod %s", f.p, f.m, f.generator.ToString())
}
