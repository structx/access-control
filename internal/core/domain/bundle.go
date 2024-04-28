package domain

import "context"

// Bundler service global interface
type Bundler interface {
	// Subscribe service to all subscription topics
	Subscribe(ctx context.Context) error
	// Close external service connections
	Close() error
}
