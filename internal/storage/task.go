package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mskelton/tsk/internal/sql_builder"
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

func ListTasks(filters []sql_builder.Filter) ([]Task, error) {
	conn, err := connect()
	if err != nil {
		return nil, fmt.Errorf("Failed to list tasks: %w", err)
	}

	builder := sql_builder.New().
		Select("tasks.id, tasks.template_id, assignments.id, tasks.data").
		From("tasks").
		Join("assignments", "tasks.id = assignments.task_id").
		Filter(sql_builder.Filter{
			Key:      "tasks.data ->> '$.status'",
			Operator: sql_builder.Neq,
			Value:    "'done'",
		})

	for _, filter := range filters {
		builder.Filter(filter)
	}

	if os.Getenv("DEBUG") != "" {
		log.Println(builder.SQL())
	}

	rows, err := conn.Query(builder.SQL())
	if err != nil {
		return nil, fmt.Errorf("Failed to list tasks: %w", err)
	}

	var tasks []Task

	for rows.Next() {
		var taskId string
		var templateId sql.NullString
		var shortId sql.NullInt64
		var data []byte

		err = rows.Scan(&taskId, &templateId, &shortId, &data)
		if err != nil {
			return nil, fmt.Errorf("Failed to list tasks: %w", err)
		}

		var task Task
		err = json.Unmarshal(data, &task)
		if err != nil {
			return nil, fmt.Errorf("Failed to list tasks: %w", err)
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
		return 0, fmt.Errorf("Failed to add task: %w", err)
	}

	conn, err := connect()
	if err != nil {
		return 0, fmt.Errorf("Failed to add task: %w", err)
	}

	_, err = conn.Exec(
		"INSERT INTO tasks (id, template_id, data) VALUES (?, ?, ?)",
		task.Id,
		task.TemplateId,
		data,
	)
	if err != nil {
		return 0, fmt.Errorf("Failed to add task: %w", err)
	}

	// Add an id assignment for the newly created task
	res, err := conn.Exec(
		"INSERT INTO assignments VALUES ((select max(id) + 1 from assignments), ?)",
		task.Id,
	)
	if err != nil {
		return 0, fmt.Errorf("Failed to add task assignment: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Failed to get last insert id: %w", err)
	}

	// Return the id of the newly created task. Thankfully SQLite handles this
	// automatically with `LastInsertId()` since we are using a numeric id.
	return id, nil
}

func Count(filters []sql_builder.Filter) (int, error) {
	conn, err := connect()
	if err != nil {
		return 0, fmt.Errorf("Failed to count tasks: %w", err)
	}

	builder := sql_builder.New().
		Select("count(tasks.id)").
		From("tasks").
		Join("assignments", "tasks.id = assignments.task_id")

	for _, filter := range filters {
		builder.Filter(filter)
	}

	debug := os.Getenv("DEBUG") != ""
	if debug {
		log.Println(builder.SQL())
	}

	row := conn.QueryRow(builder.SQL())
	if row.Err() != nil {
		return 0, fmt.Errorf("Failed to count tasks: %w", row.Err())
	}

	var count int
	err = row.Scan(&count)
	if row.Err() != nil {
		return 0, fmt.Errorf("Failed to count tasks: %w", row.Err())
	}

	return count, nil
}

func getIds(conn *sql.DB, filters []sql_builder.Filter) ([]int, error) {
	builder := sql_builder.New().
		Select("assignments.id").
		From("tasks").
		Join("assignments", "tasks.id = assignments.task_id")

	for _, filter := range filters {
		builder.Filter(filter)
	}

	debug := os.Getenv("DEBUG") != ""
	if debug {
		log.Println(builder.SQL())
	}

	res, err := conn.Query(builder.SQL())
	if err != nil {
		return nil, fmt.Errorf("Failed to get task ids: %w", err)
	}

	var ids []int
	for res.Next() {
		var id int
		err = res.Scan(&id)

		if err != nil {
			return nil, fmt.Errorf("Failed to get task id: %w", err)
		}

		ids = append(ids, id)
	}

	return ids, nil
}

type QueryEdit struct {
	Path  string
	Value string
}

func Edit(filters []sql_builder.Filter, edits []QueryEdit) ([]int, error) {
	conn, err := connect()
	if err != nil {
		return nil, fmt.Errorf("Failed to edit tasks: %w", err)
	}

	builder := sql_builder.New().Update("tasks")
	var params []any

	for _, edit := range edits {
		params = append(params, edit.Value)
		builder.Set(fmt.Sprintf("data = json_set(data, '$.%s', ?)", edit.Path))
	}

	for _, filter := range filters {
		builder.Filter(filter)
	}

	debug := os.Getenv("DEBUG") != ""
	if debug {
		log.Println(builder.SQL())
	}

	_, err = conn.Exec(builder.SQL(), params...)
	if err != nil {
		return nil, fmt.Errorf("Failed to edit tasks: %w", err)
	}

	return getIds(conn, filters)
}

func Delete(filters []sql_builder.Filter) ([]int, error) {
	conn, err := connect()
	if err != nil {
		return nil, fmt.Errorf("Failed to delete tasks: %w", err)
	}

	builder := sql_builder.New().Delete("tasks")

	for _, filter := range filters {
		builder.Filter(filter)
	}

	debug := os.Getenv("DEBUG") != ""
	if debug {
		log.Println(builder.SQL())
	}

	_, err = conn.Exec(builder.SQL())
	if err != nil {
		return nil, fmt.Errorf("Failed to delete tasks: %w", err)
	}

	return getIds(conn, filters)
}
