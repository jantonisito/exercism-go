package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
	r float64
	i float64
}

func (n Number) Real() float64 {
	return n.r
}

func (n Number) Imaginary() float64 {
	return n.i
}

func (n1 Number) Add(n2 Number) Number {
	return Number{n1.r + n2.r, n1.i + n2.i}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{n1.r - n2.r, n1.i - n2.i}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{n1.r*n2.r - n1.i*n2.i, n1.r*n2.i + n1.i*n2.r}
}

func (n Number) Times(factor float64) Number {
	return Number{n.r * factor, n.i * factor}
}

// Divide returns the quotient of n1 and n2.
// Multiply both nominator and denominator by the conjugate of the denominator
func (n1 Number) Divide(n2 Number) Number {
	denom := n2.Multiply(n2.Conjugate()).Real()
	prod := n1.Multiply(n2.Conjugate())
	return Number{prod.Real() / denom, prod.Imaginary() / denom}
}

func (n Number) Conjugate() Number {
	return Number{n.r, -n.i}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.Multiply(n.Conjugate()).Real())
}

// Exp returns the exponential of n. Use Euler's formula.
func (n Number) Exp() Number {
	return Number{math.Exp(n.r) * math.Cos(n.i), math.Exp(n.r) * math.Sin(n.i)}
}
