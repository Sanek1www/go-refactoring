package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"refactoring/handler/request/user"
)

func (handler Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := handler.uCase.DeleteUser(id); err != nil {
		_ = render.Render(w, r, user.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}
