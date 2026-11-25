package tasks_postgres_repository

import (
	"time"

	"github.com/condonar/ToDoApp-REST-API/internal/core/domain"
)

type TaskModel struct {
	ID           int
	Version      int
	Titile       string
	Description  *string
	Completed    bool
	CreatedAt    time.Time
	CompletedAt  *time.Time
	AuthorUserID int
}

func taskDomainFromModel(taskModel TaskModel) domain.Task {
	return domain.NewTask(
		taskModel.ID,
		taskModel.Version,
		taskModel.Titile,
		taskModel.Description,
		taskModel.Completed,
		taskModel.CreatedAt,
		taskModel.CompletedAt,
		taskModel.AuthorUserID,
	)
}

func taskDomainsFromModels(taskModels []TaskModel) []domain.Task {
	domains := make([]domain.Task, len(taskModels))

	for i, model := range taskModels {
		domains[i] = taskDomainFromModel(model)
	}

	return domains
}
