package get_task_handler

import (
	"context"
	"net/http"
	"strconv"
	"testWorkmate/internal/common/types/error_with_codes"
	"testWorkmate/internal/common/types/handler_type"
	"testWorkmate/internal/data_transfer_object/result"
	"testWorkmate/internal/data_transfer_object/task_dto"
	"testWorkmate/internal/model/task_model"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const handlerName = "GetTaskHandler"

type GetTaskHandler struct {
	taskUsecase taskUsecaseI
}

type taskUsecaseI interface {
	GetTask(ctx context.Context, id uint64) (*task_model.Task, error)
}

func NewGetTaskHandler(taskUsecase taskUsecaseI) *GetTaskHandler {
	return &GetTaskHandler{
		taskUsecase: taskUsecase,
	}
}

func (h *GetTaskHandler) GetPath() handler_type.HandlerPath {
	return "/api/v1/task/{id}"
}

func (h *GetTaskHandler) GetMethod() handler_type.HandlerMethod {
	return handler_type.HandlerMethodGet
}

func (h *GetTaskHandler) ExecFunc(ctx context.Context, r *http.Request) ([]byte, error) {
	const action = "GetTaskHandler ExecFunc"
	const method = "ExecFunc"

	t := time.Now()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		err := error_with_codes.ErrorFailedToReadBody
		logrus.WithFields(logrus.Fields{
			"handlerName": handlerName,
			"method":      method,
		}).WithError(err).Error(action)
		return nil, err
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		err = error_with_codes.ErrorFailedToReadBody
		logrus.WithFields(logrus.Fields{
			"handlerName": handlerName,
			"method":      method,
		}).WithError(err).Error(action)
		return nil, err
	}

	task, err := h.taskUsecase.GetTask(ctx, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"handlerName": handlerName,
			"method":      method,
			"taskID":      id,
		}).WithError(err).Error(action)
		return nil, err
	}

	response := task_dto.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		Status:    task.Status,
		CreatedAt: task.CreatedAt,
		RunTime:   time.Since(t),
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
