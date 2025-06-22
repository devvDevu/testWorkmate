package services

import (
	"testWorkmate/cmd/app/repositories"
	"testWorkmate/internal/service/task_service"

	"github.com/sirupsen/logrus"
)

type Services struct {
	taskService *task_service.TaskService
}

func NewServices() *Services {
	return &Services{
		taskService: nil,
	}
}

func (s *Services) GetTaskService() *task_service.TaskService {
	return s.taskService
}

func (s *Services) MustInit(repositories *repositories.Repositories) {
	logrus.Info("services start initializing")
	{
		s.taskService = task_service.NewTaskService(repositories.GetTaskRepository())
		logrus.Info("taskService initialized")
	}
	logrus.Info("done")
}
