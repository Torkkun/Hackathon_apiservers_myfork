package usecase

import (
	"github.com/pkg/errors"
	"yandereca.tech/yandereca/domain"

	"log"
)

type TaskInteractor struct {
	TaskRepository TaskRepository
}

func (interactor *TaskInteractor) Create(meeting domain.Todo) (message domain.TodoSuccessMessage, err error) {
	identifier, err := interactor.TaskRepository.Store(meeting)
	if err != nil {
		message.Message = "Failed"
		message.Result = false
		return
	}
	message.TaskId = identifier
	message.Message = "Success"
	message.Result = true
	// _, err = interactor.TaskRepository.FindById(identifier)

	return
}
func (interactor *TaskInteractor) CalcProgress(id string) (progress float64, err error) {
	progress, err = interactor.TaskRepository.CalcProgress(id)
	if err != nil {
		log.Println(errors.WithStack(err))
	}
	return
}
