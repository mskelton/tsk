package cmd

import (
	"errors"
	"fmt"

	"github.com/mskelton/tsk/internal/arg_parser"
	"github.com/mskelton/tsk/internal/printer"
	"github.com/mskelton/tsk/internal/storage"
	"github.com/mskelton/tsk/internal/utils"
)

func Done(ctx arg_parser.ParseContext) {
	requireFilters(ctx, "done")

	filters := buildFilters(ctx)
	count, err := storage.Count(filters)
	if err != nil {
		printer.Error(err)
		return
	}

	if count == 0 {
		printer.Error(errors.New("No tasks match filters"))
		return
	}

	fmt.Printf(
		"This command will complete %d %s\n",
		count,
		utils.Pluralize(count, "task", "tasks"),
	)

	if utils.IsBulk(ctx, count) && !printer.Confirm("Are you sure you want to continue?") {
		return
	}

	edits := []storage.QueryEdit{{
		Path:  "status",
		Value: string(storage.TaskStatusDone),
	}}

	ids, err := storage.Edit(filters, edits)
	if err != nil {
		printer.Error(err)
		return
	}

	for _, id := range ids {
		fmt.Printf("Completed task %d\n", id)
	}
}
