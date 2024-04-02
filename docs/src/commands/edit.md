# edit

Edits an existing task.

```bash
tsk 12 edit Buy oranges +shopping
```

When editing a task, only what you specify will change. For example, the
following command will remove the "work" tag and set the priority to low.
Because we did not specify a new title, the title will remain unchanged.

```bash
tsk 12 edit -work priority:L
```

## Bulk Editing

In addition to editing a single task by id, you can specify filters to modify
multiple tasks. For example, to mark all your tasks with the tag `repair` as
high priority, you could use the following command:

```bash
tsk +repair edit priority:H
```

Refer to the [filters](../filters.md) page for more details about the available
filters and how to use them effectively.
