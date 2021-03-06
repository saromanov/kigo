package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/saromanov/kigo/internal/rest/handlers"
	"github.com/saromanov/kigo/internal/generate/service"
)

// Init provides initialization of the rest
func Init(ctx context.Context, cfg Config, srv service.Generate) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handlers.Create().ServeHTTP)
	if err := http.ListenAndServe(cfg.Address, r); err != nil {
		return err
	}
	return nil
}