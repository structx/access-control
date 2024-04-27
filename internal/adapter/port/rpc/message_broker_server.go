// Package rpc gRPC interface implementation
package rpc

import (
	"context"
	"encoding/json"

	"github.com/trevatk/anastasia/internal/core/domain"
	pb "github.com/trevatk/go-pkg/proto/messaging/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GrpcServer ...
type GrpcServer struct {
	pb.UnimplementedMessagingServiceV1Server

	log *zap.SugaredLogger
	ac  domain.AccessController
}

// interface compliance
var _ pb.MessagingServiceV1Server = (*GrpcServer)(nil)

// New constructor
func New(logger *zap.Logger, accessControl domain.AccessController) *GrpcServer {
	return &GrpcServer{
		log: logger.Sugar().Named("GrpcServer"),
		ac:  accessControl,
	}
}

// Publish not implemented
func (g *GrpcServer) Publish(context.Context, *pb.Envelope) (*pb.Stub, error) {
	// return empty response
	return &pb.Stub{}, nil
}

// RequestResponse
func (g *GrpcServer) RequestResponse(ctx context.Context, in *pb.Envelope) (*pb.Envelope, error) {

	var (
		result interface{}
		err    error
	)

	switch domain.RequestResponse(in.Topic) {
	default:
		return nil, status.Errorf(codes.InvalidArgument, "invalid topic provided")
	}

	if err != nil {
		g.log.Errorf("failed to perform message operation %v", err)
		return nil, status.Error(codes.Internal, "unable to perform message operation")
	}

	resultbytes, err := json.Marshal(result)
	if err != nil {
		g.log.Errorf("failed to marshal result %v", err)
		return nil, status.Errorf(codes.Internal, "unable to marshal result object")
	}

	return &pb.Envelope{
		Topic:   in.Topic,
		Payload: resultbytes,
	}, nil
}

// Subscribe not implemented
func (g *GrpcServer) Subscribe(*pb.Subscription, pb.MessagingServiceV1_SubscribeServer) error {
	// return empty response
	return nil
}
