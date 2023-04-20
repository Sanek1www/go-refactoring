package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"refactoring/handler/request/user"
	"refactoring/handler/transformer"
)

func (handler Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	updateRequest := user.UpdateUserRequest{}

	if err := render.Bind(r, &updateRequest); err != nil {
		_ = render.Render(w, r, user.ErrInvalidRequest(err))
		return
	}

	userToUpdate := transformer.UpdateRequestToUser(updateRequest)
	id := chi.URLParam(r, "id")

	err := handler.uCase.UpdateUser(id, userToUpdate)
	if err != nil {
		_ = render.Render(w, r, user.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}
