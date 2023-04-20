package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/golobby/container/v3"
	"refactoring/handler/v1/user"
)

func InitUsersRoute(r *chi.Mux) error {
	handler := user.Handler{}
	err := container.Fill(&handler)
	if err != nil {
		return err
	}

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/", handler.SearchUser)
				r.Post("/", handler.CreateUser)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", handler.GetUser)
					r.Patch("/", handler.UpdateUser)
					r.Delete("/", handler.DeleteUser)
				})
			})
		})
	})

	return nil
}
