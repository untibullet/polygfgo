package polygfgo

import (
	"math"
	"testing"
)

func TestSimpleField_Normalize(t *testing.T) {
	t.Run("take modulo with positive and negative coefficients", func(t *testing.T) {
		poly := Polynomial{[]int{-123, 1, 0, 234271, 32, 5}, 6, 5}
		f := SimpleField{13, false}

		got := f.Normalize(poly)
		want := Polynomial{[]int{7, 1, 0, 11, 6, 5}, 6, 5}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("take modulo when all coefficients are multiples of the field size", func(t *testing.T) {
		poly := Polynomial{[]int{26, 39, 52, 65}, 4, 3}
		f := SimpleField{13, false}

		got := f.Normalize(poly)
		want := Polynomial{[]int{}, 0, -1}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("take modulo for zero polynomial", func(t *testing.T) {
		poly := Polynomial{[]int{0, 0, 0}, 3, -1}
		f := SimpleField{7, false}

		got := f.Normalize(poly)
		want := Polynomial{[]int{}, 0, -1}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("take modulo for large positive coefficients", func(t *testing.T) {
		poly := Polynomial{[]int{12345, 54321, 99999}, 3, 2}
		f := SimpleField{17, false}

		got := f.Normalize(poly)
		want := Polynomial{[]int{3, 6, 5}, 3, 2}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("take modulo with all coefficients already in range", func(t *testing.T) {
		poly := Polynomial{[]int{3, 7, 10, 6}, 4, 3}
		f := SimpleField{11, false}

		got := f.Normalize(poly)
		want := Polynomial{[]int{3, 7, 10, 6}, 4, 3}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})
}

func TestSimpleField_Add(t *testing.T) {
	t.Run("add polynomials", func(t *testing.T) {
		f := SimpleField{5, false}
		p1 := Polynomial{[]int{1, 2, 3}, 3, 2}
		p2 := Polynomial{[]int{4, 3, 1}, 3, 2}

		got := f.AddPolynomials(p1, p2)
		want := Polynomial{[]int{0, 0, 4}, 3, 2}

		if !got.Equals(want) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})
}

func TestSimpleField_Sub(t *testing.T) {
	t.Run("subtract polynomials", func(t *testing.T) {
		f := SimpleField{7, false}
		p1 := Polynomial{[]int{6, 5, 4}, 3, 2}
		p2 := Polynomial{[]int{3, 2, 1}, 3, 2}

		got := f.SubPolynomials(p1, p2)
		want := Polynomial{[]int{3, 3, 3}, 3, 2}

		if !got.Equals(want) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})
}

func TestSimpleField_Mul(t *testing.T) {
	t.Run("multiplication of polynomials #1", func(t *testing.T) {
		f := SimpleField{11, false}
		p1 := Polynomial{[]int{2, 3, 0, 299}, 4, 3}
		p2 := Polynomial{[]int{-4, 5}, 2, 1}

		got := f.MulPolynomials(p1, p2)
		want := Polynomial{[]int{3, 9, 4, 3, 10}, 5, 4}.Normalize()

		if !got.Equals(want) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("multiplication of polynomials #2", func(t *testing.T) {
		f := SimpleField{7, false}
		p1 := NewPolynomial([]int{2, 3, 4, 3})
		p2 := NewPolynomial([]int{5, 0, 0})

		got := f.MulPolynomials(p1, p2)
		want := NewPolynomial([]int{3, 1, 6, 1, 0, 0})

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("multiplication of polynomials #3", func(t *testing.T) {
		f := SimpleField{101, false}
		p1 := NewPolynomial([]int{77, 38, 39, 25})
		p2 := NewPolynomial([]int{70, 14, 96, 54, 55, 2, 87})

		got := f.MulPolynomials(p1, p2)
		want := NewPolynomial([]int{37, 1, 49, 2, 79, 84, 69, 12, 9, 54})

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("multiplication of polynomials #4", func(t *testing.T) {
		f := SimpleField{104729, false}
		p1 := NewPolynomial([]int{43068, 29273, 102881, 104460, 76030, 74011, 81127, 31023, 28077})
		p2 := NewPolynomial([]int{98010, 46658, 66335, 83646, 11212, 81169, 69139})

		got := f.MulPolynomials(p1, p2)
		want := NewPolynomial([]int{97064, 27196, 16095, 36853, 67884, 92566, 38286, 56450, 31237, 85799, 677, 50232, 54361, 23521, 63688})

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("multiplication of polynomials #5", func(t *testing.T) {
		f := SimpleField{104729, false}
		p1 := NewPolynomial([]int{98010, 0, 0, 0, 0, 0, 69139})
		p2 := NewPolynomial([]int{98010, 46658, 66335, 83646, 11212, 81169, 69139})

		got := f.MulPolynomials(p1, p2)
		want := NewPolynomial([]int{6762, 63524, 21759, 63069, 71452, 54121, 65806, 24804, 43197, 65414, 87139, 40026, 55574})

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})
}

func TestSimpleField_Div(t *testing.T) {
	t.Run("division polynomials #1", func(t *testing.T) {
		f := SimpleField{7, false}
		poly1 := newPolynomialNoReverse([]int{6, 0, 1, 3})
		poly2 := newPolynomialNoReverse([]int{0, 1, 5})

		gotQuot, gotRem, _ := f.DivPolynomials(poly1, poly2)
		wantQuot := newPolynomialNoReverse([]int{4, 2})
		wantRem := newPolynomialNoReverse([]int{6, 3})

		if !(gotQuot.Equals(wantQuot) && gotRem.Equals(wantRem)) {
			t.Errorf(
				"Expected %s, %s but got %s, %s in GF(%d)",
				wantQuot.Sprint(), wantRem.Sprint(),
				gotQuot.Sprint(), gotRem.Sprint(),
				f.p,
			)
		}
	})

	t.Run("division polynomials #2", func(t *testing.T) {
		f := SimpleField{101, false}
		poly1 := newPolynomialNoReverse([]int{78, 90, 94, 30, 11, 47, 93, 42, 7})
		poly2 := newPolynomialNoReverse([]int{87, 2, 55, 54, 96, 14, 70})

		gotQuot, gotRem, _ := f.DivPolynomials(poly1, poly2)
		wantQuot := newPolynomialNoReverse([]int{5, 43, 91})
		wantRem := newPolynomialNoReverse([]int{47, 76, 98, 41, 82, 25})

		if !(gotQuot.Equals(wantQuot) && gotRem.Equals(wantRem)) {
			t.Errorf(
				"Expected %s, %s but got %s, %s in GF(%d)",
				wantQuot.Sprint(), wantRem.Sprint(),
				gotQuot.Sprint(), gotRem.Sprint(),
				f.p,
			)
		}
	})

	t.Run("division polynomials #3", func(t *testing.T) {
		f := SimpleField{104729, false}
		poly1 := newPolynomialNoReverse([]int{28077, 31023, 81127, 74011, 76030, 104460, 102881, 29273, 43068})
		poly2 := newPolynomialNoReverse([]int{69139, 81169, 11212, 83646, 66335, 46658, 98010})

		gotQuot, gotRem, _ := f.DivPolynomials(poly1, poly2)
		wantQuot := newPolynomialNoReverse([]int{83370, 55844, 8130})
		wantRem := newPolynomialNoReverse([]int{89078, 81555, 102427, 37407, 2958, 97383})

		if !(gotQuot.Equals(wantQuot) && gotRem.Equals(wantRem)) {
			t.Errorf(
				"Expected %s, %s but got %s, %s in GF(%d)",
				wantQuot.Sprint(), wantRem.Sprint(),
				gotQuot.Sprint(), gotRem.Sprint(),
				f.p,
			)
		}
	})

	t.Run("division with zero remainder", func(t *testing.T) {
		f := SimpleField{5, false}
		poly1 := newPolynomialNoReverse([]int{1, 2, 1})
		poly2 := newPolynomialNoReverse([]int{1, 1})

		gotQuot, gotRem, _ := f.DivPolynomials(poly1, poly2)
		wantQuot := newPolynomialNoReverse([]int{1, 1})
		wantRem := newPolynomialNoReverse([]int{})

		if !(gotQuot.Equals(wantQuot) && gotRem.Equals(wantRem)) {
			t.Errorf(
				"Expected %s, %s but got %s, %s in GF(%d)",
				wantQuot.Sprint(), wantRem.Sprint(),
				gotQuot.Sprint(), gotRem.Sprint(),
				f.p,
			)
		}
	})

	t.Run("division with larger divisor", func(t *testing.T) {
		f := SimpleField{3, false}
		poly1 := newPolynomialNoReverse([]int{2, 1})
		poly2 := newPolynomialNoReverse([]int{2, 1, 1})

		gotQuot, gotRem, _ := f.DivPolynomials(poly1, poly2)
		wantQuot := newPolynomialNoReverse([]int{})
		wantRem := newPolynomialNoReverse([]int{2, 1})

		if !(gotQuot.Equals(wantQuot) && gotRem.Equals(wantRem)) {
			t.Errorf(
				"Expected %s, %s but got %s, %s in GF(%d)",
				wantQuot.Sprint(), wantRem.Sprint(),
				gotQuot.Sprint(), gotRem.Sprint(),
				f.p,
			)
		}
	})

	t.Run("division of zero polynomial", func(t *testing.T) {
		f := SimpleField{11, false}
		poly1 := newPolynomialNoReverse([]int{0})
		poly2 := newPolynomialNoReverse([]int{1, 1, 43, 10, 2, 5912441})

		gotQuot, gotRem, _ := f.DivPolynomials(poly1, poly2)
		wantQuot := newPolynomialNoReverse([]int{})
		wantRem := newPolynomialNoReverse([]int{})

		if !(gotQuot.Equals(wantQuot) && gotRem.Equals(wantRem)) {
			t.Errorf(
				"Expected %s, %s but got %s, %s in GF(%d)",
				wantQuot.Sprint(), wantRem.Sprint(),
				gotQuot.Sprint(), gotRem.Sprint(),
				f.p,
			)
		}
	})

	t.Run("division by zero polynomial", func(t *testing.T) {
		f := SimpleField{7, false}
		poly1 := newPolynomialNoReverse([]int{3, 6, 2})
		poly2 := newPolynomialNoReverse([]int{})

		_, _, err := f.DivPolynomials(poly1, poly2)

		if err == nil {
			t.Errorf("Expected panic when dividing by zero in field GF(%d)", f.p)
		}
	})
}

func TestSimpleField_PowMod(t *testing.T) {
	t.Run("exponentiation of polynomials #1", func(t *testing.T) {
		f := SimpleField{37, false}
		p1 := Polynomial{[]int{23, 28, 26, 30, 22, 7, 9, 25, 1}, 9, 8}
		p2 := Polynomial{[]int{2, 4, 10, 6, 18}, 2, 1}

		got := f.PowModPolynomial(p2, int(math.Pow(37, 8))-2, p1)
		want := newPolynomialNoReverse([]int{26, 12, 16, 19, 33, 16, 7, 14})

		if !got.Equals(want) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("exponentiation of polynomials #2", func(t *testing.T) {
		f := SimpleField{5, false}
		p1 := Polynomial{[]int{2, 4, 4, 2, 4, 2, 1}, 7, 6}
		p2 := Polynomial{[]int{1, 1, 4}, 3, 2}

		got := f.PowModPolynomial(p2, int(math.Pow(float64(f.p), float64(p1.deg)))-2, p1)
		want := newPolynomialNoReverse([]int{4, 3, 4, 4, 2, 2})

		if !got.Equals(want) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("exponentiation of polynomials #3", func(t *testing.T) {
		f := SimpleField{7, false}
		p1 := Polynomial{[]int{1, 4, 3, 4, 1}, 5, 4}
		p2 := Polynomial{[]int{1, 4, 2, 6}, 4, 3}

		got := f.PowModPolynomial(p2, int(math.Pow(float64(f.p), float64(p1.deg)))-2, p1)
		want := newPolynomialNoReverse([]int{1, 3, 5, 2})

		if !got.Equals(want) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})
}

func TestGenerateIrreducible(t *testing.T) {
	t.Run("16 thread calculation in the field GF(3) and degree 4", func(t *testing.T) {
		f := SimpleField{3, false}
		degree := 4

		ch, _ := GenerateIrreduciblePolynomials(f, degree+1, 16, -1)
		got := 0
		for range ch {
			got++
		}
		want := 18

		if got != want {
			t.Errorf("Expected %d gut got %d", want, got)
		}
	})

	t.Run("2 thread calculation in the field GF(3) and degree 4", func(t *testing.T) {
		f := SimpleField{3, false}
		degree := 4

		ch, _ := GenerateIrreduciblePolynomials(f, degree+1, 2, -1)
		got := 0
		for range ch {
			got++
		}
		want := 18

		if got != want {
			t.Errorf("Expected %d gut got %d", want, got)
		}
	})

	t.Run("16 thread calculation in the field GF(5) and degree 4", func(t *testing.T) {
		f := SimpleField{5, false}
		degree := 4

		ch, _ := GenerateIrreduciblePolynomials(f, degree+1, 16, -1)
		got := 0
		for range ch {
			got++
		}
		want := 150

		if got != want {
			t.Errorf("Expected %d gut got %d", want, got)
		}
	})

	t.Run("1 thread calculation in the field GF(5) and degree 4", func(t *testing.T) {
		f := SimpleField{5, false}
		degree := 4

		ch, _ := GenerateIrreduciblePolynomials(f, degree+1, 1, -1)
		got := 0
		for range ch {
			got++
		}
		want := 150

		if got != want {
			t.Errorf("Expected %d gut got %d", want, got)
		}
	})

	t.Run("16 thread calculation in the field GF(5), degree 4 and limit = 14", func(t *testing.T) {
		f := SimpleField{5, false}
		degree := 4

		ch, _ := GenerateIrreduciblePolynomials(f, degree+1, 0, 14)
		got := 0
		for range ch {
			got++
		}
		want := 14

		if got != want {
			t.Errorf("Expected %d gut got %d", want, got)
		}
	})
}

func TestIsIrreducible(t *testing.T) {
	t.Run("irreducible test number #1", func(t *testing.T) {
		f := SimpleField{37, false}
		poly := Polynomial{[]int{27, 29, 18, 29, 17, 23, 25, 24, 14, 1}, 10, 9}

		got := f.IsIrreducible(poly)
		want := true

		if got != want {
			t.Errorf("Expected %t but got %t", want, got)
		}
	})

	t.Run("irreducible test number #2", func(t *testing.T) {
		f := SimpleField{11, false}
		poly := Polynomial{[]int{6, 6, 4, 0, 1, 5, 1}, 7, 6}

		got := f.IsIrreducible(poly)
		want := false

		if got != want {
			t.Errorf("Expected %t but got %t", want, got)
		}
	})

	t.Run("irreducible test number #3", func(t *testing.T) {
		f := SimpleField{199933, false}
		poly := Polynomial{[]int{41194, 79985, 163946, 161238, 52940, 80299, 96191, 15330, 133939, 194819, 160338, 189015, 176142, 188277, 99410, 123846, 188414, 64313, 68982, 116765, 28267, 173093, 106559, 1}, 24, 23}

		got := f.IsIrreducible(poly)
		want := true

		if got != want {
			t.Errorf("Expected %t but got %t", want, got)
		}
	})

	t.Run("irreducible test number #1", func(t *testing.T) {
		f := SimpleField{3, false}
		poly := Polynomial{[]int{2, 2, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, 101, 100}

		got := f.IsIrreducible(poly)
		want := true

		if got != want {
			t.Errorf("Expected %t but got %t", want, got)
		}
	})
}

func TestGCD(t *testing.T) {
	t.Run("calculating GCD of two reducible polinomials", func(t *testing.T) {
		f := SimpleField{11, false}
		p1 := Polynomial{[]int{7, 3, 5, 8, 7, 8}, 6, 5}
		p2 := Polynomial{[]int{4, 0, 5, 3, 3, 9, 6}, 7, 6}

		got := f.GCD(p1, p2)
		want := Polynomial{[]int{7, 0, 10}, 3, 2}
		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.ToString(), got.ToString())
		}
	})

	t.Run("calculating GCD = 1 of two reducible polinomials", func(t *testing.T) {
		f := SimpleField{11, false}
		p1 := Polynomial{[]int{10, 0, 10, 2, 5}, 5, 4}
		p2 := Polynomial{[]int{1, 2, 4, 3}, 4, 3}

		got := f.GCD(p1, p2)
		want := Polynomial{[]int{1}, 1, 0}
		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.ToString(), got.ToString())
		}
	})
}

func TestExtendedField_ModInverse(t *testing.T) {
	t.Run("calculationg inverse element #1", func(t *testing.T) {
		f := ExtendedField{
			SimpleField{37, false}, 37, 8, // Простое поле, простое число, степень расширения
			Polynomial{[]int{23, 28, 26, 30, 22, 7, 9, 25, 1}, 9, 8},
			false,
		}
		poly := Polynomial{[]int{2, 4, 10, 6, 18}, 5, 4}

		inverse, _ := f.modInverse(poly)
		checked := f.MulPolynomials(poly, inverse)

		if !checked.Equals(Polynomial{[]int{1}, 1, 0}) {
			t.Errorf("Got %s", inverse.Sprint())
		}
	})

	t.Run("inverse of a polynomial with higher degree", func(t *testing.T) {
		f := ExtendedField{
			SimpleField{19, false}, 19, 4,
			Polynomial{[]int{1, 0, 0, 1}, 4, 3}, // Неприводимый многочлен
			false,
		}
		poly := Polynomial{[]int{5, 3, 7}, 3, 2}

		inverse, _ := f.modInverse(poly)
		checked := f.MulPolynomials(poly, inverse)

		if !checked.Equals(Polynomial{[]int{1}, 1, 0}) {
			t.Errorf("Failed to find inverse for %s. Got %s", poly.Sprint(), inverse.Sprint())
		}
	})

	t.Run("inverse of a constant polynomial", func(t *testing.T) {
		f := ExtendedField{
			SimpleField{11, false}, 11, 3,
			Polynomial{[]int{1, 1, 0, 1}, 4, 3},
			false,
		}
		poly := Polynomial{[]int{3}, 1, 0}

		inverse, _ := f.modInverse(poly)
		checked := f.MulPolynomials(poly, inverse)

		if !checked.Equals(Polynomial{[]int{1}, 1, 0}) {
			t.Errorf("Failed to find inverse for constant polynomial %s. Got %s", poly.Sprint(), inverse.Sprint())
		}
	})

	t.Run("non-invertible polynomial", func(t *testing.T) {
		f := ExtendedField{
			SimpleField{7, false}, 7, 2,
			Polynomial{[]int{1, 0, 1}, 3, 2},
			false,
		}
		poly := Polynomial{[]int{}, 0, -1} // Нулевой многочлен

		_, err := f.modInverse(poly)
		if err == nil {
			t.Errorf("Expected error for non-invertible polynomial %s", poly.Sprint())
		}
	})

	t.Run("polynomial equal to modulus", func(t *testing.T) {
		f := ExtendedField{
			SimpleField{13, false}, 13, 5,
			Polynomial{[]int{1, 1, 0, 0, 1}, 5, 4},
			false,
		}
		poly := Polynomial{[]int{1, 1, 0, 0, 1}, 5, 4} // Полный модуль

		_, err := f.modInverse(poly)
		if err == nil {
			t.Errorf("Expected error for polynomial equal to modulus %s", poly.Sprint())
		}
	})

	t.Run("inverse of irreducible polynomial", func(t *testing.T) {
		f := ExtendedField{
			SimpleField{17, false}, 17, 3,
			Polynomial{[]int{1, 0, 1, 1}, 4, 3},
			false,
		}
		poly := Polynomial{[]int{1, 0, 0}, 3, 2} // Пример простого многочлена

		inverse, _ := f.modInverse(poly)
		checked := f.MulPolynomials(poly, inverse)

		if !checked.Equals(Polynomial{[]int{1}, 1, 0}) {
			t.Errorf("Failed to find inverse for irreducible polynomial %s. Got %s", poly.Sprint(), inverse.Sprint())
		}
	})
}

func BenchmarkSimpleField_IsIrreducible(b *testing.B) {
	f := SimpleField{3, false}
	poly := Polynomial{[]int{2, 2, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, 101, 100}
	for i := 0; i < 10; i++ {
		f.IsIrreducible(poly)
	}
}
