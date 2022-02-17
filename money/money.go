package money

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/shopspring/decimal"
)

type Money struct {
	value    decimal.Decimal
	currency currency
	unit     unit
}

type currency int64

const (
	EUR currency = iota
	USD
)

func parseCurrency(s string) (c currency, ok bool) {
	s = strings.ToUpper(s)
	ok = true
	switch s {
	case "EUR":
		c = EUR
	case "USD":
		c = USD
	default:
		c = -1
		ok = false
	}

	return
}

func (c currency) string() (s string) {
	switch c {
	case EUR:
		s = "EUR"
	case USD:
		s = "USD"
	}
	return
}

type unit int64

const (
	CENT unit = iota
	EURO
	DOLLAR
)

func parseUnit(s string) (u unit, ok bool) {
	s = strings.ToUpper(s)
	ok = true
	switch s {
	case "CENT":
		u = CENT
	case "DOLLAR":
		u = DOLLAR
	case "EURO":
		u = EURO
	default:
		u = -1
		ok = false
	}

	return
}

func (u unit) string() (s string) {
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

func defaultMoney() Money {
	return Money{
		currency: -1,
		unit:     -1,
	}
}

func new(v decimal.Decimal, c string, u string) Money {
	currency, okc := parseCurrency(c)
	unit, oku := parseUnit(u)

	if !okc || !oku {
		return defaultMoney()
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
		return defaultMoney(), false
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
		return defaultMoney(), false
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
		return defaultMoney(), false
	}

	return Money{
		value:    m1.value.Mul(m2.value),
		unit:     m1.unit,
		currency: m1.currency,
	}, true
}

func (m Money) MultiplyFloat(f float64) Money {
	return Money{
		value:    m.value.Mul(decimal.NewFromFloat(f)),
		unit:     m.unit,
		currency: m.currency,
	}
}

// returns m1 / m2, ok
func (m1 Money) Divide(m2 Money) (Money, bool) {
	if !m1.sameUnit(m2) {
		return defaultMoney(), false
	}

	return Money{
		value:    m1.value.Div(m2.value),
		unit:     m1.unit,
		currency: m1.currency,
	}, true
}

func (m1 Money) Quotient(m2 Money) (int64, bool) {
	if !m1.sameUnit(m2) {
		return 0, false
	}

	q, _ := m1.value.QuoRem(m2.value, 0)

	return q.IntPart(), true
}

func (m Money) QutoientFloat(f float64) int64 {
	q, _ := m.value.QuoRem(decimal.NewFromFloat(f), 0)
	return q.IntPart()
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
	var tmp struct {
		Value    decimal.Decimal `json:"value"`
		Currency string          `json:"currency"`
		Unit     string          `json:"unit"`
	}
	err := json.Unmarshal(data, &tmp)

	if err != nil {
		return nil
	}

	c, okc := parseCurrency(tmp.Currency)
	u, oku := parseUnit(string(tmp.Unit))

	// must have a currency
	if !okc {
		return nil
	}

	// default to base unit if undefined
	if !oku {
		switch c {
		case EUR:
			u = EURO
		case USD:
			u = DOLLAR
		default:
			err = fmt.Errorf("could not derive default unit for provided currency %s", c.string())
		}
	}

	if err != nil {
		return err
	}

	m.value = tmp.Value
	m.currency = c
	m.unit = u

	return nil
}
