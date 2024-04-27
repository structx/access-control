package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/trevatk/anastasia/internal/core/domain"
	pkgdomain "github.com/trevatk/go-pkg/domain"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

// Bundle
type Bundle struct {
	log *zap.SugaredLogger
	ac  domain.AccessController
	m   pkgdomain.MessageBroker
}

// NewBundle
func NewBundle(logger *zap.Logger, accessControl domain.AccessController) *Bundle {
	return &Bundle{
		ac:  accessControl,
		log: logger.Sugar().Named("AccessControlSubscriber"),
	}
}

// Subscribe
func (b *Bundle) Subscribe(ctx context.Context) error {

	var result error
	cs := make([]<-chan pkgdomain.Envelope, 0)

	subs := domain.ListSubscriptions()
	for _, sub := range subs {
		t := sub.String()
		ch, err := b.m.Subscribe(ctx, t)
		if err != nil {
			result = multierr.Append(result, fmt.Errorf("unable to subscribe to %s %v"))
		}
		cs = append(cs, ch)
	}

	ch := merge(ctx, cs...)

	g, ctx := errgroup.WithContext(ctx)
	for i := 0; i < len(subs); i++ {
		g.Go(func() error {
			return b.subscriber(ctx, ch)
		})
	}

	err := g.Wait()
	if err != nil {
		result = multierr.Append(result, fmt.Errorf("unable to wait for subscribers to work %v", err))
	}

	return result
}

func (b *Bundle) subscriber(ctx context.Context, ch <-chan pkgdomain.Envelope) error {

	for {

		var (
			err error
		)

		select {
		case <-ctx.Done():
			return nil
		case msg, ok := <-ch:

			if !ok {
				return nil
			}

			t := msg.GetTopic()
			switch domain.Subscriptions(msg.GetTopic()) {
			case domain.ModifyAccessControlList:
				_, err = b.modifyAccessControlList(ctx, msg.GetPayload())
			default:
				return errors.New("invalid topic " + t)
			}

			if err != nil {
				// error occured
				b.log.Errorf("failed to complete operation %s %v", t, err)
			}
		}
	}
}

func merge(ctx context.Context, cs ...<-chan pkgdomain.Envelope) <-chan pkgdomain.Envelope {

	out := make(chan pkgdomain.Envelope)

	output := func(c <-chan pkgdomain.Envelope) {
		for n := range c {
			select {
			case <-ctx.Done():
				return
			case out <- n:

			}
		}
	}

	for _, c := range cs {
		go output(c)
	}

	go func() {
		<-ctx.Done()
		close(out)
	}()

	return out
}
