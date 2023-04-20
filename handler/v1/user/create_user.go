package user

import (
	"github.com/go-chi/render"
	"net/http"
	"refactoring/handler/request/user"
	"refactoring/handler/transformer"
)

func (handler Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	createRequest := user.CreateUserRequest{}

	if err := render.Bind(r, &createRequest); err != nil {
		_ = render.Render(w, r, user.ErrInvalidRequest(err))
		return
	}

	userToCreate := transformer.CreateRequestToUser(createRequest)

	id, err := handler.uCase.CreateUser(userToCreate)
	if err != nil {
		_ = render.Render(w, r, user.ErrInvalidRequest(err))
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})
}
