package create_task_handler

import (
	"context"
	"io"
	"net/http"
	"testWorkmate/internal/common/types/error_with_codes"
	"testWorkmate/internal/common/types/handler_type"
	"testWorkmate/internal/data_transfer_object/result"
	"testWorkmate/internal/data_transfer_object/task_dto"
	"testWorkmate/internal/model/task_model"
	"time"

	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
)

const handlerName = "CreateTaskHandler"

type CreateTaskHandler struct {
	taskUsecase taskUsecaseI
}

type taskUsecaseI interface {
	CreateTask(ctx context.Context, task *task_model.Task) (*task_model.Task, error)
}

func NewCreateTaskHandler(taskUsecase taskUsecaseI) *CreateTaskHandler {
	return &CreateTaskHandler{
		taskUsecase: taskUsecase,
	}
}

func (h *CreateTaskHandler) GetPath() handler_type.HandlerPath {
	return "/api/v1/task"
}

func (h *CreateTaskHandler) GetMethod() handler_type.HandlerMethod {
	return handler_type.HandlerMethodPost
}

func (h *CreateTaskHandler) ExecFunc(ctx context.Context, r *http.Request) ([]byte, error) {
	const action = "CreateTaskHandler ExecFunc"
	const method = "ExecFunc"

	t := time.Now()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		err = error_with_codes.ErrorFailedToReadBody
		logrus.WithFields(logrus.Fields{
			"handlerName": handlerName,
			"method":      method,
		}).WithError(err).Error(action)
		return nil, err
	}

	taskDto := new(task_dto.TaskRequest)
	err = json.Unmarshal(body, taskDto)
	if err != nil {
		err = error_with_codes.ErrorFailedToCast
		logrus.WithFields(logrus.Fields{
			"handlerName": handlerName,
			"method":      method,
		}).WithError(err).Error(action)
		return nil, err
	}

	taskVO := new(task_model.Task)
	taskVO.Title = taskDto.Title

	task, err := h.taskUsecase.CreateTask(ctx, taskVO)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"handlerName": handlerName,
			"method":      method,
		}).WithError(err).Error(action)
		return nil, err
	}

	response := task_dto.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		Status:    task.Status,
		CreatedAt: task.CreatedAt,
		RunTime:   task.RunTime.String(),
	}

	json, err := result.NewResultOk(response, time.Since(t)).GetJson()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"handlerName": handlerName,
			"method":      method,
		}).WithError(err).Error(action)
		return nil, err
	}

	return json, nil
}
