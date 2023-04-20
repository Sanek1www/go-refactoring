package transformer

import (
	request "refactoring/handler/request/user"
	"refactoring/internal/pkg/entity/user"
	"time"
)

func CreateRequestToUser(request request.CreateUserRequest) *user.User {
	return &user.User{
		CreatedAt:   time.Now(),
		DisplayName: request.DisplayName,
		Email:       request.Email,
	}
}

func UpdateRequestToUser(request request.UpdateUserRequest) *user.User {
	return &user.User{
		CreatedAt:   time.Now(),
		DisplayName: request.DisplayName,
	}
}
