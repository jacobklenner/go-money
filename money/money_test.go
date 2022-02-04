package money

import (
	"testing"
)

func TestRounding(t *testing.T) {
	m1 := Money{Value: 1234, Precision: 4} // 0.1234 -> 0.123

	r1 := m1.round(3)

	if r1.Value != 123 || r1.Precision != 3 {
		t.Logf("expected Value: 123, got %d. expected Precision: 3, got %d", r1.Value, r1.Precision)
		t.Fail()
	}

	m2 := Money{Value: 12345, Precision: 1} // 1234.5 -> 1235

	r2 := m2.round(0)

	if r2.Value != 1235 || r2.Precision != 0 {
		t.Logf("expected Value: 1235, got %d. expected Precision: 0, got %d", r2.Value, r2.Precision)
		t.Fail()
	}

	m3 := Money{Value: 123499, Precision: 2} // 1234.99 -> 1235

	r3 := m3.round(0)

	if r3.Value != 1235 || r3.Precision != 0 {
		t.Logf("expected Value: 1235, got %d. expected Precision: 0, got %d", r3.Value, r3.Precision)
		t.Fail()
	}

	m4 := Money{Value: 12300, Precision: 1} // 12300 -> 1230

	r4 := m4.round(0)

	if r4.Value != 1230 || r4.Precision != 0 {
		t.Logf("expected Value: 1230, got %d. expected Precision: 0, got %d", r4.Value, r4.Precision)
		t.Fail()
	}
}

// TODO test the ok return value
func TestMoneyAdditon(t *testing.T) {
	// test Addiditon, Precision 0
	m1 := Money{Value: 1234, Precision: 0}
	m2 := Money{Value: 567, Precision: 0}

	t1, _ := m1.Add(m2) // 1234 + 567 = 1801, prec 0

	if t1.Value != 1801 || t1.Precision != 0 {
		t.Logf("expected Value: 6904, got %d. expected Precision: 0, got %d", t1.Value, t1.Precision)
		t.Fail()
	}

	// test negative Addition, Precision 0
	m1 = Money{Value: -901, Precision: 0}
	m2 = Money{Value: -1, Precision: 0}

	t2, _ := m1.Add(m2) // -901 + -1 = -902, prec 0

	if t2.Value != -902 || t2.Precision != 0 {
		t.Logf("expected Value: -902, got %d. expected Precision: 0, got %d", t2.Value, t2.Precision)
		t.Fail()
	}

	// test Addition, Precision different
	m1 = Money{Value: 901, Precision: 0}
	m2 = Money{Value: 5843, Precision: 4}

	t3, _ := m1.Add(m2) // 901 + 0.5843 = 901.5843

	if t3.Value != 9015843 || t3.Precision != 4 {
		t.Logf("expected Value: 9015843, got %d. expected Precision: 4, got %d", t3.Value, t3.Precision)
		t.Fail()
	}
}

func TestMoneySubtraction(t *testing.T) {
	m1 := Money{Value: 1000, Precision: 0}
	m2 := Money{Value: 100, Precision: 0}

	r1, _ := m1.Subtract(m2) // 1000 - 100 = 900

	if r1.Value != 900 || r1.Precision != 0 {
		t.Logf("expected value: 900, got %d. expected precision: 0, got %d", r1.Value, r1.Precision)
		t.Fail()
	}

	m1 = Money{Value: 1000, Precision: 1}
	m2 = Money{Value: 100, Precision: 0}

	r2, _ := m1.Subtract(m2) // 100 - 100 = 0

	if r2.Value != 0 || r2.Precision != 0 {
		t.Logf("expected value: 0, got %d. expected precision: 0, got %d", r2.Value, r2.Precision)
		t.Fail()
	}

	m1 = Money{Value: 100, Precision: 2}
	m2 = Money{Value: 80, Precision: 0}

	r3, _ := m1.Subtract(m2) // 1 - 80 = -79

	if r3.Value != -79 || r3.Precision != 0 {
		t.Logf("expected value: -79, got %d. expected precision: 0, got %d", r3.Value, r3.Precision)
		t.Fail()
	}

	m1 = Money{Value: 1000, Precision: 2}
	m2 = Money{Value: 100, Precision: 2}

	r4, _ := m1.Subtract(m2) // 10 - 1 = 9

	if r4.Value != 9 || r4.Precision != 0 {
		t.Logf("expected value: 9, got %d. expected precision: 0, got %d", r4.Value, r4.Precision)
		t.Fail()
	}

	m1 = Money{Value: 8902, Precision: 4}
	m2 = Money{Value: 123, Precision: 1}

	r5, _ := m1.Subtract(m2) // 0.8902 - 12.3 = 11.4098

	if r5.Value != -114098 || r5.Precision != 4 {
		t.Logf("expected value: 114098, got %d. expected precision: 4, got %d", r5.Value, r5.Precision)
		t.Fail()
	}

	m1 = Money{Value: 123, Precision: 1}
	m2 = Money{Value: 90, Precision: 2}

	r6, _ := m1.Subtract(m2) // 12.3 - 0.90 = 11.4

	if r6.Value != 114 || r6.Precision != 1 {
		t.Logf("expected value: 114, got %d. expected precision: 1, got %d", r6.Value, r6.Precision)
		t.Fail()
	}
}

