package user

import "refactoring/internal/app/usecase/user"

type Handler struct {
	uCase *user.Usecase `container:"type"`
}
