package money

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/shopspring/decimal"
)

type Currency int64

const (
	def_currency Currency = iota
	EUR
	USD
)

func parseCurrency(s string) (c Currency, err error) {
	s = strings.ToUpper(s)

	switch s {
	case "EUR":
		c = EUR
	case "USD":
		c = USD
	default:
		err = fmt.Errorf("provided currency %s not supported", s)
	}

	return
}

func (c Currency) string() (s string) {
	switch c {
	case EUR:
		s = "EUR"
	case USD:
		s = "USD"
	}
	return
}

type Unit int64

const (
	def_unit Unit = iota
	CENT
	EURO
	DOLLAR
)

func parseUnit(s string) (u Unit, err error) {
	s = strings.ToUpper(s)

	switch s {
	case "CENT":
		u = CENT
	case "DOLLAR":
		u = DOLLAR
	case "EURO":
		u = EURO
	default:
		err = fmt.Errorf("provided unit %s is not supported", s)
	}

	return
}

func (u Unit) string() (s string) {
	switch u {
	case CENT:
		s = "CENT"
	case EURO:
		s = "EURO"
	case DOLLAR:
		s = "DOLLAR"
	}
	return
}

type Money struct {
	value    decimal.Decimal
	currency Currency
	unit     Unit
}

func new(v decimal.Decimal, c string, u string) Money {
	currency, erc := parseCurrency(c)
	unit, eru := parseUnit(u)

	if erc != nil || eru != nil {
		return Money{}
	}

	return Money{
		value:    v,
		currency: currency,
		unit:     unit,
	}
}

func newEuro(v decimal.Decimal) Money {
	return new(v, "EUR", "EURO")
}

func newEuroCent(v decimal.Decimal) Money {
	return new(v, "EUR", "CENT")
}

func newUsdollar(v decimal.Decimal) Money {
	return new(v, "USD", "DOLLAR")
}

func New(val int64, exp int32, c string, u string) Money {
	v := decimal.New(val, exp)

	return new(v, c, u)
}

func NewEuro(val int64, exp int32) Money {
	v := decimal.New(val, exp)

	return newEuro(v)
}

func ZeroEuro() Money {
	v := decimal.Zero

	return newEuro(v)
}

func ZeroUsDollar() Money {
	v := decimal.Zero

	return newUsdollar(v)
}

func NewEuroFromDecimal(d decimal.Decimal) Money {
	return newEuro(d)
}

func NewEuroFromFloat(f float64) Money {
	v := decimal.NewFromFloat(f)

	return newEuro(v)
}

func NewEuroFromFloat32(f float32) Money {
	v := decimal.NewFromFloat32(f)

	return newEuro(v)
}

func NewEuroCent(val int64, exp int32) Money {
	v := decimal.New(val, exp)

	return newEuroCent(v)
}

func NewEuroCentFromDecimal(d decimal.Decimal) Money {
	return newEuroCent(d)
}

func NewEuroCentFromFloat32(f float32) Money {
	v := decimal.NewFromFloat32(f)

	return newEuroCent(v)
}

func NewEuroCentFromFloat(f float64) Money {
	v := decimal.NewFromFloat(f)

	return newEuroCent(v)
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
	if m1.unit == EURO || m1.unit == DOLLAR {
		m1.value = m1.value.Mul(decimal.New(10, 0))
		m1.unit = CENT
	} else if m2.unit == EURO || m2.unit == DOLLAR {
		m2.value = m2.value.Mul(decimal.New(10, 0))
		m2.unit = CENT
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
	return m.currency.string()
}

func (m Money) Unit() string {
	return m.unit.string()
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
	kvp := map[string]string{}
	err := json.Unmarshal(data, &kvp)

	if err != nil {
		return nil
	}

	v, erv := decimal.NewFromString(kvp["value"])
	c, erc := parseCurrency(string(kvp["currency"]))
	u, eru := parseUnit(string(kvp["unit"]))

	if erv != nil || erc != nil || eru != nil {
		return nil
	}

	m.value = v
	m.currency = c
	m.unit = u

	return nil
}
