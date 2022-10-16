package ports

import "github.com/quangtran88/anifni-authentication/core/domain"

type OTPService interface {
	SendEmailOTP(email string) error
	CheckEmailOTP(otp domain.EmailOTP) (bool, error)
}
