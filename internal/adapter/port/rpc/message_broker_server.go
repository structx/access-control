// Package rpc gRPC interface implementation
package rpc

import (
	"context"
	"encoding/json"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/structx/go-pkg/proto/messaging/v1"

	"github.com/structx/access-control/internal/core/domain"
)

// GRPCServer protobuf implementation of messaging service
type GRPCServer struct {
	pb.UnimplementedMessagingServiceV1Server

	log *zap.SugaredLogger
	ac  domain.AccessController
}

// interface compliance
var _ pb.MessagingServiceV1Server = (*GRPCServer)(nil)

// New constructor
func New(logger *zap.Logger, accessControl domain.AccessController) *GRPCServer {
	return &GRPCServer{
		log: logger.Sugar().Named("GrpcServer"),
		ac:  accessControl,
	}
}

// Publish not implemented
func (g *GRPCServer) Publish(context.Context, *pb.Envelope) (*pb.Stub, error) {
	// return empty response
	return &pb.Stub{}, nil
}

// RequestResponse messaging pattern
func (g *GRPCServer) RequestResponse(ctx context.Context, in *pb.Envelope) (*pb.Envelope, error) {

	var (
		result interface{}
		err    error
	)

	switch domain.RequestResponse(in.Topic) {
	case domain.VerifyServiceAccess:

		var ace domain.AccessControlEntry
		err = json.Unmarshal(in.Payload, &ace)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid topic payload")
		}

		err = g.ac.VerifyServiceAccess(ctx, &ace)
		if err != nil {

			resultbytes, err := json.Marshal(domain.EntryResponse{Granted: false})
			if err != nil {
				g.log.Errorf("failed to marshal result %v", err)
				return nil, status.Error(codes.Internal, "failed to marshal result")
			}

			return &pb.Envelope{
				Topic:   in.Topic,
				Payload: resultbytes,
			}, nil
		}

		result = domain.EntryResponse{Granted: true}

	case domain.VerifyUserAccess:

		var ace domain.AccessControlEntry
		err = json.Unmarshal(in.Payload, &ace)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid topic payload")
		}

		err = g.ac.VerifyUserAccess(ctx, &ace)
		if err != nil {

			resultbytes, err := json.Marshal(domain.EntryResponse{Granted: false})
			if err != nil {
				g.log.Errorf("failed to marshal result %v", err)
				return nil, status.Error(codes.Internal, "failed to marshal result")
			}

			return &pb.Envelope{
				Topic:   in.Topic,
				Payload: resultbytes,
			}, nil
		}

		result = domain.EntryResponse{Granted: true}

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
func (g *GRPCServer) Subscribe(*pb.Subscription, pb.MessagingServiceV1_SubscribeServer) error {
	// return empty response
	return nil
}
