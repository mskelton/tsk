package main

import (
	"os"

	"github.com/mskelton/tsk/internal/arg_parser"
	"github.com/mskelton/tsk/pkg/cmd"
)

func main() {
	args := os.Args[1:]
	parser := arg_parser.New()
	context := parser.Parse(args)

	switch context.Command {
	case arg_parser.List:
		cmd.List(context)
	case arg_parser.Add:
		cmd.Add(context)
	// case arg_parser.Done:
	// 	cmd.Done(context)
	// case arg_parser.Edit:
	// 	cmd.Edit(context)
	case arg_parser.Show:
		cmd.Show(context)
	// case arg_parser.Start:
	// 	cmd.Start(context)
	// case arg_parser.Stop:
	// 	cmd.Stop(context)
	// case arg_parser.Get:
	// 	cmd.Get(context)
	// case arg_parser.Delete:
	// 	cmd.Delete(context)
	// case arg_parser.Help:
	// 	cmd.Help()
	// case arg_parser.Version:
	// 	cmd.Version()
	default:
		var ids []int

		for _, filter := range context.Filters {
			if filter, ok := filter.(arg_parser.IdFilter); ok {
				ids = append(ids, filter.Ids...)
			}
		}

		// If there is only one id, show the task, otherwise list all tasks
		// that match the filters.
		if len(ids) == 1 {
			cmd.Show(context)
		} else {
			cmd.List(context)
		}
	}

}
