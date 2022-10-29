package grpcHandler

import (
	"context"
	"github.com/quangtran88/anifni-authentication/core/domain"
	"github.com/quangtran88/anifni-authentication/core/ports"
	authGRPC "github.com/quangtran88/anifni-grpc/authentication"
)

type OTPHandler struct {
	authGRPC.UnimplementedOTPServiceServer
	otpService ports.OTPService
}

func NewOTPHandler(otpService ports.OTPService) *OTPHandler {
	return &OTPHandler{otpService: otpService}
}

func (handler OTPHandler) CheckEmailOTP(ctx context.Context, in *authGRPC.CheckEmailOTPInput) (*authGRPC.CheckEmailOTPResult, error) {
	otp := domain.EmailOTP{Email: in.Email, Code: in.Code}
	ok, err := handler.otpService.CheckEmailOTP(ctx, otp)
	if err != nil {
		return &authGRPC.CheckEmailOTPResult{Ok: false}, err
	}
	return &authGRPC.CheckEmailOTPResult{Ok: ok}, nil
}
