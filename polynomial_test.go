package polygfgo

import "testing"

func TestNewPolynomialNoReverse(t *testing.T) {
	t.Run("create a polynomial without trailing zeros", func(t *testing.T) {
		coefs := []int{1, 2, 3}
		got := newPolynomialNoReverse(coefs)
		want := Polynomial{[]int{1, 2, 3}, 3, 2}

		if !got.Equals(want) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("create a polynomial with trailing zeros", func(t *testing.T) {
		coefs := []int{4, 5, 0, 0, 0}
		got := newPolynomialNoReverse(coefs)
		want := Polynomial{[]int{4, 5}, 2, 1}

		if !got.Equals(want) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("create a polynomial with all zero coefficients", func(t *testing.T) {
		coefs := []int{0, 0, 0}
		got := newPolynomialNoReverse(coefs)
		want := Polynomial{[]int{}, 0, -1}

		if !got.Equals(want) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("create a polynomial with single non-zero coefficient", func(t *testing.T) {
		coefs := []int{7}
		got := newPolynomialNoReverse(coefs)
		want := Polynomial{[]int{7}, 1, 0}

		if !got.Equals(want) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("create a polynomial from empty slice", func(t *testing.T) {
		coefs := []int{}
		got := newPolynomialNoReverse(coefs)
		want := Polynomial{[]int{}, 0, -1}

		if !got.Equals(want) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})
}

func TestEquals(t *testing.T) {
	t.Run("collection with different coefficients", func(t *testing.T) {
		p1 := Polynomial{[]int{1, 2, 3}, 3, 2}
		p2 := Polynomial{[]int{1, 2, 4}, 3, 2}

		got := p1.Equals(p2)
		want := false

		if got != want {
			t.Errorf("Expected %t but got %t", want, got)
		}
	})

	t.Run("collection with different lengths", func(t *testing.T) {
		p1 := Polynomial{[]int{1, 2, 3}, 3, 2}
		p2 := Polynomial{[]int{1, 2}, 2, 1}

		got := p1.Equals(p2)
		want := false

		if got != want {
			t.Errorf("Expected %t but got %t", want, got)
		}
	})

	t.Run("collection with identical coefficients and degree", func(t *testing.T) {
		p1 := Polynomial{[]int{1, -2, 3}, 3, 2}
		p2 := Polynomial{[]int{1, -2, 3}, 3, 2}

		got := p1.Equals(p2)
		want := true

		if got != want {
			t.Errorf("Expected %t but got %t", want, got)
		}
	})

	t.Run("collection with identical coefficients but different degree", func(t *testing.T) {
		p1 := Polynomial{[]int{1, -2, 3}, 3, 1}
		p2 := Polynomial{[]int{1, -2, 3}, 3, 2}

		got := p1.Equals(p2)
		want := false

		if got != want {
			t.Errorf("Expected %t but got %t", want, got)
		}
	})

	t.Run("empty polynomial comparison", func(t *testing.T) {
		p1 := Polynomial{[]int{}, 0, 0}
		p2 := Polynomial{[]int{}, 0, 0}

		got := p1.Equals(p2)
		want := true

		if got != want {
			t.Errorf("Expected %t but got %t", want, got)
		}
	})
}

func TestNormalize(t *testing.T) {
	t.Run("normalize a polynomial with trailing zeros", func(t *testing.T) {
		poly := Polynomial{[]int{42, -1, 0, 5, 0, 0}, 6, 3}

		got := poly.Normalize()
		want := Polynomial{[]int{42, -1, 0, 5}, 4, 3}

		if !got.Equals(want) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("normalize a polynomial without trailing zeros", func(t *testing.T) {
		poly := Polynomial{[]int{5, -2, 3}, 3, 2}

		got := poly.Normalize()
		want := Polynomial{[]int{5, -2, 3}, 3, 2}

		if !got.Equals(want) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("normalize an already empty polynomial", func(t *testing.T) {
		poly := Polynomial{[]int{}, 0, 0}

		got := poly.Normalize()
		want := Polynomial{[]int{}, 0, -1}

		if !got.Equals(want) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("normalize a polynomial with only zeros", func(t *testing.T) {
		poly := Polynomial{[]int{0, 0, 0, 0}, 4, 3}

		got := poly.Normalize()
		want := Polynomial{[]int{}, 0, -1}

		if !got.Equals(want) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("normalize a single non-zero coefficient", func(t *testing.T) {
		poly := Polynomial{[]int{7}, 1, 0}

		got := poly.Normalize()
		want := Polynomial{[]int{7}, 1, 0}

		if !got.Equals(want) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("adding with different lengths", func(t *testing.T) {
		p1 := Polynomial{[]int{0, 2301, 0, 19, -2}, 5, 4}
		p2 := Polynomial{[]int{21, -9, -4, 0, 5, 0, 3}, 7, 6}

		got := p1.Add(p2)
		want := Polynomial{[]int{21, 2292, -4, 19, 3, 0, 3}, 7, 6}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("adding two empty polynomials", func(t *testing.T) {
		p1 := Polynomial{[]int{}, 0, -1}
		p2 := Polynomial{[]int{}, 0, -1}

		got := p1.Add(p2)
		want := Polynomial{[]int{}, 0, -1}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("adding polynomial to zero polynomial", func(t *testing.T) {
		p1 := Polynomial{[]int{1, 2, 3}, 3, 2}
		p2 := Polynomial{[]int{0, 0, 0}, 3, 2}

		got := p1.Add(p2)
		want := Polynomial{[]int{1, 2, 3}, 3, 2}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("adding zero polynomial to polynomial", func(t *testing.T) {
		p1 := Polynomial{[]int{0, 0, 0}, 3, 2}
		p2 := Polynomial{[]int{1, 2, 3}, 3, 2}

		got := p1.Add(p2)
		want := Polynomial{[]int{1, 2, 3}, 3, 2}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("adding negative coefficients", func(t *testing.T) {
		p1 := Polynomial{[]int{-1, -2, -3}, 3, 2}
		p2 := Polynomial{[]int{-4, -5, -6}, 3, 2}

		got := p1.Add(p2)
		want := Polynomial{[]int{-5, -7, -9}, 3, 2}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})
}

func TestSub(t *testing.T) {
	t.Run("subtraction of polynomials with different lengths", func(t *testing.T) {
		p1 := Polynomial{[]int{0, 2301, 0, 19, -2}, 5, 4}
		p2 := Polynomial{[]int{21, -9, -4, 0, 5, 0, 3}, 7, 6}

		got := p1.Sub(p2)
		want := Polynomial{[]int{-21, 2310, 4, 19, -7, 0, -3}, 7, 6}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("subtraction of identical polynomials", func(t *testing.T) {
		p1 := Polynomial{[]int{1, 2, 3}, 3, 2}
		p2 := Polynomial{[]int{1, 2, 3}, 3, 2}

		got := p1.Sub(p2)
		want := Polynomial{[]int{}, 0, -1}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("subtraction of zero polynomial", func(t *testing.T) {
		p1 := Polynomial{[]int{1, 2, 3}, 3, 2}
		p2 := Polynomial{[]int{0, 0, 0}, 3, 2}

		got := p1.Sub(p2)
		want := Polynomial{[]int{1, 2, 3}, 3, 2}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("subtraction resulting in negative coefficients", func(t *testing.T) {
		p1 := Polynomial{[]int{1, 2, 3}, 3, 2}
		p2 := Polynomial{[]int{3, 4, 5}, 3, 2}

		got := p1.Sub(p2)
		want := Polynomial{[]int{-2, -2, -2}, 3, 2}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("subtraction of empty polynomials", func(t *testing.T) {
		p1 := Polynomial{[]int{}, 0, -1}
		p2 := Polynomial{[]int{}, 0, -1}

		got := p1.Sub(p2)
		want := Polynomial{[]int{}, 0, -1}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})
}

func TestMul(t *testing.T) {
	t.Run("multiply of polynomials with different lengths", func(t *testing.T) {
		p1 := Polynomial{[]int{9, 0, 5, -2, 8}, 5, 4}
		p2 := Polynomial{[]int{7, -6, 1}, 3, 2}

		got := p1.Mul(p2)
		want := Polynomial{[]int{63, -54, 44, -44, 73, -50, 8}, 7, 6}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("multiply a polynomial by zero polynomial", func(t *testing.T) {
		p1 := Polynomial{[]int{3, 5, 7}, 3, 2}
		p2 := Polynomial{[]int{}, 0, -1}

		got := p1.Mul(p2)
		want := Polynomial{[]int{}, 0, -1}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("multiply a polynomial by one", func(t *testing.T) {
		p1 := Polynomial{[]int{4, -3, 6}, 3, 2}
		p2 := Polynomial{[]int{1}, 1, 0}

		got := p1.Mul(p2)
		want := Polynomial{[]int{4, -3, 6}, 3, 2}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("multiply a polynomial by a constant", func(t *testing.T) {
		p1 := Polynomial{[]int{2, -1, 4}, 3, 2}
		p2 := Polynomial{[]int{3}, 1, 0}

		got := p1.Mul(p2)
		want := Polynomial{[]int{6, -3, 12}, 3, 2}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("multiply two zero polynomials", func(t *testing.T) {
		p1 := Polynomial{[]int{}, 0, -1}
		p2 := Polynomial{[]int{}, 0, -1}

		got := p1.Mul(p2)
		want := Polynomial{[]int{}, 0, -1}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("multiply two polynomials with one term each", func(t *testing.T) {
		p1 := Polynomial{[]int{5}, 1, 0}
		p2 := Polynomial{[]int{4}, 1, 0}

		got := p1.Mul(p2)
		want := Polynomial{[]int{20}, 1, 0}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})
}

func TestMulScalar(t *testing.T) {
	t.Run("multiply a polynomial with trailing zero by positive number", func(t *testing.T) {
		poly := Polynomial{[]int{0, 10, -6, 1, 7, 0}, 6, 5}
		alpha := 13

		got := poly.MulScalar(alpha)
		want := Polynomial{[]int{0, 130, -78, 13, 91}, 5, 4}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("multiply a polynomial by zero", func(t *testing.T) {
		poly := Polynomial{[]int{3, 5, -2, 0}, 4, 3}
		alpha := 0

		got := poly.MulScalar(alpha)
		want := Polynomial{[]int{}, 0, -1}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("multiply a polynomial by one", func(t *testing.T) {
		poly := Polynomial{[]int{-7, 2, 0, 5}, 4, 3}
		alpha := 1

		got := poly.MulScalar(alpha)
		want := Polynomial{[]int{-7, 2, 0, 5}, 4, 3}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("multiply a polynomial by negative number", func(t *testing.T) {
		poly := Polynomial{[]int{6, -3, 0, 4}, 4, 3}
		alpha := -2

		got := poly.MulScalar(alpha)
		want := Polynomial{[]int{-12, 6, 0, -8}, 4, 3}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("multiply an empty polynomial by any number", func(t *testing.T) {
		poly := Polynomial{[]int{}, 0, -1}
		alpha := 5

		got := poly.MulScalar(alpha)
		want := Polynomial{[]int{}, 0, -1}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})

	t.Run("multiply a single-term polynomial by positive number", func(t *testing.T) {
		poly := Polynomial{[]int{3}, 1, 0}
		alpha := 4

		got := poly.MulScalar(alpha)
		want := Polynomial{[]int{12}, 1, 0}

		if !(got.Equals(want)) {
			t.Errorf("Expected %s but got %s", want.Sprint(), got.Sprint())
		}
	})
}

func TestSprint(t *testing.T) {
	t.Run("to string test for {-2, 19, 0, 2301, 0}", func(t *testing.T) {
		poly := Polynomial{[]int{-2, 19, 0, 2301, 0}, 5, 4}

		got := poly.Sprint()
		want := "[-2 19 0 2301 0]5:4"

		if got != want {
			t.Errorf("Expected %s but got %s", want, got)
		}
	})

	t.Run("to string test for empty polynomial", func(t *testing.T) {
		poly := Polynomial{[]int{}, 0, -1}

		got := poly.Sprint()
		want := "[]0:-1"

		if got != want {
			t.Errorf("Expected %s but got %s", want, got)
		}
	})

	t.Run("to string test for single coefficient", func(t *testing.T) {
		poly := Polynomial{[]int{42}, 1, 0}

		got := poly.Sprint()
		want := "[42]1:0"

		if got != want {
			t.Errorf("Expected %s but got %s", want, got)
		}
	})
}