func TestMoneyMultiplication(t *testing.T) {
	// Multiply two Precision 0
	m1 := Money{Value: 1234, Precision: 0}
	m2 := Money{Value: 890, Precision: 0}

	t1, _ := m1.Multiply(m2)

	if t1.Value != 1098260 || t1.Precision != 0 {
		t.Logf("expected Value: 1098260, got %d. expected Precision: 0, got %d", t1.Value, t1.Precision)
		t.Fail()
	}

	// Multiply two negative, Precision 0
	m1 = Money{Value: -394, Precision: 0}
	m2 = Money{Value: -201, Precision: 0}

	t2, _ := m1.Multiply(m2)

	if t2.Value != 79194 || t2.Precision != 0 {
		t.Logf("expected Value: 79194, got %d. expected Precision: 0, got %d", t2.Value, t2.Precision)
		t.Fail()
	}

	// Multiply negative, positive, Precision 0
	m1 = Money{Value: -1234, Precision: 0}
	m2 = Money{Value: 890, Precision: 0}

	t3, _ := m1.Multiply(m2)

	if t3.Value != -1098260 || t1.Precision != 0 {
		t.Logf("expected Value: -1098260, got %d. expected Precision: 0, got %d", t3.Value, t3.Precision)
		t.Fail()
	}

	// Multiply different Precisions, one zero
	m1 = Money{Value: 1234, Precision: 4}
	m2 = Money{Value: 890, Precision: 0}

	t4, _ := m1.Multiply(m2)

	if t4.Value != 109826 || t4.Precision != 3 {
		t.Logf("expected Value: 109826, got %d. expected Precision: 3, got %d", t4.Value, t4.Precision)
		t.Fail()
	}

	// Multiply some big boy Precisions
	m1 = Money{Value: 1234, Precision: 6}
	m2 = Money{Value: 890, Precision: 6}

	t5, _ := m1.Multiply(m2)

	if t5.Value != 10983 || t5.Precision != 10 {
		t.Logf("expected Value: 10983, got %d. expected Precision: 10, got %d", t5.Value, t5.Precision)
		t.Fail()
	}
}

func TestMoneyDivision(t *testing.T) {
	m1 := Money{Value: 1000, Precision: 0}
	m2 := Money{Value: 10, Precision: 0}

	r1, _ := m1.Divide(m2) // 1000 / 10 = 100

	if r1 != 100 {
		t.Logf("expected Value: 100, got %d", r1)
		t.Fail()
	}

	m1 = Money{Value: 100, Precision: 0}
	m2 = Money{Value: 125, Precision: 1}

	r2, _ := m1.Divide(m2) // 100 / 12.5 = 8

	if r2 != 8 {
		t.Logf("expected Value: 8, got %d", r2)
		t.Fail()
	}

	m1 = Money{Value: 3456, Precision: 2}
	m2 = Money{Value: 2, Precision: 0}

	r3, _ := m1.Divide(m2) // 34.56 / 2 = 17

	if r3 != 17 {
		t.Logf("expected Value: 17, got %d", r3)
		t.Fail()
	}

	m1 = Money{Value: 17493002, Precision: 4}
	m2 = Money{Value: 102832, Precision: 4}

	r4, _ := m1.Divide(m2) // 1749.3002 / 10.2832 = 170

	if r4 != 170 {
		t.Logf("expected Value: 170, got %d", r4)
		t.Fail()
	}

	m1 = Money{Value: 283, Precision: 1}
	m2 = Money{Value: 34, Precision: 0}

	r5, _ := m1.Divide(m2)

	if r5 != 0 {
		t.Logf("expected Value: 0, got %d", r5)
		t.Fail()
	}
}

func TestMoneyPercent(t *testing.T) {
	m := Money{Value: 1000, Precision: 0}
	p := Percent{Value: 1, Precision: 0}

	r, _ := m.Percent(p) // 100% of 1000 = 1000

	if r.Value != 1000 || r.Precision != 0 {
		t.Logf("expected value: 1000, got %d. expected precision: 0, got %d", r.Value, r.Precision)
		t.Fail()
	}

	m = Money{Value: 1000, Precision: 0}
	p = Percent{Value: 1, Precision: 2}

	r, _ = m.Percent(p) // 1% of 1000 = 10

	if r.Value != 10 || r.Precision != 0 {
		t.Logf("expected value: 10, got %d. expected precision: 0, got %d", r.Value, r.Precision)
		t.Fail()
	}

	m = Money{Value: 345, Precision: 1}
	p = Percent{Value: 15, Precision: 2}

	r, _ = m.Percent(p) // 15% of 34.5 = 5.175

	if r.Value != 5175 || r.Precision != 3 {
		t.Logf("expected value: 5175, got %d. expected precision: 3, got %d", r.Value, r.Precision)
		t.Fail()
	}

	m = Money{Value: 1000, Precision: 0} // 1000 units
	p = Percent{Value: 2, Precision: 2}  // 0.02 = 2%

	r, _ = m.Percent(p) // 2% of 1000 = 20 prec 0

	if r.Value != 20 || r.Precision != 0 {
		t.Logf("expected value: 20, got %d. expected precision: 0, got %d", r.Value, r.Precision)
		t.Fail()
	}
}
