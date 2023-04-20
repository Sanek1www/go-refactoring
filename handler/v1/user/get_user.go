package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"refactoring/handler/request/user"
)

func (handler Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	u, err := handler.uCase.GetUser(id)
	if err != nil {
		_ = render.Render(w, r, user.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, u)
}
