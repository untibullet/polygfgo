package polygfgo

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"runtime"
	"sync"
	"sync/atomic"
)

const UNIT_DEGREE = 1

type FieldInterface interface {
	GetPrime() int
	GetDegree() int
	GetIrreducible() Polynomial
	AddPolynomials(p1, p2 Polynomial) Polynomial
	SubPolynomials(p1, p2 Polynomial) Polynomial
	MulPolynomials(p1, p2 Polynomial) Polynomial
	DivPolynomials(p1, p2 Polynomial) (Polynomial, Polynomial, error)
	IsIrreducible(poly Polynomial) bool
	GCD(p1, p2 Polynomial) Polynomial
	ToString() string
}

func FieldFactory(p, m int, generator Polynomial, enableLogging bool) (field FieldInterface, err error) {
	if generator.deg > m {
		err = fmt.Errorf("the degree of the generator must be lower than or equal to %d", m)
		tryLog(enableLogging, err)
		return
	}
	if p < 2 || m < 1 {
		err = fmt.Errorf("invalid values of the numbers p=%d < 2 or m=%d < 1", p, m)
		tryLog(enableLogging, err)
		return
	}
	if m == 1 || generator.deg < 1 {
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

func (sf SimpleField) GetPrime() int {
	return sf.p
}

func (sf SimpleField) GetDegree() int {
	return 1 // Simple degree
}

func (sf SimpleField) GetIrreducible() Polynomial {
	return newZeroPolynomial() // Simple irreducible polinomial
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
		err := fmt.Errorf("division by zero is not supported")
		tryLog(f.enableLogging, err)
		return newZeroPolynomial(), newZeroPolynomial(), err
	}

	if p1.deg < p2.deg {
		return newZeroPolynomial(), f.Normalize(p1), nil
	}

	// Находим коэффициент для вычитания
	inv := modInverse(d[0], f.p)
	if inv == -1 {
		err := fmt.Errorf("there is no reverse element")
		tryLog(f.enableLogging, err)
		return newZeroPolynomial(), newZeroPolynomial(), err
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

// GenerateIrreduciblePolynomials генерирует все комбинации длины k из диапазона [0..n-1] с повторениями.
func GenerateIrreduciblePolynomials(simpleField SimpleField, length, workers, totalCount int) (<-chan Polynomial, error) {
	prime := simpleField.p
	if prime < 0 || length < 0 {
		return nil, errors.New("n и k должны быть неотрицательными")
	}
	if prime == 0 && length == 0 {
		return nil, errors.New("нельзя генерировать комбинации для n=0 и k=0")
	}

	// Используем big.Int для расчёта p^d, чтобы избежать переполнения
	total := new(big.Int).Exp(big.NewInt(int64(prime)), big.NewInt(int64(length)), nil)
	if total.Cmp(big.NewInt(0)) == 0 {
		// Нет комбинаций для n=0 и k>0
		out := make(chan Polynomial)
		close(out)
		return out, nil
	}

	// Проверяем, помещается ли total в uint64
	if !total.IsInt64() {
		return nil, errors.New("the value of p^m is too large for processing")
	}
	totalInt := total.Int64()

	// Разрешаем d=0 (пустая комбинация)
	if length == 0 {
		out := make(chan Polynomial, 1)
		out <- newZeroPolynomial()
		close(out)
		return out, nil
	}

	// Определяем количество воркеров
	if workers <= 0 {
		workers = runtime.NumCPU()
	}
	if workers > int(totalInt) {
		workers = int(totalInt) // Не создавать больше воркеров, чем комбинаций
	}

	// Распределение работы между горутинами
	chunkSize := totalInt / int64(workers)
	remainder := totalInt % int64(workers)

	counter := int32(totalCount)

	out := make(chan Polynomial, 100)
	var wg sync.WaitGroup
	wg.Add(workers)

	startIndex := int64(0)
	for w := 0; w < workers; w++ {
		size := chunkSize
		if int64(w) < remainder {
			size++
		}

		endIndex := startIndex + size

		go func(start, end int64) {
			defer wg.Done()
			for i := start; i < end; i++ {
				if totalCount != -1 && atomic.LoadInt32(&counter) <= 0 {
					break
				}
				comb, err := nthCombination(prime, length, i)
				if err != nil {
					tryLog(simpleField.enableLogging, err)
					continue
				}
				if comb[0] == 0 || comb[length-1] != 1 {
					continue
				}
				poly := Polynomial{comb, length, length - 1}
				if !(simpleField.IsIrreducible(poly)) {
					continue
				}
				if totalCount == -1 {
					out <- poly
					continue
				}
				if atomic.LoadInt32(&counter) <= 0 {
					break
				}
				// Атомарное уменьшение счетчика
				if atomic.AddInt32(&counter, -1) >= 0 {
					out <- poly
				} else {
					break
				}
			}
		}(startIndex, endIndex)
		startIndex = endIndex
	}

	// Закрываем канал после завершения всех горутин
	go func() {
		wg.Wait()
		close(out)
	}()

	return out, nil
}

// nthCombination вычисляет i-ю комбинацию для p^d.
func nthCombination(p, d int, i int64) ([]int, error) {
	if p <= 0 || d <= 0 {
		return nil, errors.New("p и d должны быть положительными")
	}

	comb := make([]int, d)
	current := i
	for j := d - 1; j >= 0; j-- {
		comb[j] = int(current % int64(p))
		current /= int64(p)
	}
	return comb, nil
}

func (f SimpleField) IsIrreducible(poly Polynomial) bool {
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

func (sf SimpleField) GCD(p1, p2 Polynomial) Polynomial {
	for !(p2.isZeroPolynomial()) {
		_, mod, _ := sf.DivPolynomials(p1, p2)
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

func (ef ExtendedField) GetPrime() int {
	return ef.p
}

func (ef ExtendedField) GetDegree() int {
	return ef.m
}

func (ef ExtendedField) GetIrreducible() Polynomial {
	return ef.generator
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
		err := fmt.Errorf("polinomial cannot be zero")
		tryLog(f.enableLogging, err)
		return newZeroPolynomial(), err
	}

	q := int(math.Pow(float64(f.p), float64(f.generator.deg)))

	return f.simple.PowModPolynomial(poly, q-2, f.generator), nil
}

func (ex ExtendedField) IsIrreducible(poly Polynomial) bool {
	return ex.simple.IsIrreducible(poly)
}

func (ef ExtendedField) GCD(p1, p2 Polynomial) Polynomial {
	return ef.simple.GCD(p1, p2)
}

func (f ExtendedField) ToString() string {
	return fmt.Sprintf("GF(%d^%d) mod %s", f.p, f.m, f.generator.ToString())
}
