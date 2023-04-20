package user

import (
	"refactoring/internal/pkg/entity/user"
	userRepo "refactoring/internal/pkg/repository/user"
)

type Usecase struct {
	repo userRepo.ContractRepository
}

func New(userRepo userRepo.ContractRepository) *Usecase {
	return &Usecase{repo: userRepo}
}

func (u *Usecase) SearchUser() (map[string]user.User, error) {
	return u.repo.GetAll()
}

func (u *Usecase) GetUser(id string) (user.User, error) {
	return u.repo.Find(id)
}

func (u *Usecase) CreateUser(user *user.User) (string, error) {
	err := u.repo.Inc()
	if err != nil {
		return "", err
	}

	return u.repo.Add(user)
}

func (u *Usecase) UpdateUser(id string, user *user.User) error {
	return u.repo.Update(id, user)
}

func (u *Usecase) DeleteUser(id string) error {
	return u.repo.Delete(id)
}
