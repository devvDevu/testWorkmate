package handlers

import (
	"context"
	"net/http"
	"testWorkmate/cmd/app/usecases"
	"testWorkmate/internal/common/types/handler_type"
	"testWorkmate/internal/data_transfer_object/result"
	"testWorkmate/internal/handler/create_task_handler"
	"testWorkmate/internal/handler/delete_task_handler"
	"testWorkmate/internal/handler/get_task_handler"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const defaultExecTimeout = 3 * time.Second

type Handlers struct {
}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) MustInit(usecases *usecases.Usecases, router *mux.Router) {
	logrus.Info("handlers start initializing")
	{
		createTaskHandler := create_task_handler.NewCreateTaskHandler(usecases.GetTaskUsecase())
		initHandler(router, createTaskHandler)
	}
	{
		getTaskHandler := get_task_handler.NewGetTaskHandler(usecases.GetTaskUsecase())
		initHandler(router, getTaskHandler)
	}
	{
		deleteTaskHandler := delete_task_handler.NewDeleteTaskHandler(usecases.GetTaskUsecase())
		initHandler(router, deleteTaskHandler)
	}
	logrus.Info("done")
}

type handlerI interface {
	GetMethod() handler_type.HandlerMethod
	GetPath() handler_type.HandlerPath
	ExecFunc(ctx context.Context, r *http.Request) ([]byte, error)
}

func initHandler(router *mux.Router, h handlerI) {
	const action = "handler "

	l := logrus.WithFields(logrus.Fields{
		"handler": h.GetPath(),
		"method":  h.GetMethod(),
	})

	l.Info(action, "init")

	router.HandleFunc(h.GetPath().String(), func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		ctxToRun, cancel := context.WithTimeout(r.Context(), defaultExecTimeout)
		defer cancel()

		response, err := h.ExecFunc(ctxToRun, r)
		if err != nil {
			resultJson, err := result.NewResultErr(err).GetJson()
			if err != nil {
				l.WithError(err).
					Error(action, "NewResultErr(err).GetJson()")

				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			w.Write(resultJson)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		}
	}).Methods(h.GetMethod().String())
}
