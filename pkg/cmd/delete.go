package cmd

import (
	"errors"
	"fmt"

	"github.com/mskelton/tsk/internal/arg_parser"
	"github.com/mskelton/tsk/internal/printer"
	"github.com/mskelton/tsk/internal/storage"
)

func Delete(ctx arg_parser.ParseContext) {
	requireFilters(ctx, "delete")

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

	if count != 1 {
		printer.Error(errors.New("Bulk delete is not supported"))
		return
	}

	if !printer.Confirm("Are you sure you want to continue?") {
		return
	}

	ids, err := storage.Delete(filters)
	if err != nil {
		printer.Error(err)
		return
	}

	for _, id := range ids {
		fmt.Printf("Deleted task %d\n", id)
	}
}
