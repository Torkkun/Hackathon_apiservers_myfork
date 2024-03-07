package usecase

import (
	"yandereca.tech/yandereca/domain"
)

type UserDataInteractor struct {
	UserDataRepository UserDataRepository
}

func (interactor *UserDataInteractor) Create(meeting domain.UserData) (message domain.UserDataSuccessMessage, err error) {
	identifier, err := interactor.UserDataRepository.Store(meeting)
	if err != nil {
		message.Message = "Failed"
		message.Result = false
		return
	}

	message.UserId = identifier
	message.Message = "Success"
	message.Result = true
	return
}
