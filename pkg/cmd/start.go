package cmd

import "github.com/mskelton/tsk/internal/arg_parser"

func Start(context arg_parser.ParseContext) {
	requireFilters(context, "start")
}
