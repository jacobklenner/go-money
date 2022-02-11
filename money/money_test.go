package money

import (
	"math/big"
	"testing"

	"github.com/shopspring/decimal"
)

type moneyTest struct {
	*testing.T
}

func (mt moneyTest) assertMoneyEqual(expected Money, result Money) {
	if !result.exactEqual(expected) {
		mt.Fatalf("expected %+v but got %+v", expected, result)
	}
}

func TestNew(t *testing.T) {
	e := Money{
		value:    decimal.New(100, 1),
		currency: "EUR",
		unit:     "cent",
	}

	r := New(100, 1, "EUR", "cent")

	moneyTest{t}.assertMoneyEqual(e, r)
}

func TestNewEuro(t *testing.T) {
	e := Money{
		value:    decimal.New(345, 2),
		currency: "EUR",
		unit:     "euro",
	}

	r := NewEuro(345, 2)

	moneyTest{t}.assertMoneyEqual(e, r)
}

func TestZeroEuro(t *testing.T) {
	e := Money{
		value:    decimal.Zero,
		unit:     Euro,
		currency: EUR,
	}

	r := ZeroEuro()

	moneyTest{t}.assertMoneyEqual(e, r)
}

func TestZeroUsDollar(t *testing.T) {
	e := Money{
		value:    decimal.Zero,
		unit:     Dollar,
		currency: USD,
	}

	r := ZeroUsDollar()

	moneyTest{t}.assertMoneyEqual(e, r)
}

func TestNewEuroFromFloat32(t *testing.T) {
	e := Money{
		value:    decimal.NewFromFloat32(4503.203),
		currency: "EUR",
		unit:     "euro",
	}

	r := NewEuroFromFloat32(4503.203)

	moneyTest{t}.assertMoneyEqual(e, r)
}

func TestNewEuroFromFloat(t *testing.T) {
	e := Money{
		value:    decimal.NewFromFloat(4503.203),
		currency: "EUR",
		unit:     "euro",
	}

	r := NewEuroFromFloat(4503.203)

	moneyTest{t}.assertMoneyEqual(e, r)
}

func TestNewEuroFromDecimal(t *testing.T) {
	e := Money{
		value:    decimal.New(4539, 3),
		currency: "EUR",
		unit:     "euro",
	}

	r := NewEuroFromDecimal(decimal.New(4539, 3))

	moneyTest{t}.assertMoneyEqual(e, r)
}

func TestNewEuroCent(t *testing.T) {
	e := Money{
		value:    decimal.New(583920, -1),
		currency: "EUR",
		unit:     "cent",
	}

	r := NewEuroCent(583920, -1)

	moneyTest{t}.assertMoneyEqual(e, r)
}

func TestNewEuroCentFromFloat32(t *testing.T) {
	e := Money{
		value:    decimal.NewFromFloat32(58292.304),
		currency: "EUR",
		unit:     "cent",
	}

	r := NewEuroCentFromFloat32(58292.304)

	moneyTest{t}.assertMoneyEqual(e, r)
}

func TestNewEuroCentFromFloat(t *testing.T) {
	e := Money{
		value:    decimal.NewFromFloat(58292.304),
		currency: "EUR",
		unit:     "cent",
	}

	r := NewEuroCentFromFloat(58292.304)

	moneyTest{t}.assertMoneyEqual(e, r)
}

func TestNewEuroCentFromDecimal(t *testing.T) {
	d := decimal.New(4820, 4)

	e := Money{
		value:    d,
		currency: "EUR",
		unit:     "cent",
	}

	r := NewEuroCentFromDecimal(d)

	moneyTest{t}.assertMoneyEqual(e, r)
}

func TestEqual(t *testing.T) {
	m1 := NewEuroFromFloat(5738.0)
	m2 := NewEuroFromFloat(5738.0)

	if !m1.Equal(m2) {
		t.Fatal("expected money would be exact")
	}

	m3 := NewEuroFromFloat(6930.20)
	m4 := NewEuroCentFromFloat(693020)

	if m3.Equal(m4) {
		t.Fatal("expected money would not be exact")
	}

	if m4.Equal(m3) {
		t.Fatal("expected money would not be exact")
	}

	m5 := ZeroUsDollar()
	m6 := ZeroEuro()

	if m5.Equal(m6) {
		t.Fatal("expected money would not be exact")
	}
}

func TestSameCurrencyAddition(t *testing.T) {
	m1 := NewEuroCentFromFloat(14.5677)
	m2 := NewEuroCentFromFloat(100.0000)

	r, ok := m1.Add(m2)

	e := NewEuroCentFromFloat(114.5677)

	if !ok {
		t.Fatalf("incompatable units. expected same units.")
	}

	moneyTest{t}.assertMoneyEqual(e, r)
}

