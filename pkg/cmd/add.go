package cmd

import (
	"errors"
	"fmt"

	"github.com/mskelton/tsk/internal/arg_parser"
	"github.com/mskelton/tsk/internal/printer"
	"github.com/mskelton/tsk/internal/storage"
)

func Add(context arg_parser.ParseContext) {
	task := storage.NewTask()

	for _, arg := range context.Args {
		if text, ok := arg.(arg_parser.TextArg); ok {
			task.Title = text.Text
		} else if tag, ok := arg.(arg_parser.TagArg); ok {
			task.Tags = append(task.Tags, tag.Tag)
		} else if scoped, ok := arg.(arg_parser.ScopedArg); ok {
			if scoped.Scope == arg_parser.ScopePriority {
				task.Priority = scoped.Value
			} else {
				printer.Error(fmt.Sprintf("Missing value for \"%s:\"", scoped.Scope), errors.New("invalid args"))
			}
		}
	}

	if task.Title == "" {
		printer.Error("Missing title", errors.New("invalid args"))
	}

	id, err := storage.Add(task)
	if err != nil {
		printer.Error("Failed to add task", err)
	}

	fmt.Println("Created task", id)
}
