package sql_builder_test

import (
	"testing"

	"github.com/mskelton/tsk/internal/sql_builder"
	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	assert.Equal(t, sql_builder.New().SQL(), "")
}

func TestSelect(t *testing.T) {
	sql := sql_builder.New().Select("id, name").From("users").SQL()
	assert.Equal(t, sql, "select id, name from users")
}

func TestUpdate(t *testing.T) {
	sql := sql_builder.New().Update("users").Set("name = 'foo'").SQL()
	assert.Equal(t, sql, "update users set name = 'foo'")
}

func TestJoins(t *testing.T) {
	sql := sql_builder.New().
		Select("id, name").
		From("users").
		Join("roles", "users.role_id = roles.id").
		SQL()

	assert.Equal(t, sql, "select id, name from users join roles on users.role_id = roles.id")
}

func TestFilterSingleCondition(t *testing.T) {
	sql := sql_builder.New().
		Select("id, name").
		From("users").
		Filter(sql_builder.Filter{
			Key:      "id",
			Operator: sql_builder.Eq,
			Value:    "1",
		}).
		SQL()

	assert.Equal(t, sql, "select id, name from users where id = 1")
}

func TestFilterMultipleConditions(t *testing.T) {
	sql := sql_builder.New().
		Select("id, name").
		From("users").
		Filter(sql_builder.Filter{
			Key:      "id",
			Operator: sql_builder.Eq,
			Value:    "1",
		}).
		Filter(sql_builder.Filter{
			Key:      "name",
			Operator: sql_builder.Like,
			Value:    "'John'",
		}).
		SQL()

	assert.Equal(t, sql, "select id, name from users where id = 1 and name like 'John'")
}