func TestSameCurrencySubtraction(t *testing.T) {
	m1 := NewEuroCentFromFloat(14.5677)
	m2 := NewEuroCentFromFloat(100.0000)

	r, ok := m1.Subtract(m2)

	e := NewEuroCentFromFloat(-85.4323)

	if !ok {
		t.Fatalf("incompatable units. expected same units.")
	}

	moneyTest{t}.assertMoneyEqual(e, r)
}

func TestSameCurrencyMultiplication(t *testing.T) {
	m1 := NewEuroCentFromFloat(14.5677)
	m2 := NewEuroCentFromFloat(100.0000)

	r, ok := m1.Multiply(m2)

	e := NewEuroCentFromFloat(1456.77)

	if !ok {
		t.Fatalf("incompatable units. expected same units.")
	}

	moneyTest{t}.assertMoneyEqual(e, r)
}

func TestSameCurrencyDivision(t *testing.T) {
	m1 := NewEuroCentFromFloat(14.5677)
	m2 := NewEuroCentFromFloat(100.0000)

	r, ok := m1.Divide(m2)

	e := NewEuroCentFromFloat(0.145677)

	if !ok {
		t.Fatalf("incompatable units. expected same units.")
	}

	moneyTest{t}.assertMoneyEqual(e, r)
}

func TestSameCurrencyPercent(t *testing.T) {
	m1 := NewEuroCentFromFloat(14.5677)
	p := decimal.NewFromFloat(0.000653)

	r, ok := m1.Percent(p)

	e := NewEuroCentFromFloat(0.0095127081)

	if !ok {
		t.Fatalf("incompatable units. expected same units.")
	}

	moneyTest{t}.assertMoneyEqual(e, r)
}

func TestDifferentCurrencyAddition(t *testing.T) {
	usd := New(123, 1, "USD", "dollar")
	eur := New(123, 1, "EUR", "euro")

	_, ok := usd.Add(eur)

	if ok {
		t.Fatalf("should not return ok")
	}
}

func TestDifferentCurrencySubtraction(t *testing.T) {
	usd := New(123, 1, "USD", "dollar")
	eur := New(123, 1, "EUR", "euro")

	_, ok := usd.Subtract(eur)

	if ok {
		t.Fatalf("should not return ok")
	}
}

func TestDifferentCurrencyMultiplication(t *testing.T) {
	usd := New(123, 1, "USD", "dollar")
	eur := New(123, 1, "EUR", "euro")

	_, ok := usd.Multiply(eur)

	if ok {
		t.Fatalf("should not return ok")
	}
}

func TestDifferentCurrencyDivision(t *testing.T) {
	usd := New(123, 1, "USD", "dollar")
	eur := New(123, 1, "EUR", "euro")

	_, ok := usd.Divide(eur)

	if ok {
		t.Fatalf("should not return ok")
	}
}

func TestValueToFloat(t *testing.T) {
	e := 573.402
	m := NewEuroCentFromFloat(e)

	v, _ := m.ValueFloat64()

	if v != e {
		t.Fatalf("expected %f but got %f", e, v)
	}
}

func TestValueToBigInt(t *testing.T) {
	e := big.NewInt(68302029485030130)
	m := NewEuroCent(68302029485030130, 0)

	v := m.ValueBigInt()

	if v.String() != e.String() {
		t.Fatalf("expected %d but got %d", e, v)
	}
}

func TestGetCurrency(t *testing.T) {
	m1 := ZeroEuro()
	eur := m1.Currency()

	if eur != "EUR" {
		t.Fatalf(`expected "EUR" but got %s`, eur)
	}

	m2 := ZeroUsDollar()
	usd := m2.Currency()

	if usd != "USD" {
		t.Fatalf(`"expected USD but got %s`, usd)
	}
}

func TestGetUnit(t *testing.T) {
	m1 := ZeroEuro()
	euro := m1.Unit()

	if euro != "euro" {
		t.Fatalf(`expected "euro" but got %s`, euro)
	}

	m2 := NewEuroCent(1, 0)
	cent := m2.Unit()

	if cent != "cent" {
		t.Fatalf(`expected "cent" but got %s`, euro)
	}

	m3 := ZeroUsDollar()
	dollar := m3.Unit()

	if dollar != "dollar" {
		t.Fatalf(`expected "dollar" but got %s`, euro)
	}
}

func TestEqualUnit(t *testing.T) {
	m1 := NewEuroFromFloat(34920.43)
	m2 := NewEuroFromFloat(505)

	if !m1.EqualUnit(m2) {
		t.Fatalf("expected equal units")
	}
}

func TestEqualCurrency(t *testing.T) {
	m1 := NewEuroFromFloat(34920.43)
	m2 := NewEuroCentFromFloat(58302.405)

	if !m1.EqualCurrency(m2) {
		t.Fatalf("expected equal currency")
	}
}
