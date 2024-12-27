package polygfgo

import (
	"reflect"
	"testing"
)

func TestExpand(t *testing.T) {
	t.Run("expand slice to larger size", func(t *testing.T) {
		s := []int{1, 2, 3}
		n := 6

		got := expand(s, n)
		want := []int{1, 2, 3, 0, 0, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})

	t.Run("no expansion needed when n equals slice length", func(t *testing.T) {
		s := []int{4, 5, 6}
		n := 3

		got := expand(s, n)
		want := []int{4, 5, 6}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})

	t.Run("no expansion needed when n is smaller than slice length", func(t *testing.T) {
		s := []int{7, 8, 9, 10}
		n := 2

		got := expand(s, n)
		want := []int{7, 8, 9, 10}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})

	t.Run("expand an empty slice", func(t *testing.T) {
		s := []int{}
		n := 4

		got := expand(s, n)
		want := []int{0, 0, 0, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})

	t.Run("no expansion for empty slice when n is 0", func(t *testing.T) {
		s := []int{}
		n := 0

		got := expand(s, n)
		want := []int{}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})
}

func TestReverse(t *testing.T) {
	t.Run("reverse a slice with odd number of elements", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		got := reverse(input)
		want := []int{5, 4, 3, 2, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})

	t.Run("reverse a slice with even number of elements", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		got := reverse(input)
		want := []int{4, 3, 2, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})

	t.Run("reverse a single-element slice", func(t *testing.T) {
		input := []int{1}
		got := reverse(input)
		want := []int{1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})

	t.Run("reverse an empty slice", func(t *testing.T) {
		input := []int{}
		got := reverse(input)
		want := []int{}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})

	t.Run("reverse a slice with duplicate elements", func(t *testing.T) {
		input := []int{1, 2, 3, 2, 1}
		got := reverse(input)
		want := []int{1, 2, 3, 2, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})

	t.Run("reverse a slice with negative numbers", func(t *testing.T) {
		input := []int{-1, -2, -3, -4, -5}
		got := reverse(input)
		want := []int{-5, -4, -3, -2, -1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})
}

func TestNextPOT(t *testing.T) {
	t.Run("number is already a power of two", func(t *testing.T) {
		n := 8
		got := nextPOT(n)
		want := 8

		if got != want {
			t.Errorf("Expected %d but got %d", want, got)
		}
	})

	t.Run("number is less than the next power of two", func(t *testing.T) {
		n := 7
		got := nextPOT(n)
		want := 8

		if got != want {
			t.Errorf("Expected %d but got %d", want, got)
		}
	})

	t.Run("number is greater than the previous power of two", func(t *testing.T) {
		n := 33
		got := nextPOT(n)
		want := 64

		if got != want {
			t.Errorf("Expected %d but got %d", want, got)
		}
	})

	t.Run("number is zero", func(t *testing.T) {
		n := 0
		got := nextPOT(n)
		want := 1

		if got != want {
			t.Errorf("Expected %d but got %d", want, got)
		}
	})

	t.Run("number is negative", func(t *testing.T) {
		n := -5
		got := nextPOT(n)
		want := 1

		if got != want {
			t.Errorf("Expected %d but got %d", want, got)
		}
	})

	t.Run("number is already a large power of two", func(t *testing.T) {
		n := 1024
		got := nextPOT(n)
		want := 1024

		if got != want {
			t.Errorf("Expected %d but got %d", want, got)
		}
	})

	t.Run("number is just below a large power of two", func(t *testing.T) {
		n := 1023
		got := nextPOT(n)
		want := 1024

		if got != want {
			t.Errorf("Expected %d but got %d", want, got)
		}
	})

	t.Run("number is between two large powers of two", func(t *testing.T) {
		n := 3000
		got := nextPOT(n)
		want := 4096

		if got != want {
			t.Errorf("Expected %d but got %d", want, got)
		}
	})
}

func TestIsPOT(t *testing.T) {
	t.Run("number is a power of two", func(t *testing.T) {
		n := 8
		got := isPOT(n)
		want := true

		if got != want {
			t.Errorf("Expected %t but got %t for input %d", want, got, n)
		}
	})

	t.Run("number is not a power of two", func(t *testing.T) {
		n := 10
		got := isPOT(n)
		want := false

		if got != want {
			t.Errorf("Expected %t but got %t for input %d", want, got, n)
		}
	})

	t.Run("number is zero", func(t *testing.T) {
		n := 0
		got := isPOT(n)
		want := false

		if got != want {
			t.Errorf("Expected %t but got %t for input %d", want, got, n)
		}
	})

	t.Run("number is negative", func(t *testing.T) {
		n := -8
		got := isPOT(n)
		want := false

		if got != want {
			t.Errorf("Expected %t but got %t for input %d", want, got, n)
		}
	})

	t.Run("number is 1", func(t *testing.T) {
		n := 1
		got := isPOT(n)
		want := true

		if got != want {
			t.Errorf("Expected %t but got %t for input %d", want, got, n)
		}
	})

	t.Run("large power of two", func(t *testing.T) {
		n := 1024
		got := isPOT(n)
		want := true

		if got != want {
			t.Errorf("Expected %t but got %t for input %d", want, got, n)
		}
	})

	t.Run("large number that is not a power of two", func(t *testing.T) {
		n := 3000
		got := isPOT(n)
		want := false

		if got != want {
			t.Errorf("Expected %t but got %t for input %d", want, got, n)
		}
	})
}
