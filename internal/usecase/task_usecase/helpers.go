package task_usecase

import (
	"context"
	"sync/atomic"
	"testWorkmate/internal/common/types/task_type"
	"testWorkmate/internal/model/task_model"
	"time"
)

var taskIDCounter atomic.Uint64

func (t *TaskUsecase) GenerateTask(ctx context.Context, task *task_model.Task) (*task_model.Task, error) {
	timeStart := time.Now()
	task.ID = taskIDCounter.Add(1)
	task.CreatedAt = time.Now()

	task.Status = task_type.TaskStatusPending
	task, err := t.taskService.CreateTask(ctx, task)
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Second * 120) // 120 seconds
	task.Status = task_type.TaskStatusInProgress
	task, err = t.taskService.UpdateTask(ctx, task)
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Second * 60) // 60 seconds
	task.Status = task_type.TaskStatusCompleted
	task, err = t.taskService.UpdateTask(ctx, task)
	if err != nil {
		return nil, err
	}
	task.RunTime = time.Since(timeStart)
	return task, nil
}
