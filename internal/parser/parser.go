package parser

type ParseStage int

const (
	ConfigStage ParseStage = iota
	FilterStage
	ArgStage
)

type ArgParser struct{}

type ParseContext struct {
	Config  []Config
	Command Command
	Filters []Filter
	Args    []Arg
}

func join(text *string, arg string) {
	if *text != "" {
		*text += " "
	}

	*text += arg
}

func (p *ArgParser) Parse(args []string) ParseContext {
	ctx := ParseContext{
		Config:  []Config{},
		Command: "",
		Filters: []Filter{},
		Args:    []Arg{},
	}

	stage := ConfigStage
	text := ""
	ids := []int{}

	for _, arg := range args {
		if stage == ConfigStage {
			// If the argument is a config override (e.g., `bulk=3`), parse
			// it and add it to the config.
			if config, ok := ConfigFromString(arg); ok {
				ctx.Config = append(ctx.Config, config)
				continue
			}

			// If the argument is not a config override, move to the filter
			// stage and re-process the argument.
			stage = FilterStage
		}

		if stage == FilterStage {
			// If we haven't yet identified the command, start by attempting
			// to parse the text as the command. Once the command has been
			// identified, we would treat further text that matches a
			// command as simply a filter/arg (e.g., `tsk 12 edit improve edit command`).
			if ctx.Command == "" {
				if command, ok := commandFromStr(arg); ok {
					// If the command accepts args, move to the arg stage of
					// parsing and push any collected text into a filter. If a
					// command doesn't accept args, we will store the
					// identified command, but remain in the filter stage.
					if commandAcceptsArgs(command) {
						stage = ArgStage

						if text != "" {
							ctx.Filters = append(ctx.Filters, TextFilter{Text: text})
							text = ""
						}
					}

					ctx.Command = command
					continue
				}
			}

			// Try parsing the argument as a number and if it parses, add
			// it as an ID filter.
			if parsedIds, ok := parseIds(arg); ok {
				ids = append(ids, parsedIds...)
				continue
			}

			// If the argument starts with a + or - and has more than one
			// character, it's a tag filter.
			if operator, tag := parseTag(arg); tag != "" {
				ctx.Filters = append(ctx.Filters, TagFilter{Operator: operator, Tag: tag})
				continue
			}

			// If the argument starts with a scope (e.g. priority:) and it is
			// a valid scope, add it as a scope filter.
			if scope, value := parseScope(arg); scope != "" {
				ctx.Filters = append(ctx.Filters, ScopedFilter{Scope: scope, Value: value})
				continue
			}

			// If the argument is not a command, an ID, a tag, or a scope,
			// then it's a text filter.
			join(&text, arg)
			continue
		}

		if stage == ArgStage {
			// If the argument starts with a + or - and has more than one
			// character, it's a tag arg.
			if operator, tag := parseTag(arg); tag != "" {
				ctx.Args = append(ctx.Args, TagArg{Operator: operator, Tag: tag})
				continue
			}

			// If the argument starts with a scope (e.g. priority:) and it is
			// a valid scope, add it as a scope arg.
			if scope, value := parseScope(arg); scope != "" {
				ctx.Args = append(ctx.Args, ScopedArg{Scope: scope, Value: value})
				continue
			}

			// If the argument is not a command, an ID, a tag, or a scope,
			// then it's a text arg.
			join(&text, arg)
		}
	}

	// After finishing iterating over the arguments, we need to add the last
	// id and text arguments to the result.
	if len(ids) > 0 {
		ctx.Filters = append(ctx.Filters, IdFilter{Ids: ids})
	}

	if text != "" {
		switch stage {
		case ConfigStage, FilterStage:
			ctx.Filters = append(ctx.Filters, TextFilter{Text: text})
		case ArgStage:
			ctx.Args = append(ctx.Args, TextArg{Text: text})
		}
	}

	return ctx
}

func New() ArgParser {
	return ArgParser{}
}
