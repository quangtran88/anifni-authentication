package serviceAdapters

import (
	"context"
	"encoding/json"
	"github.com/quangtran88/anifni-authentication/constants"
	"github.com/quangtran88/anifni-authentication/core/domain"
	"github.com/quangtran88/anifni-authentication/core/ports"
)

type NotificationService struct {
	kafka ports.KafkaProducer
}

func NewNotificationService(kafkaProducer ports.KafkaProducer) *NotificationService {
	return &NotificationService{kafkaProducer}
}

func (srv NotificationService) SendOTPEmail(ctx context.Context, otp domain.EmailOTP) error {
	body, err := json.Marshal(otp)
	if err != nil {
		return err
	}
	return srv.kafka.Produce(ctx, constants.SendOTPEmailTopic, otp.Email, string(body))
}
