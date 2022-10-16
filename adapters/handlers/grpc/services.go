package grpcHandler

import (
	otpGRPC "github.com/quangtran88/anifni-grpc/authentication"
	"google.golang.org/grpc"
)

func InitGRPCServices(s *grpc.Server) {
	otpGRPC.RegisterOTPServiceServer(s, NewOTPHandler())
}
