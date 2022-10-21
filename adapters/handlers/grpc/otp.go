package grpcHandler

import (
	"context"
	"github.com/quangtran88/anifni-authentication/core/domain"
	"github.com/quangtran88/anifni-authentication/core/ports"
	authGRPC "github.com/quangtran88/anifni-grpc/authentication"
)

type AuthHandler struct {
	authGRPC.UnimplementedOTPServiceServer
	otpService ports.OTPService
}

func NewOTPHandler(otpService ports.OTPService) *AuthHandler {
	return &AuthHandler{otpService: otpService}
}

func (handler AuthHandler) CheckEmailOTP(ctx context.Context, in *authGRPC.CheckEmailOTPInput) (*authGRPC.CheckEmailOTPResult, error) {
	otp := domain.EmailOTP{Email: in.Email, Code: in.Code}
	ok, err := handler.otpService.CheckEmailOTP(ctx, otp)
	if err != nil {
		return &authGRPC.CheckEmailOTPResult{Ok: false}, err
	}
	return &authGRPC.CheckEmailOTPResult{Ok: ok}, nil
}

func (handler AuthHandler) SendEmailOTP(ctx context.Context, in *authGRPC.SendEmailOTPInput) (*authGRPC.SendEmailOTPResult, error) {
	err := handler.otpService.SendEmailOTP(ctx, in.Email)
	if err != nil {
		return &authGRPC.SendEmailOTPResult{Ok: false}, err
	}
	return &authGRPC.SendEmailOTPResult{Ok: true}, nil
}
