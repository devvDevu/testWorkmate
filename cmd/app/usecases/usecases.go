package usecases

import (
	"testWorkmate/cmd/app/services"
	"testWorkmate/internal/usecase/task_usecase"

	"github.com/sirupsen/logrus"
)

type Usecases struct {
	taskUsecase *task_usecase.TaskUsecase
}

func NewUsecases() *Usecases {
	return &Usecases{
		taskUsecase: nil,
	}
}

func (u *Usecases) GetTaskUsecase() *task_usecase.TaskUsecase {
	return u.taskUsecase
}

func (u *Usecases) MustInit(services *services.Services) {
	logrus.Info("usecases start initializing")
	{
		u.taskUsecase = task_usecase.NewTaskUsecase(services.GetTaskService())
		logrus.Info("taskUsecase initialized")
	}
	logrus.Info("done")
}
