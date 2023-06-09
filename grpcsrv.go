// ## gRPC Server: musayerapi/grpcsrv.go
//
// ServeGrpc serves the gRPC server for the SayerService.
package sayerapigo

import (
	"context"
	"net"
	"os"
	"reflect"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"github.com/cdfmlr/ellipsis"
	sayerv1 "github.com/murchinroom/sayerapigo/proto"
)

// sayerServiceServerImpl implements the SayerServiceServer interface
// generated by buf.
type sayerServiceServerImpl struct {
	sayer Sayer
	sayerv1.UnimplementedSayerServiceServer
}

func newSayerServiceServerImpl(sayer Sayer) *sayerServiceServerImpl {
	return &sayerServiceServerImpl{
		sayer: sayer,
	}
}

func (s *sayerServiceServerImpl) Say(
	ctx context.Context, in *sayerv1.SayRequest) (
	*sayerv1.SayResponse, error) {
	logger := slog.With("role", ellipsis.Centering(in.Role, 9), "text", ellipsis.Centering(in.Text, 9))

	resp, err := s.say(ctx, in)
	if err != nil {
		logger.Warn("[sayerServiceServer] Say (TTS) failed.", "err", err)
		return nil, err
	}
	logger.Info("[sayerServiceServer] Say (TTS) succeeded.")
	return resp, nil
}

func (s *sayerServiceServerImpl) say(
	ctx context.Context, in *sayerv1.SayRequest) (
	*sayerv1.SayResponse, error) {

	// if in.Role == "" {
	// 	return nil, status.Error(codes.InvalidArgument, "role is required")
	// }
	if in.Text == "" {
		return nil, status.Error(codes.InvalidArgument, "text is required")
	}

	format, audio, err := s.sayer.Say(in.Role, in.Text)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &sayerv1.SayResponse{
		Format: format,
		Audio:  audio,
	}, nil
}

// ServeGrpc serves the Sayer service on the given address.
//
// The context is used to control the lifetime of the server.
func ServeGrpc(ctx context.Context, sayer Sayer, addr string) error {
	server := grpc.NewServer()

	sayerv1.RegisterSayerServiceServer(server, newSayerServiceServerImpl(sayer))
	reflection.Register(server)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		server.GracefulStop()
		lis.Close()
	}()

	slog.Info("gRPC API server started.",
		"addr", addr,
		"sayer", reflect.TypeOf(sayer).String(),
		"pid", os.Getpid())
	return server.Serve(lis)
}
