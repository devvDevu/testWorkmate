package task_model

import (
	"testWorkmate/internal/common/types/task_type"
	"time"
)

// Task is a type that represents a task.
type Task struct {
	ID        uint64
	Title     string
	CreatedAt time.Time
	RunTime   time.Duration
	Status    task_type.TaskStatus
}
