package polygfgo

type Num uint

const UNIT_DEGREE = 1

type FieldInterface interface {
	// AddPolynomials(p1, p2 Polynomial) Polynomial
	// SubPolynomials(p1, p2 Polynomial) Polynomial
	// MulPolynomials(p1, p2 Polynomial) Polynomial
	// DivPolynomials(p1, p2 Polynomial) Polynomial
	// Normalize(poly Polynomial) Polynomial
	// ToString() string
}

func FieldFactory(p, m uint, generator Polynomial, enableLogging bool) (field FieldInterface) {
	if m == 1 {
		field = SimpleField{p, enableLogging}
		return
	}
	field = ExtendedField{p, m, generator, enableLogging}
	return
}

// Представление конечного поля GF(p)
type SimpleField struct {
	p             uint
	enableLogging bool
}

func (f *SimpleField) Normalize(poly *Polynomial) (product Polynomial) {
	return
}

func (f *SimpleField) AddPolynomials(p1, p2 Polynomial) (product Polynomial) {
	product = p1.Add(p2)

	f.Normalize(&product)

	return
}

func (f *SimpleField) SubPolynomials(p1, p2 Polynomial) (product Polynomial) {
	return
}

// Представление конечного поля GF(q), q = p^m
type ExtendedField struct {
	// q             uint
	p, m          uint
	generator     Polynomial
	enableLogging bool
}
