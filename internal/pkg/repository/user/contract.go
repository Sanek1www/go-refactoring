package user

import "refactoring/internal/pkg/entity/user"

type ContractRepository interface {
	Add(*user.User) (string, error)
	GetAll() (map[string]user.User, error)
	Find(id string) (user.User, error)
	Inc() error
	Update(string, *user.User) error
	Delete(id string) error
}
