package sql_builder

import "fmt"

type Builder struct {
	query     string
	usedWhere bool
	usedSet   bool
}

type Operator string

const (
	Eq      Operator = "="
	Neq     Operator = "!="
	In      Operator = "in"
	Like    Operator = "like"
	NotLike Operator = "not like"
)

type Filter struct {
	Key      string
	Operator Operator
	Value    string
}

func New() *Builder {
	return &Builder{}
}

func (b *Builder) Select(columns string) *Builder {
	b.query += "select " + columns
	return b
}

func (b *Builder) Update(table string) *Builder {
	b.query += "update " + table
	return b
}

func (b *Builder) Delete(table string) *Builder {
	b.query += "delete from " + table
	return b
}

func (b *Builder) From(table string) *Builder {
	b.query += " from " + table
	return b
}

func (b *Builder) Join(table string, condition string) *Builder {
	b.query += fmt.Sprintf(" join %s on %s", table, condition)
	return b
}

func (b *Builder) Filter(filter Filter) *Builder {
	if !b.usedWhere {
		b.query += " where "
		b.usedWhere = true
	} else {
		b.query += " and "
	}

	b.query += fmt.Sprintf("%s %s %s", filter.Key, filter.Operator, filter.Value)
	return b
}

func (b *Builder) Set(fields string) *Builder {
	if !b.usedSet {
		b.query += " set "
		b.usedSet = true
	} else {
		b.query += ", "
	}

	b.query += fields
	return b
}

func (b *Builder) SQL() string {
	return b.query
}
