# polygfgo

`polygfgo` is a Go library for performing polynomial arithmetic over finite fields (Galois Fields). It provides a simple and efficient way to work with mathematical constructs essential in areas like cryptography, error-correcting codes (e.g., Reed-Solomon codes), and digital signal processing.

The library allows you to define a finite field GF(pⁿ) and then create and manipulate polynomials with coefficients from that field.

## Features

- **Finite Field Arithmetic**: Create and operate within any finite field GF(pⁿ). All basic arithmetic operations (`Add`, `Subtract`, `Multiply`, `Inverse`, `Divide`) on field elements are supported.
    
- **Polynomial Operations**:
    
    - Create polynomials with coefficients from a specified finite field.
        
    - Perform basic polynomial arithmetic: `Add`, `Subtract`, `Multiply`.
        
    - Evaluate a polynomial at a specific point in the field.
        
- **Ease of Use**: A clean and straightforward API for defining fields and polynomials.
    

## Installation

To use `polygfgo` in your project, install it using `go get`:

text

`go get github.com/untibullet/polygfgo`

## Usage

Here is a basic example of how to use the library. We will create a finite field GF(2³), define two polynomials, and then perform addition and multiplication.

```go
package main

import (
	"fmt"
	"log"

	"github.com/untibullet/polygfgo"
)

func main() {
	// 1. Create a finite field GF(2^3).
	// The library automatically finds a suitable irreducible polynomial.
	field, err := polygfgo.NewField(2, 3)
	if err != nil {
		log.Fatalf("Failed to create field: %v", err)
	}

	// 2. Define two polynomials over the field GF(2^3).
	// P(x) = x^2 + 1
	p1 := polygfgo.NewPolynomial(field, []uint64{1, 0, 1})

	// Q(x) = x + 1
	p2 := polygfgo.NewPolynomial(field, []uint64{1, 1})

	fmt.Printf("Field: GF(%d^%d)\n", field.P, field.N)
	fmt.Printf("P(x) = %s\n", p1)
	fmt.Printf("Q(x) = %s\n", p2)

	// 3. Add the two polynomials: R(x) = P(x) + Q(x)
	sum := p1.Add(p2)
	fmt.Printf("P(x) + Q(x) = %s\n", sum)

	// 4. Multiply the two polynomials: S(x) = P(x) * Q(x)
	product := p1.Mul(p2)
	fmt.Printf("P(x) * Q(x) = %s\n", product)
}
```
## Explanation

1. **Field Creation**: `polygfgo.NewField(2, 3)` initializes the finite field GF(2³). The elements of this field are polynomials of degree less than 3 with coefficients in GF(2).
    
2. **Polynomial Definition**: `polygfgo.NewPolynomial(field, coeffs)` creates a new polynomial. The coefficients are passed as a slice of `uint64`, where each element is a value in the base field (GF(2) in this case).
    
    - `[]uint64{1, 0, 1}` corresponds to the polynomial `1*x^2 + 0*x^1 + 1*x^0`.
        
3. **Addition**: The `Add` method performs polynomial addition. In GF(2), this is equivalent to a bitwise XOR on the coefficients.
    
    - `(x^2 + 1) + (x + 1) = x^2 + x`
        
4. **Multiplication**: The `Mul` method performs polynomial multiplication. The result is taken modulo the field's irreducible polynomial.
    
    - `(x^2 + 1) * (x + 1) = x^3 + x^2 + x + 1`.
        
    - Within GF(2³) using the irreducible polynomial `x^3 + x + 1`, `x^3` is equivalent to `x + 1`.
        
    - The result simplifies to `(x + 1) + x^2 + x + 1 = x^2`.
        

## API Overview

The library exposes two main types: `Field` and `Polynomial`.

## `Field`

Represents a finite field GF(pⁿ).

- `NewField(p, n uint64) (*Field, error)`: Creates a new finite field.
    
- `Add(a, b uint64) uint64`: Adds two field elements.
    
- `Mul(a, b uint64) uint64`: Multiplies two field elements.
    
- `Inv(a uint64) uint64`: Computes the multiplicative inverse of a field element.
    

## `Polynomial`

Represents a polynomial with coefficients in a given `Field`.

- `NewPolynomial(field *Field, coeffs []uint64) *Polynomial`: Creates a new polynomial.
    
- `Add(other *Polynomial) *Polynomial`: Returns the sum of two polynomials.
    
- `Sub(other *Polynomial) *Polynomial`: Returns the difference of two polynomials.
    
- `Mul(other *Polynomial) *Polynomial`: Returns the product of two polynomials.
    
- `Eval(x uint64) uint64`: Evaluates the polynomial at a point `x` in the field.
    
- `Degree() int`: Returns the degree of the polynomial.
    

## Contributing

Contributions are welcome! If you find a bug or have a feature request, please open an issue. If you would like to contribute code, please open a pull request.

1. Fork the repository.
    
2. Create your feature branch (`git checkout -b feature/my-new-feature`).
    
3. Commit your changes (`git commit -am 'Add some feature'`).
    
4. Push to the branch (`git push origin feature/my-new-feature`).
    
5. Create a new Pull Request.
    

## License

No license has been specified for this project.
