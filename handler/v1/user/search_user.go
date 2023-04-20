package user

import (
	"github.com/go-chi/render"
	"net/http"
	"refactoring/handler/request/user"
)

func (handler Handler) SearchUser(w http.ResponseWriter, r *http.Request) {
	users, err := handler.uCase.SearchUser()
	if err != nil {
		_ = render.Render(w, r, user.ErrInvalidRequest(err))
	}

	render.JSON(w, r, users)
}
