package utils_test

import (
	"testing"

	"github.com/mskelton/tsk/internal/arg_parser"
	"github.com/mskelton/tsk/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestPluralize(t *testing.T) {
	assert.Equal(t, utils.Pluralize(0, "task", "tasks"), "tasks")
	assert.Equal(t, utils.Pluralize(1, "task", "tasks"), "task")
	assert.Equal(t, utils.Pluralize(2, "task", "tasks"), "tasks")
}

func TestBulk(t *testing.T) {
	context := arg_parser.ParseContext{}

	assert.Equal(t, utils.IsBulk(context, 0), false)
	assert.Equal(t, utils.IsBulk(context, 1), false)
	assert.Equal(t, utils.IsBulk(context, 3), false)
	assert.Equal(t, utils.IsBulk(context, 4), true)
	assert.Equal(t, utils.IsBulk(context, 10), true)
}
