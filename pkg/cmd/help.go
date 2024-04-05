package cmd

import (
	"fmt"
	"strings"
)

const HELP = `
Usage: tsk <filters> <command> <args>

Commands:
  list          Show the task list
  add           Add a new task
  done          Mark a task as done
  edit          Edit a task
  show          Show a task
  start         Start a task
  stop          Stop a task
  get           Get a task
  delete        Delete a task
  help          Show this help message
  version       Show the version

For more information and examples, see https://tsk.mskelton.dev for the full
documentation or run ` + "`" + "tsk help <command>" + "`" + ` for help with a specific command.
`

func Help() {
	fmt.Println(strings.TrimSpace(HELP))
}
