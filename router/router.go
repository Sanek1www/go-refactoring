package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	v1 "refactoring/router/routes/v1"
	"time"
)

func MakeRouter() (*chi.Mux, error) {
	r := chi.NewRouter()
	initBaseMiddlewares(r)

	return r, initRoutes(r)
}

func initBaseMiddlewares(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Use(middleware.Timeout(60 * time.Second))
}

func initRoutes(r *chi.Mux) error {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})

	err := v1.InitUsersRoute(r)
	if err != nil {
		return err
	}

	return nil
}
