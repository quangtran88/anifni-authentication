package services

import (
	"context"
	"github.com/quangtran88/anifni-authentication/core/domain"
	"github.com/quangtran88/anifni-authentication/core/ports"
)

type OTPService struct {
	repo    ports.OTPRepository
	notiSrv ports.NotificationService
	random  ports.RandomGenerator
	hash    ports.HashGenerator
}

func NewOTPService(
	repo ports.OTPRepository,
	notiSrv ports.NotificationService,
	random ports.RandomGenerator,
	hash ports.HashGenerator,
) *OTPService {
	return &OTPService{repo, notiSrv, random, hash}
}

func (srv OTPService) SendEmailOTP(ctx context.Context, email string) error {
	code := srv.random.GetDigit(domain.EmailOTPLength)
	hashedCode, err := srv.hash.HashPassword(code)
	if err != nil {
		return err
	}

	otp := domain.EmailOTP{
		Email: email,
		Code:  hashedCode,
	}

	err = srv.notiSrv.SendOTPEmail(otp)
	if err != nil {
		return err
	}

	err = srv.repo.DeleteEmailOTP(ctx, email)
	if err != nil {
		return err
	}

	err = srv.repo.SaveEmailOTP(ctx, otp)
	if err != nil {
		return err
	}

	return nil
}

func (srv OTPService) CheckEmailOTP(ctx context.Context, otp domain.EmailOTP) (bool, error) {
	fetched, err := srv.repo.GetEmailOTP(ctx, otp.Email)
	if err != nil {
		return false, err
	}

	if ok := srv.hash.CheckPasswordHash(otp.Code, fetched.Code); !ok {
		return false, nil
	}

	return true, nil
}
