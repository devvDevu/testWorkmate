package repositories

import (
	"testWorkmate/cmd/app/adapters"
	"testWorkmate/internal/repository/task_repository"

	"github.com/sirupsen/logrus"
)

type Repositories struct {
	taskRepository *task_repository.TaskRepository
}

func NewRepositories() *Repositories {
	return &Repositories{
		taskRepository: nil,
	}
}

func (r *Repositories) GetTaskRepository() *task_repository.TaskRepository {
	return r.taskRepository
}

func (r *Repositories) MustInit(adapters *adapters.Adapters) {
	logrus.Info("repositories start initializing")
	{
		r.taskRepository = task_repository.NewTaskRepository(adapters.GetImDb())
		logrus.Info("taskRepository initialized")
	}
	logrus.Info("done")
}
