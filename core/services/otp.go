package services

import (
	"github.com/quangtran88/anifni-authentication/core/domain"
	"github.com/quangtran88/anifni-authentication/core/ports"
)

type OTPService struct {
	repo    ports.OTPRepository
	notiSrv ports.NotificationService
}

func NewOTPService(repo ports.OTPRepository, notiSrv ports.NotificationService) *OTPService {
	return &OTPService{repo: repo, notiSrv: notiSrv}
}

func (srv OTPService) SendEmailOTP(email string) error {
	//TODO implement me
	panic("implement me")
}

func (srv OTPService) CheckEmailOTP(otp domain.EmailOTP) (bool, error) {
	//TODO implement me
	panic("implement me")
}
