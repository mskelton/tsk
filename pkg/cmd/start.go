package cmd

import (
	"fmt"

	"github.com/mskelton/tsk/internal/arg_parser"
	"github.com/mskelton/tsk/internal/printer"
	"github.com/mskelton/tsk/internal/storage"
	"github.com/mskelton/tsk/internal/utils"
)

func Start(ctx arg_parser.ParseContext) {
	requireFilters(ctx, "start")

	filters := buildFilters(ctx)
	count, err := storage.Count(filters)
	if err.Message != "" {
		printer.Error(err)
		return
	}

	if count == 0 {
		printer.Error(utils.CLIError{
			Message: "No tasks match filters",
		})
		return
	}

	fmt.Printf(
		"This command will start %d %s\n",
		count,
		utils.Pluralize(count, "task", "tasks"),
	)

	if utils.IsBulk(ctx, count) && !printer.Confirm("Are you sure you want to continue?") {
		return
	}

	edits := []storage.QueryEdit{{
		Path:  "status",
		Value: string(storage.TaskStatusActive),
	}}

	ids, err := storage.Edit(filters, edits)
	if err.Message != "" {
		printer.Error(err)
		return
	}

	for _, id := range ids {
		fmt.Printf("Started task %d\n", id)
	}
}
