package task_service

import (
	"context"
	"testWorkmate/internal/model/task_model"
)

// TaskService is a service for tasks.
// There we can add some actions with data and cache data(future).
type TaskService struct {
	taskRepository taskRepositoryI
}

type taskRepositoryI interface {
	CreateTask(ctx context.Context, task *task_model.Task) (*task_model.Task, error)
	GetTask(ctx context.Context, id uint64) (*task_model.Task, error)
	DeleteTask(ctx context.Context, id uint64) error
	UpdateTask(ctx context.Context, task *task_model.Task) (*task_model.Task, error)
}

func NewTaskService(taskRepository taskRepositoryI) *TaskService {
	return &TaskService{taskRepository: taskRepository}
}

func (s *TaskService) CreateTask(ctx context.Context, task *task_model.Task) (*task_model.Task, error) {
	return s.taskRepository.CreateTask(ctx, task)
}

func (s *TaskService) GetTask(ctx context.Context, id uint64) (*task_model.Task, error) {
	return s.taskRepository.GetTask(ctx, id)
}

func (s *TaskService) DeleteTask(ctx context.Context, id uint64) error {
	return s.taskRepository.DeleteTask(ctx, id)
}

func (s *TaskService) UpdateTask(ctx context.Context, task *task_model.Task) (*task_model.Task, error) {
	return s.taskRepository.UpdateTask(ctx, task)
}
