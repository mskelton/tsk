package cmd

import (
	"fmt"

	"github.com/mskelton/tsk/internal/arg_parser"
	"github.com/mskelton/tsk/internal/printer"
	"github.com/mskelton/tsk/internal/utils"
)

func requireFilters(context arg_parser.ParseContext, command string) {
	if len(context.Filters) == 0 {
		printer.Error(utils.CLIError{
			Code:    utils.ErrorCodeInvalidArgs,
			Message: fmt.Sprintf("The %s command requires filters", command),
		})
	}
}
