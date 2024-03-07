package usecase

import "yandereca.tech/yandereca/domain"

type TaskRepository interface {
	Store(domain.Todo) (string, error)
	FindById(string) (domain.Todos, error)
	CalcProgress(string) (float64, error)
}
