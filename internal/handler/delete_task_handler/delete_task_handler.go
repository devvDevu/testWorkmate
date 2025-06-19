package delete_task_handler

import (
	"context"
	"net/http"
	"strconv"
	"testWorkmate/internal/common/types/error_with_codes"
	"testWorkmate/internal/common/types/handler_type"
	"testWorkmate/internal/data_transfer_object/result"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const handlerName = "DeleteTaskHandler"

type DeleteTaskHandler struct {
	taskUsecase taskUsecaseI
}

type taskUsecaseI interface {
	DeleteTask(ctx context.Context, id uint64) error
}

func NewDeleteTaskHandler(taskUsecase taskUsecaseI) *DeleteTaskHandler {
	return &DeleteTaskHandler{
		taskUsecase: taskUsecase,
	}
}

func (h *DeleteTaskHandler) GetPath() handler_type.HandlerPath {
	return "/api/v1/task/{id}"
}

func (h *DeleteTaskHandler) GetMethod() handler_type.HandlerMethod {
	return handler_type.HandlerMethodDelete
}

func (h *DeleteTaskHandler) ExecFunc(ctx context.Context, r *http.Request) ([]byte, error) {
	const action = "DeleteTaskHandler ExecFunc"
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

	if err := h.taskUsecase.DeleteTask(ctx, id); err != nil {
		logrus.WithFields(logrus.Fields{
			"handlerName": handlerName,
			"method":      method,
			"taskID":      id,
		}).WithError(err).Error(action)
		return nil, err
	}

	json, err := result.NewResultOk("ok", time.Since(t)).GetJson()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"handlerName": handlerName,
			"method":      method,
		}).WithError(err).Error(action)
		return nil, err
	}

	return json, nil
}
