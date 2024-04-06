package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mskelton/tsk/internal/arg_parser"
	"github.com/mskelton/tsk/internal/printer"
	"github.com/mskelton/tsk/internal/sql_builder"
	"github.com/mskelton/tsk/internal/utils"
)

func requireFilters(ctx arg_parser.ParseContext, command string) {
	if len(ctx.Filters) == 0 {
		printer.Error(utils.CLIError{
			Message: fmt.Sprintf("The %s command requires filters", command),
		})
	}
}

func buildFilters(ctx arg_parser.ParseContext) []sql_builder.Filter {
	var filters []sql_builder.Filter

	for _, f := range ctx.Filters {
		switch filter := f.(type) {
		case arg_parser.IdFilter:
			var ids []string
			for _, id := range filter.Ids {
				ids = append(ids, strconv.Itoa(id))
			}

			filters = append(filters, sql_builder.Filter{
				Key:      "tasks.id",
				Operator: sql_builder.In,
				Value: fmt.Sprintf(
					"(select task_id from assignments where id in (%s))",
					strings.Join(ids, ", "),
				),
			})

		case arg_parser.TextFilter:
			filters = append(filters, sql_builder.Filter{
				Key:      "data ->> 'title'",
				Operator: sql_builder.Like,
				Value:    fmt.Sprintf("'%%%s%%'", filter.Text),
			})

		case arg_parser.TagFilter:
			operator := sql_builder.Like
			if filter.Operator == arg_parser.Exclude {
				operator = sql_builder.NotLike
			}

			filters = append(filters, sql_builder.Filter{
				Key:      "JSON_EXTRACT(data, '$.tags')",
				Operator: operator,
				Value:    fmt.Sprintf("'%%\"%s\"%%'", filter.Tag),
			})

		case arg_parser.ScopedFilter:
			filters = append(filters, sql_builder.Filter{
				Key:      fmt.Sprintf("data ->> '%s'", filter.Scope),
				Operator: sql_builder.Eq,
				Value:    fmt.Sprintf("'%s'", filter.Value),
			})
		}
	}

	return filters
}
