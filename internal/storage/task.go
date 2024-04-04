package storage

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/mskelton/tsk/internal/utils"
)

type TaskStatus string

const (
	TaskStatusPending TaskStatus = "pending"
	TaskStatusActive  TaskStatus = "active"
	TaskStatusDone    TaskStatus = "done"
)

type Task struct {
	// The unique identifier for the task
	Id string
	// A short numerical identifier for the task, used for quick reference in the UI.
	ShortId int
	// The parent recurrence template the task was created from (if any). This
	// is used when finding other tasks from the same recurrence template or
	// when modifying the recurrence options.
	TemplateId string
	// The title of the task
	Title string `json:"title"`
	// The priority of the task, typically something like `H`, `M`, or `L`,
	// though the values are user-defined.
	Priority string `json:"priority"`
	// The status of the task, one of `pending`, `active`, or `done`. Tasks
	// start as `pending`, and can move between `active`, `pending`, and `done`
	// as the user sees fit. Typically a task does not move from done to the
	// other statuses, but it is not enforced.
	Status TaskStatus `json:"status"`
	// A list of tags for the task. Tags are useful for grouping tasks together
	// and can be used to filter tasks in the UI.
	Tags []string `json:"tags"`
	// The time the task was created
	CreatedAt time.Time `json:"created_at"`
	// The time the task was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTask() Task {
	return Task{
		Id:        utils.GenerateId(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    TaskStatusPending,
		Tags:      make([]string, 0),
	}
}

func ListTasks() ([]Task, error) {
	conn, err := connect()
	if err != nil {
		return nil, err
	}

	// let mut sql_builder = Builder::new();
	// sql_builder
	//     .select("tasks.id, tasks.template_id, assignments.id, tasks.data")
	//     .from("tasks")
	//     .join("assignments", "tasks.id = assignments.task_id")
	//     .filter("tasks.data ->> '$.status' != 'done'");
	//
	// filters.iter().for_each(|f| {
	//     sql_builder.filter(&f.to_sql());
	// });
	//
	// let sql = sql_builder.sql();
	// debug!("{}", &sql);

	query := "SELECT tasks.id, tasks.template_id, assignments.id, tasks.data FROM tasks JOIN assignments ON tasks.id = assignments.task_id WHERE tasks.data ->> '$.status' != 'done'"

	stmt, err := conn.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var tasks []Task

	for rows.Next() {
		var taskId string
		var templateId sql.NullString
		var shortId sql.NullInt64
		var data []byte

		err = rows.Scan(&taskId, &templateId, &shortId, &data)
		if err != nil {
			return nil, err
		}

		var task Task
		err = json.Unmarshal(data, &task)
		if err != nil {
			return nil, err
		}

		task.Id = taskId

		if shortId.Valid {
			task.ShortId = int(shortId.Int64)
		}

		if templateId.Valid {
			task.TemplateId = templateId.String
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func Add(task Task) (int64, error) {
	data, err := json.Marshal(task)
	if err != nil {
		return 0, err
	}

	conn, err := connect()
	if err != nil {
		return 0, err
	}

	conn.Exec(
		"INSERT INTO tasks (id, template_id, data) VALUES (?, ?, ?)",
		task.Id,
		task.TemplateId,
		data,
	)

	// Add an id assignment for the newly created task
	conn.Exec(
		"INSERT INTO assignments VALUES ((select max(id) + 1 from assignments), ?)",
		task.Id,
	)

	// Return the id of the newly created task. Thankfully SQLite handles this
	// automatically with `last_insert_rowid()` since we are using a numeric id.
	// return conn.last_insert_rowid(), nil
	return 0, nil
}
