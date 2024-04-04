package cmd

import (
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/mskelton/tsk/internal/arg_parser"
	"github.com/mskelton/tsk/internal/printer"
	"github.com/mskelton/tsk/internal/storage"
	"github.com/mskelton/tsk/internal/utils"
)

func List(ctx arg_parser.ParseContext) {
	filters := buildFilters(ctx)
	tasks, err := storage.ListTasks(filters)
	if err.Code != "" {
		printer.Error(err)
	}

	if len(tasks) == 0 {
		printer.Message("No tasks match filters")
		return
	}

	table := printer.Table{
		Columns: []string{"ID", "Active", "Age", "P", "Tags", "Title"},
		Rows:    []printer.Row{},
	}

	for _, task := range tasks {
		var status string
		if task.Status == storage.TaskStatusActive && color.NoColor {
			status = "✔︎"
		}

		table.Rows = append(table.Rows, printer.Row{
			Cells: []string{
				strconv.Itoa(task.ShortId),
				status,
				utils.ShortDuration(task.CreatedAt),
				task.Priority,
				strings.Join(task.Tags, " "),
				task.Title,
			},
			Highlight: task.Status == storage.TaskStatusActive,
		})
	}

	table.Print()
}
