package grpcHandler

import (
	"github.com/quangtran88/anifni-authentication/adapters/repositories"
	serviceAdapters "github.com/quangtran88/anifni-authentication/adapters/services"
	"github.com/quangtran88/anifni-authentication/core/services"
	baseUtils "github.com/quangtran88/anifni-base/libs/utils"
	authGRPC "github.com/quangtran88/anifni-grpc/authentication"
	"google.golang.org/grpc"
)

func InitGRPCServices(s *grpc.Server) {
	redisService := serviceAdapters.NewRedisService()
	kafkaProducer := serviceAdapters.NewKafkaProducer()

	otpRepo := repositories.NewOTPRepository(redisService)

	notiService := serviceAdapters.NewNotificationService(kafkaProducer)

	random := baseUtils.GetRandomGenerator()
	hash := baseUtils.GetHashGenerator()

	otpService := services.NewOTPService(otpRepo, notiService, random, hash)

	authGRPC.RegisterOTPServiceServer(s, NewOTPHandler(otpService))
}
