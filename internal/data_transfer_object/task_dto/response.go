package task_dto

import (
	"testWorkmate/internal/common/types/task_type"
	"time"
)

type TaskResponse struct {
	ID        uint64               `json:"id"`
	Title     string               `json:"title"`
	Status    task_type.TaskStatus `json:"status"`
	CreatedAt time.Time            `json:"created_at"`
	RunTime   time.Duration        `json:"run_time"`
}
