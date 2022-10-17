package grpcHandler

import (
	"context"
	"github.com/quangtran88/anifni-authentication/core/domain"
	"github.com/quangtran88/anifni-authentication/core/ports"
	"github.com/quangtran88/anifni-grpc/authentication"
)

type OTPHandler struct {
	otpGRPC.UnimplementedOTPServiceServer
	otpService ports.OTPService
}

func NewOTPHandler(otpService ports.OTPService) *OTPHandler {
	return &OTPHandler{otpService: otpService}
}

func (handler OTPHandler) CheckEmailOTP(ctx context.Context, in *otpGRPC.CheckEmailOTPInput) (*otpGRPC.CheckEmailOTPResult, error) {
	otp := domain.EmailOTP{Email: in.Email, Code: in.Code}
	ok, err := handler.otpService.CheckEmailOTP(ctx, otp)
	if err != nil {
		return &otpGRPC.CheckEmailOTPResult{Ok: false}, err
	}
	return &otpGRPC.CheckEmailOTPResult{Ok: ok}, nil
}

func (handler OTPHandler) SendEmailOTP(ctx context.Context, in *otpGRPC.SendEmailOTPInput) (*otpGRPC.SendEmailOTPResult, error) {
	err := handler.otpService.SendEmailOTP(ctx, in.Email)
	if err != nil {
		return &otpGRPC.SendEmailOTPResult{Ok: false}, err
	}
	return &otpGRPC.SendEmailOTPResult{Ok: true}, nil
}
