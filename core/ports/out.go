package ports

import "github.com/quangtran88/anifni-authentication/core/domain"

type OTPRepository interface {
	SaveEmailOTP(otp domain.EmailOTP) error
	GetEmailOTP(email string) (string, error)
}

type NotificationService interface {
	SendOTPEmail(otp domain.EmailOTP) error
}
