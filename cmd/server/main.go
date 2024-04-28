// Package main entrypoint of application
package main

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/multierr"

	"github.com/trevatk/anastasia/internal/adapter/port/http/router"
	"github.com/trevatk/anastasia/internal/adapter/port/rpc"
	"github.com/trevatk/anastasia/internal/core/domain"
	"github.com/trevatk/anastasia/internal/core/service"
)

func main() {
	fx.New(
		fx.Provide(service.NewBundle),
		fx.Provide(fx.Annotate(service.NewAccessControl, fx.As(new(domain.AccessController)))),
		fx.Provide(fx.Annotate(router.New, fx.As(new(http.Handler)))),
		fx.Provide(rpc.New),
		fx.Invoke(registerHooks),
	).Run()
}

func registerHooks(lc fx.Lifecycle, bundle *service.Bundle) error {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {

				var result error

				err := bundle.Subscribe(ctx)
				if err != nil {
					result = multierr.Append(result, fmt.Errorf("failed to subscribe to service to topics %v", err))
				}

				return result
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		},
	)
	return nil
}
