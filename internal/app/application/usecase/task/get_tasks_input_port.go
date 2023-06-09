package task

import taskDomain "go-ca/internal/app/domain/task"

type GetTasksInputPort interface {
	GetTasks() ([]GetTasksContainContractDto, *taskDomain.Task)
}
