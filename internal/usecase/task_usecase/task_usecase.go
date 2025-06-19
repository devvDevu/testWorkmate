package task_usecase

import (
	"context"
	"errors"
	"sync"
	"testWorkmate/internal/model/task_model"

	"github.com/sirupsen/logrus"
)

const usecaseName = "TaskUsecase"

type TaskUsecase struct {
	taskService taskServiceI
}

type taskServiceI interface {
	CreateTask(ctx context.Context, task *task_model.Task) (*task_model.Task, error)
	GetTask(ctx context.Context, id uint64) (*task_model.Task, error)
	DeleteTask(ctx context.Context, id uint64) error
	UpdateTask(ctx context.Context, task *task_model.Task) (*task_model.Task, error)
}

func NewTaskUsecase(taskService taskServiceI) *TaskUsecase {
	return &TaskUsecase{taskService: taskService}
}

func (t *TaskUsecase) CreateTask(ctx context.Context, task *task_model.Task) (*task_model.Task, error) {
	const action = "TaskUsecase CreateTask"
	const method = "CreateTask"

	wg := sync.WaitGroup{}
	var taskResult chan *task_model.Task

	wg.Add(1)
	go func() {
		defer wg.Done()
		task, err := t.GenerateTask(ctx, task)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"usecaseName": usecaseName,
				"method":      method,
			}).WithError(err).Error(action)
		}
		taskResult <- task
	}()

	wg.Wait()

	val, ok := <-taskResult
	if !ok {
		return nil, errors.New("task result channel closed")
	}

	task = val

	logrus.WithFields(logrus.Fields{
		"usecaseName": usecaseName,
		"method":      method,
	}).Info("Task successfully created ", action)

	return task, nil
}

func (t *TaskUsecase) GetTask(ctx context.Context, id uint64) (*task_model.Task, error) {
	const action = "TaskUsecase GetTask"
	const method = "GetTask"

	logrus.WithFields(logrus.Fields{
		"usecaseName": usecaseName,
		"method":      method,
	}).Info("Task successfully got ", action)

	return t.taskService.GetTask(ctx, id)
}

func (t *TaskUsecase) DeleteTask(ctx context.Context, id uint64) error {
	const action = "TaskUsecase DeleteTask"
	const method = "DeleteTask"

	logrus.WithFields(logrus.Fields{
		"usecaseName": usecaseName,
		"method":      method,
	}).Info("Task successfully deleted ", action)

	return t.taskService.DeleteTask(ctx, id)
}
