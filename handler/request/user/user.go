package user

import "net/http"

type baseUserRequest struct {
}

func (c *baseUserRequest) Bind(r *http.Request) error { return nil }

type CreateUserRequest struct {
	baseUserRequest
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

type UpdateUserRequest struct {
	baseUserRequest
	DisplayName string `json:"display_name"`
}
