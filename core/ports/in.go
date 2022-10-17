package ports

import (
	"context"
	"github.com/quangtran88/anifni-authentication/core/domain"
)

type OTPService interface {
	SendEmailOTP(ctx context.Context, email string) error
	CheckEmailOTP(ctx context.Context, otp domain.EmailOTP) (bool, error)
}
