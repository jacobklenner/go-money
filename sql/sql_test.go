package sql

import (
	"fmt"
	"testing"
)

func assert(t *testing.T, exp string, res string) {
	if exp != res {
		t.Logf("expected %s, got %s", exp, res)
		t.Fail()
	}
}

func getTestQuery() Query {

	cols := map[string]Column{
		"account_id": {
			Name: "account_id",
		},
		"value": {
			Name: "value",
		},
		"status": {
			Name: "status",
		},
		"created_at": {
			Name: "created_at",
		},
	}

	return Query{
		Database: "client_db",
		Table:    "accounts",
		Columns:  cols,
	}
}

func TestEmptySelect(t *testing.T) {
	q := getTestQuery()

	q.Select([]string{})

	exp := fmt.Sprintf("SELECT * FROM %s.%s;", q.Database, q.Table)

	assert(t, exp, q.Query)
}

func TestSingleSelect(t *testing.T) {
	q := getTestQuery()

	q.Select([]string{"account_id"})

	exp := fmt.Sprintf("SELECT account_id FROM %s.%s;", q.Database, q.Table)

	assert(t, exp, q.Query)
}

func TestDoubleSelect(t *testing.T) {
	q := getTestQuery()

	sel := []string{
		"account_id",
		"value",
	}

	q.Select(sel)

	exp := fmt.Sprintf("SELECT account_id, value FROM %s.%s;", q.Database, q.Table)

	assert(t, exp, q.Query)
}

func TestSelectOne(t *testing.T) {
	q := getTestQuery()

	q.SelectOne()

	exp := fmt.Sprintf("SELECT 1 FROM %s.%s;", q.Database, q.Table)

	assert(t, exp, q.Query)
}

func TestSelectAll(t *testing.T) {
	q := getTestQuery()

	q.SelectAll()

	exp := fmt.Sprintf("SELECT * FROM %s.%s;", q.Database, q.Table)

	assert(t, exp, q.Query)
}

func TestEqualVarchar(t *testing.T) {
	col := Column{Name: "account_id"}
	res := col.Equal("'acc-001'")
	exp := "account_id = 'acc-001'"
	assert(t, exp, res)
}

func TestEqualNum(t *testing.T) {
	col := Column{Name: "account_id"}
	res := col.Equal("5")
	exp := "account_id = 5"
	assert(t, exp, res)
}

func TestNotEqual(t *testing.T) {
	col := Column{Name: "account_id"}
	res := col.NotEqual("'acc-001'")
	exp := "account_id <> 'acc-001'"
	assert(t, exp, res)
}

func TestWhereEquals(t *testing.T) {
	q := getTestQuery()

	q.SelectAll().Where(q.Columns["account_id"].Equal("'acc-001'"))

	exp := fmt.Sprintf("SELECT * FROM %s.%s WHERE account_id = 'acc-001';", q.Database, q.Table)

	assert(t, exp, q.Query)
}

func TestWhereLikeAndBetweenOrderByAsc(t *testing.T) {
	q := getTestQuery()
	like := `'acc-1%'`
	q.SelectAll().Where(q.Columns["account_id"].Like(like)).And(q.Columns["value"].Between(`500`, `1000`)).OrderByAsc(q.Columns["created_at"])

	exp := fmt.Sprintf("SELECT * FROM %s.%s WHERE account_id LIKE %s AND value BETWEEN 500 AND 1000 ORDER BY created_at ASC;", q.Database, q.Table, like)

	assert(t, exp, q.Query)
}

func TestWhereInOrNotEqualOrderByDesc(t *testing.T) {
	q := getTestQuery()

	q.SelectAll().Where(q.Columns["status"].In([]string{"failed", "deleted"})).Or(q.Columns["value"].NotEqual(`100`)).OrderByDesc(q.Columns["created_at"])

	exp := fmt.Sprintf("SELECT * FROM %s.%s WHERE status IN ('failed','deleted') OR value <> 100 ORDER BY created_at DESC;", q.Database, q.Table)

	assert(t, exp, q.Query)
}

func TestEmptyWhereAnd(t *testing.T) {
	q := getTestQuery()

	q.SelectAll().Where(q.Columns["status"].In([]string{})).And(q.Columns["value"].GreaterThan("100"))

	exp := fmt.Sprintf("SELECT * FROM %s.%s WHERE value > 100;", q.Database, q.Table)

	assert(t, exp, q.Query)
}

func TestEmptyWhereOr(t *testing.T) {
	q := getTestQuery()

	q.SelectAll().Where(q.Columns["status"].In([]string{})).Or(q.Columns["value"].GreaterThan("100"))

	exp := fmt.Sprintf("SELECT * FROM %s.%s WHERE value > 100;", q.Database, q.Table)

	assert(t, exp, q.Query)
}
