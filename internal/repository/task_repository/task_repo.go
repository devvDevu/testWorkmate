package task_repository

import (
	"context"
	"testWorkmate/internal/model/task_model"
)

// TaskRepository is a repository for tasks.
// There we can add grouping sql queries(future).
type TaskRepository struct {
	dbAdapter dbAdapterI
}

// dbAdapterI is an interface that defines the methods for the database adapter(slick to clean archeticture).
type dbAdapterI interface {
	Exec(ctx context.Context, values *task_model.Task, dest *task_model.Task) (error)
	Get(ctx context.Context,dest *task_model.Task, id uint64) (error)
	Delete(ctx context.Context, id uint64) (error)
}

func NewTaskRepository(dbAdapter dbAdapterI) *TaskRepository {
	return &TaskRepository{
		dbAdapter: dbAdapter,
	}
}

func (r *TaskRepository) CreateTask(ctx context.Context, task *task_model.Task) (*task_model.Task, error) {
	dest := new(task_model.Task)

	err := r.dbAdapter.Exec(ctx, task, dest)

	return dest, err
}

func (r *TaskRepository) GetTask(ctx context.Context, id uint64) (*task_model.Task, error) {
	dest := new(task_model.Task)

	err := r.dbAdapter.Get(ctx, dest, id)

	return dest, err
}

func (r *TaskRepository) DeleteTask(ctx context.Context, id uint64) error {
	err := r.dbAdapter.Delete(ctx, id)

	return err
}

func (r *TaskRepository) UpdateTask(ctx context.Context, task *task_model.Task) (*task_model.Task, error) {
	dest := new(task_model.Task)

	err := r.dbAdapter.Exec(ctx, task, dest)

	return dest, err
}