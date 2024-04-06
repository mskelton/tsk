package cmd

import (
	"fmt"

	"github.com/mskelton/tsk/internal/arg_parser"
	"github.com/mskelton/tsk/internal/printer"
	"github.com/mskelton/tsk/internal/storage"
	"github.com/mskelton/tsk/internal/utils"
)

func Add(ctx arg_parser.ParseContext) {
	task := storage.NewTask()

	for _, arg := range ctx.Args {
		switch v := arg.(type) {
		case arg_parser.TextArg:
			task.Title = v.Text
		case arg_parser.TagArg:
			task.Tags = append(task.Tags, v.Tag)
		case arg_parser.ScopedArg:
			if v.Scope == arg_parser.ScopePriority {
				task.Priority = v.Value
			} else {
				printer.Error(utils.CLIError{
					Message: fmt.Sprintf("Missing value for \"%s:\"", v.Scope),
				})
			}
		}
	}

	if task.Title == "" {
		printer.Error(utils.CLIError{
			Message: "Missing title",
		})
	}

	id, err := storage.Add(task)
	if err.Message != "" {
		printer.Error(err)
	}

	fmt.Println("Created task", id)
}
