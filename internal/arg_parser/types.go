package arg_parser

import (
	"strconv"
	"strings"
)

type Operator string

const (
	Include Operator = "Include"
	Exclude Operator = "Exclude"
)

type Scope string

const (
	ScopePriority Scope = "priority"
)

type Command string

const (
	List    Command = "list"
	Add     Command = "add"
	Done    Command = "done"
	Edit    Command = "edit"
	Show    Command = "show"
	Start   Command = "start"
	Stop    Command = "stop"
	Get     Command = "get"
	Delete  Command = "delete"
	Help    Command = "help"
	Version Command = "version"
)

type Filter interface{}

type IdFilter struct {
	Ids []int
}

type TagFilter struct {
	Operator Operator
	Tag      string
}

type ScopedFilter struct {
	Scope Scope
	Value string
}

type TextFilter struct {
	Text string
}

type Arg interface{}

type TagArg struct {
	Operator Operator
	Tag      string
}

type ScopedArg struct {
	Scope Scope
	Value string
}

type TextArg struct {
	Text string
}

type Config interface{}

type BulkConfig struct {
	Size int
}

type ContextConfig struct {
	Context string
}

func commandFromStr(str string) (Command, bool) {
	switch Command(str) {
	case List, Add, Done, Edit, Show, Start, Stop, Get, Delete, Help, Version:
		return Command(str), true
	default:
		return "", false
	}
}

func commandAcceptsArgs(command Command) bool {
	switch command {
	case Add, Edit, Get:
		return true
	default:
		return false
	}
}

func ConfigFromStr(str string) (Config, bool) {
	parts := strings.Split(str, "=")

	// Scoped arguments are only valid if they have exactly 2 parts. This
	// includes the case where the second part is empty.
	if len(parts) != 2 {
		return nil, false
	}

	// Attempt to parse each configuration option.
	switch parts[0] {
	case "bulk":
		if size, err := strconv.Atoi(parts[1]); err == nil {
			return BulkConfig{Size: size}, true
		} else {
			return nil, false
		}

	default:
		return nil, false
	}
}
