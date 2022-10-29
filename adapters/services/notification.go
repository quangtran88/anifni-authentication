package serviceAdapters

import (
	"context"
	"encoding/json"
	"github.com/quangtran88/anifni-authentication/core/domain"
	"github.com/quangtran88/anifni-base/libs/event"
	basePorts "github.com/quangtran88/anifni-base/libs/ports"
)

type NotificationService struct {
	kafka basePorts.EventProducer
}

func NewNotificationService(kafkaProducer basePorts.EventProducer) *NotificationService {
	return &NotificationService{kafkaProducer}
}

func (srv NotificationService) SendOTPEmail(ctx context.Context, otp domain.EmailOTP) error {
	msg := event.SendEmailRequestMessage{
		Email:    "",
		Template: "",
		Params:   nil,
	}
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return srv.kafka.Produce(ctx, event.SendEmailRequestTopic, otp.Email, string(body))
}
