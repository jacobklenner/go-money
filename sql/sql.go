package sql

import (
	"fmt"
	"strings"
)

// dont think this should be called 'query' as it contains more information than just a simple query
type Query struct {
	// Language string // should support syntactic differences in mysql, postgres, ms sql etc
	// Version string // may be differences in symbols based on versions
	Database string
	Table    string
	Columns  Columns // TODO add support for table meta data, and can do some validation on the built query
	Query    string
}

type Columns map[string]Column

type Column struct {
	Name string
	// DataType string // TODO could add some data type validation
}

// name these based on what reads well when constructing? or
func (c Column) Equal(v string) string {
	return fmt.Sprintf("%s = %s", c.Name, v)
}

func (c Column) NotEqual(v string) string {
	return fmt.Sprintf("%s <> %s", c.Name, v)
}

func (c Column) GreaterThan(v string) string {
	return fmt.Sprintf("%s > %s", c.Name, v)
}

func (c Column) GreaterThanOrEqual(v string) string {
	return fmt.Sprintf("%s >= %s", c.Name, v)
}

func (c Column) LessThan(v string) string {
	return fmt.Sprintf("%s < %s", c.Name, v)
}

func (c Column) LessThanOrEqual(v string) string {
	return fmt.Sprintf("%s <= %s", c.Name, v)
}

func (c Column) Like(p string) string {
	return fmt.Sprintf("%s LIKE %s", c.Name, p)
}

func (c Column) NotLike(p string) string {
	return fmt.Sprintf("%s NOT LIKE %s", c.Name, p)
}

func (c Column) Between(l string, u string) string {
	return fmt.Sprintf("%s BETWEEN %s AND %s", c.Name, l, u)
}

func (c Column) In(r string) string {
	return fmt.Sprintf("%s IN (%s)", c.Name, r)
}

func (c Column) NotIn(r string) string {
	return fmt.Sprintf("%s NOT IN (%s)", c.Name, r)
}

func (q *Query) Select(cols []string) *Query {
	var s string

	// do we expect this as default behaviour?
	if len(cols) == 0 {
		s = "*"
	} else {
		for _, c := range cols {
			s = s + fmt.Sprintf("%s, ", c)
		}
		s = strings.TrimSuffix(s, ", ")
	}

	q.Query = fmt.Sprintf("SELECT %s FROM %s.%s;", s, q.Database, q.Table)
	return q
}

func (q *Query) SelectOne() *Query {
	q.Query = fmt.Sprintf("SELECT 1 FROM %s.%s;", q.Database, q.Table)
	return q
}

func (q *Query) SelectAll() *Query {
	q.Query = fmt.Sprintf("SELECT * FROM %s.%s;", q.Database, q.Table)
	return q
}

func (q *Query) Where(cond string) *Query {
	q.Query = strings.TrimSuffix(q.Query, ";")
	q.Query = fmt.Sprintf("%s WHERE %s;", q.Query, cond)

	return q
}

func (q *Query) And(cond string) *Query {
	q.Query = strings.TrimSuffix(q.Query, ";")
	q.Query = fmt.Sprintf("%s AND %s;", q.Query, cond)

	return q
}

func (q *Query) Or(cond string) *Query {
	q.Query = strings.TrimSuffix(q.Query, ";")
	q.Query = fmt.Sprintf("%s OR %s;", q.Query, cond)

	return q
}

func (q *Query) Limit(val int) *Query {
	q.Query = strings.TrimSuffix(q.Query, ";")
	q.Query = fmt.Sprintf("%s LIMIT(%d);", q.Query, val)

	return q
}

func (q *Query) OrderByAsc(col Column) *Query {
	q.Query = strings.TrimSuffix(q.Query, ";")
	q.Query = fmt.Sprintf("%s ORDER BY %s ASC;", q.Query, col.Name)

	return q
}

func (q *Query) OrderByDesc(col Column) *Query {
	q.Query = strings.TrimSuffix(q.Query, ";")
	q.Query = fmt.Sprintf("%s ORDER BY %s DESC;", q.Query, col.Name)
	return q
}
