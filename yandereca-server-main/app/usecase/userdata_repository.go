package usecase

import "yandereca.tech/yandereca/domain"

type UserDataRepository interface {
	Store(domain.UserData) (string, error)
	FindById(string) (domain.UserDataList, error)
	Remove(string) error
}
