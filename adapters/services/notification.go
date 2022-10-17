package serviceAdatpers

import (
	"github.com/quangtran88/anifni-authentication/core/domain"
)

type NotificationService struct {
}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (srv NotificationService) SendOTPEmail(otp domain.EmailOTP) error {
	return nil
}
