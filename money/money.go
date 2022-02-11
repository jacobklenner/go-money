package money

import (
	"encoding/json"
	"math/big"

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

func ZeroEuro() Money {
	return Money{
		value:    decimal.Zero,
		unit:     Euro,
		currency: EUR,
	}
}

func ZeroUsDollar() Money {
	return Money{
		value:    decimal.Zero,
		unit:     Dollar,
		currency: USD,
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

func (m1 Money) exactEqual(m2 Money) bool {
	return m1.sameValue(m2) && m1.sameUnit(m2)
}

// evalutates whether two money structs share the same value, irrespective of the underlying currency
// e.g. 100 EUR cent == 1 EUR euro
func (m1 Money) Equal(m2 Money) bool {
	if !m1.sameCurrency(m2) {
		return false
	}

	if m1.sameUnit(m2) {
		return m1.sameValue(m2)
	}

	// scale both to cent to evaluate
	if m1.unit == Euro || m1.unit == Dollar {
		m1.value = m1.value.Mul(decimal.New(10, 0))
		m1.unit = Cent
	} else if m2.unit == Euro || m2.unit == Dollar {
		m2.value = m2.value.Mul(decimal.New(10, 0))
		m2.unit = Cent
	}

	return m1.exactEqual(m2)
}

func (m1 Money) EqualCurrency(m2 Money) bool {
	return m1.sameCurrency(m2)
}

// evaluates whether two money structs share the same base units
// e.g. EUR cent != USD cent
// e.g. EUR cent != EUR euro
func (m1 Money) EqualUnit(m2 Money) bool {
	return m1.sameUnit(m2)
}

func (m1 Money) sameValue(m2 Money) bool {
	return m1.value.Equal(m2.value)
}

func (m1 Money) sameCurrency(m2 Money) bool {
	return m1.currency == m2.currency
}

func (m1 Money) sameUnit(m2 Money) bool {
	return m1.sameCurrency(m2) && m1.unit == m2.unit
}

// returns m1 + m2, ok
func (m1 Money) Add(m2 Money) (Money, bool) {
	if !m1.sameUnit(m2) {
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
	if !m1.sameUnit(m2) {
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
	if !m1.sameUnit(m2) {
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
	if !m1.sameUnit(m2) {
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

// returns float64 representation of the money, and flag indicating if this value is exact
func (m Money) ValueFloat64() (val float64, exact bool) {
	val, exact = m.value.Float64()
	return
}

// returns big int representation of the money
func (m Money) ValueBigInt() *big.Int {
	return m.value.BigInt()
}

// returns the currency
func (m Money) Currency() string {
	return string(m.currency)
}

func (m Money) Unit() string {
	return string(m.unit)
}

func (m Money) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"value":    m.value,
		"currency": m.currency,
		"unit":     m.unit,
	})
}

func (m *Money) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	temp := map[string]string{}
	err := json.Unmarshal(data, &temp)

	if err != nil {
		return err
	}

	v, err := decimal.NewFromString(temp["value"])

	if err != nil {
		return err
	}

	m.value = v
	m.currency = Currency(string(temp["currency"]))
	m.unit = Unit(string(temp["unit"]))

	return err
}
