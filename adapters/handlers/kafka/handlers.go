package kafkaHandlers

import (
	"github.com/quangtran88/anifni-authentication/adapters/repositories"
	serviceAdapters "github.com/quangtran88/anifni-authentication/adapters/services"
	"github.com/quangtran88/anifni-authentication/core/services"
	"github.com/quangtran88/anifni-base/libs/event"
	baseServices "github.com/quangtran88/anifni-base/libs/services"
	baseUtils "github.com/quangtran88/anifni-base/libs/utils"
)

func InitKafkaHandlers() {
	redisService := serviceAdapters.NewRedisService()
	kafkaProducer := baseServices.NewKafkaProducer()
	notiService := serviceAdapters.NewNotificationService(kafkaProducer)

	otpRepo := repositories.NewOTPRepository(redisService)

	random := baseUtils.GetRandomGenerator()
	hash := baseUtils.GetHashGenerator()

	otpSrv := services.NewOTPService(otpRepo, notiService, random, hash)
	otpHandler := NewOTPHandler(otpSrv)

	consumer := baseServices.NewKafkaConsumer()

	consumer.Consume(event.SendOTPRequestTopic, event.HandleSendOTPGroup, otpHandler.HandleSendOTP)
}
