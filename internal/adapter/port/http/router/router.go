// Package router http chi router provider
package router

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/structx/go-pkg/adapter/port/http/controller"
)

// New constructor
func New(logger *zap.Logger) *chi.Mux {

	r := chi.NewRouter()

	cc := []interface{}{
		controller.NewBundle(logger),
	}

	for _, c := range cc {

		if c0, ok := c.(controller.V0); ok {
			h := c0.RegisterRoutesV0()
			r.Mount("/", h)
		}
	}

	return r
}
