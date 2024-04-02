# Command Syntax

Commands in tsk are broken into three parts:

- Filters: Commands such as `list` or `modify` use filters to identify the
  task(s) to operate on.
- Command: The command represents the action to perform **and** separates
  filters from args.
- Args: Commands such as `add` or `modify` require additional arguments to
  perform their action.

These three parts can be visualized using the example below:

```bash
tsk <filters> <command> <args>
```

Whether a part is required is command dependent. For example, when adding a new
task, you never need to provide filters. When marking a task as done, you do not
need to provide additional args. Refer to the sections in the docs for each
command for usage examples.
