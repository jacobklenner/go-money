package money

import (
	"github.com/shopspring/decimal"
)

type Currency string

const (
	EUR Currency = "EUR"
	USD Currency = "USD"
)

type Unit string

const (
	Cent   Unit = "cent"
	Euro   Unit = "euro"
	Dollar Unit = "dollar"
)

type Money struct {
	value    decimal.Decimal
	currency Currency
	unit     Unit
}

func New(val int64, exp int32, c string, u string) Money {
	return Money{
		value:    decimal.New(val, exp),
		unit:     Unit(u),
		currency: Currency(c),
	}
}

func NewEuro(val int64, exp int32) Money {
	return Money{
		value:    decimal.New(val, exp),
		unit:     Euro,
		currency: EUR,
	}
}

func NewEuroFromDecimal(d decimal.Decimal) Money {
	return Money{
		value:    d,
		unit:     Euro,
		currency: EUR,
	}
}

func NewEuroFromFloat(f float64) Money {
	return Money{
		value:    decimal.NewFromFloat(f),
		unit:     Euro,
		currency: EUR,
	}
}

func NewEuroFromFloat32(f float32) Money {
	return Money{
		value:    decimal.NewFromFloat32(f),
		unit:     Euro,
		currency: EUR,
	}
}

func NewEuroCent(val int64, exp int32) Money {
	return Money{
		value:    decimal.New(val, exp),
		unit:     Cent,
		currency: EUR,
	}
}

func NewEuroCentFromDecimal(d decimal.Decimal) Money {
	return Money{
		value:    d,
		unit:     Cent,
		currency: EUR,
	}
}

func NewEuroCentFromFloat32(f float32) Money {
	return Money{
		value:    decimal.NewFromFloat32(f),
		unit:     Cent,
		currency: EUR,
	}
}

func NewEuroCentFromFloat(f float64) Money {
	return Money{
		value:    decimal.NewFromFloat(f),
		unit:     Cent,
		currency: EUR,
	}
}

func (m1 Money) Equal(m2 Money) bool {
	return m1.SameUnit(m2) && m1.SameValue(m2)
}

func (m1 Money) SameValue(m2 Money) bool {
	return m1.value.Equal(m2.value)
}

func (m1 Money) SameCurrency(m2 Money) bool {
	return m1.currency == m2.currency
}

func (m1 Money) SameUnit(m2 Money) bool {
	return m1.SameCurrency(m2) && m1.unit == m2.unit
}

// returns m1 + m2, ok
func (m1 Money) Add(m2 Money) (Money, bool) {
	if !m1.SameCurrency(m2) || !m1.SameUnit(m2) {
		return Money{}, false
	}

	return Money{
		value:    m1.value.Add(m2.value),
		unit:     m1.unit,
		currency: m1.currency,
	}, true
}

// returns m1 - m2, ok
func (m1 Money) Subtract(m2 Money) (Money, bool) {
	if !m1.SameUnit(m2) {
		return Money{}, false
	}

	return Money{
		value:    m1.value.Sub(m2.value),
		unit:     m1.unit,
		currency: m1.currency,
	}, true
}

// returns m1 * m2, ok
func (m1 Money) Multiply(m2 Money) (Money, bool) {
	if !m1.SameUnit(m2) {
		return Money{}, false
	}

	return Money{
		value:    m1.value.Mul(m2.value),
		unit:     m1.unit,
		currency: m1.currency,
	}, true
}

// returns m1 / m2, ok
func (m1 Money) Divide(m2 Money) (Money, bool) {
	if !m1.SameUnit(m2) {
		return Money{}, false
	}

	return Money{
		value:    m1.value.Div(m2.value),
		unit:     m1.unit,
		currency: m1.currency,
	}, true
}

// returns m * %p, ok
func (m Money) Percent(p decimal.Decimal) (Money, bool) {
	return Money{
		value:    m.value.Mul(p),
		unit:     m.unit,
		currency: m.currency,
	}, true
}
