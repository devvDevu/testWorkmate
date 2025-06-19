package task_type

type TaskStatus string

func (t TaskStatus) String() string {
	return string(t)
}

const (
	TaskStatusPending    TaskStatus = "pending"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCompleted  TaskStatus = "completed"
)