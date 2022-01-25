package money

import "math"

type Money struct {
	Value     int64
	Precision int64
	unit      string // TODO implement
	currency  string // TODO implement
}

// precision 2 is default for percentage
// e.g. Percent{Value:5, Precision:2} = 0.05 = 5%
type Percent struct {
	Value     int64
	Precision int64
}

const maxPrecision int64 = 10

// adds t to m
func (m Money) Add(t Money) (Money, bool) {
	a, b, prec := makeEquatable(m, t)

	val := a.Value + b.Value

	res := clean(Money{
		Value:     val,
		Precision: prec,
	})

	if res.Precision > maxPrecision {
		return Money{}, false
	}

	return Money{
		Value:     val,
		Precision: prec,
	}, true
}

// subtracts t from m
func (m Money) Subtract(t Money) (Money, bool) {
	a, b, prec := makeEquatable(m, t)

	val := a.Value - b.Value

	res := clean(Money{
		Value:     val,
		Precision: prec,
	})

	if res.Precision > maxPrecision {
		return Money{}, false
	}

	return clean(Money{
		Value:     val,
		Precision: prec,
	}), true
}

// multiplies m by t
func (m Money) Multiply(t Money) (Money, bool) {

	// result cannot be larger than max int
	if m.Value*t.Value > math.MaxInt {
		return Money{}, false
	}

	val := m.Value * t.Value
	prec := m.Precision + t.Precision

	res := clean(Money{
		Value:     val,
		Precision: prec,
	})

	return res, true
}

// divides m by t, returns the integer number of times t fits into m
// TODO return remainder Money object
func (m Money) Divide(t Money) (int64, bool) {
	a, b, _ := makeEquatable(m, t)
	var val int64
	if a.Value >= b.Value {
		val = a.Value / b.Value
	} else {
		val = 0
	}

	return val, true
}

// calculates the percent of an amount of money
func (m Money) Percent(p Percent) (Money, bool) {
	if m.Value*p.Value > math.MaxInt {
		return Money{}, false
	}

	val := m.Value * p.Value
	prec := m.Precision + p.Precision

	res := clean(Money{
		Value:     val,
		Precision: prec,
	})

	return res, true
}

func clean(m Money) (r Money) {
	// if Value is zero, Precision is 0 by default
	if m.Value == 0 {
		return Money{
			Value:     0,
			Precision: 0,
		}
	}

	// remove trailing zeros, and decrease Precision
	if m.Precision > 0 {
		for math.Mod(float64(m.Value), 10) == 0 {
			m.Value = m.Value / 10
			m.Precision = m.Precision - 1
			if m.Precision == 0 {
				break
			}
		}
	}

	// round to max Precision
	if m.Precision > maxPrecision {
		m = m.round(maxPrecision)
	}

	return m
}

// defualt money rounding
func (m Money) round(p int64) Money {
	if m.Precision < p {
		return m
	}

	pwr := m.Precision - p

	base := int64(math.Pow10(int(pwr)))

	rem := int64(math.Remainder(float64(m.Value), float64(base)))

	val := m.Value

	if rem >= int64(base/2) {
		// round up and decrease Precision
		val = (val + (base - rem)) / base
	}

	if rem < int64(base/2) {
		// round down and decrease Precision
		val = (val - rem) / base
	}

	return Money{
		Value:     val,
		Precision: p,
	}
}

func makeEquatable(a Money, b Money) (Money, Money, int64) {
	// a more precise than b, need to scale up b
	// e.g. a: Money{Value:1234, prec:4} = 0.1234
	// b: Money{Value:1234, prec:2} = 12.34
	var prec int64

	if a.Precision > b.Precision {
		diff := a.Precision - b.Precision
		b.Value = b.Value * int64(math.Pow(10, float64(diff)))
		b.Precision = a.Precision
		prec = a.Precision
	}

	if a.Precision < b.Precision {
		diff := b.Precision - a.Precision
		a.Value = a.Value * int64(math.Pow(10, float64(diff)))
		a.Precision = b.Precision
		prec = b.Precision
	}

	if a.Precision == b.Precision {
		prec = a.Precision
	}

	return a, b, prec
}
