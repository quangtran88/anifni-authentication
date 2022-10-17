package grpcHandler

import (
	"github.com/quangtran88/anifni-authentication/adapters/repositories"
	serviceAdatpers "github.com/quangtran88/anifni-authentication/adapters/services"
	"github.com/quangtran88/anifni-authentication/core/services"
	baseUtils "github.com/quangtran88/anifni-base/libs/utils"
	otpGRPC "github.com/quangtran88/anifni-grpc/authentication"
	"google.golang.org/grpc"
)

func InitGRPCServices(s *grpc.Server) {
	redisService := serviceAdatpers.NewRedisService()
	otpRepo := repositories.NewOTPRepository(redisService)
	notiService := serviceAdatpers.NewNotificationService()
	random := baseUtils.GetRandomGenerator()
	hash := baseUtils.GetHashGenerator()

	otpService := services.NewOTPService(otpRepo, notiService, random, hash)

	otpGRPC.RegisterOTPServiceServer(s, NewOTPHandler(otpService))
}
