// Package main entrypoint of application
package main

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/multierr"

	"github.com/structx/go-pkg/adapter/logging"
	"github.com/structx/go-pkg/adapter/port/messagebroker"
	"github.com/structx/go-pkg/adapter/setup"
	pkgdomain "github.com/structx/go-pkg/domain"
	"github.com/structx/go-pkg/util/decode"

	"github.com/structx/access-control/internal/adapter/port/http/router"
	"github.com/structx/access-control/internal/adapter/port/rpc"
	"github.com/structx/access-control/internal/core/domain"
	"github.com/structx/access-control/internal/core/graph"
	"github.com/structx/access-control/internal/core/service"
)

func main() {
	fx.New(
		fx.Provide(fx.Annotate(setup.New, fx.As(new(pkgdomain.Config)))),
		fx.Invoke(decode.ConfigFromEnv),
		fx.Provide(logging.New),
		fx.Provide(fx.Annotate(messagebroker.New, fx.As(new(pkgdomain.MessageBroker)))),
		fx.Provide(fx.Annotate(service.NewAccessControl, fx.As(new(domain.AccessController)))),
		fx.Provide(fx.Annotate(graph.New, fx.As(new(domain.Graph)))),
		fx.Provide(fx.Annotate(service.NewBundle, fx.As(new(domain.Bundler)))),
		fx.Provide(fx.Annotate(router.New, fx.As(new(http.Handler)))),
		fx.Provide(rpc.New),
		fx.Invoke(registerHooks),
	).Run()
}

func registerHooks(lc fx.Lifecycle, bundler domain.Bundler) error {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {

				var result error

				err := bundler.Subscribe(ctx)
				if err != nil {
					result = multierr.Append(result, fmt.Errorf("failed to subscribe to service to topics %v", err))
				}

				return result
			},
			OnStop: func(_ context.Context) error {

				var result error

				err := bundler.Close()
				if err != nil {
					result = multierr.Append(result, fmt.Errorf("failed to close service bundle %v", err))
				}

				return result
			},
		},
	)
	return nil
}
