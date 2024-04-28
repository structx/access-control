package domain

import "context"

// Bundle service global interface
type Bundler interface {
	Subscribe(ctx context.Context) error
	Close() error
}
