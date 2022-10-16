package grpcHandler

import (
	"context"
	"github.com/quangtran88/anifni-grpc/authentication"
)

type OTPHandler struct {
	otpGRPC.UnimplementedOTPServiceServer
}

func NewOTPHandler() *OTPHandler {
	return &OTPHandler{}
}

func (handler OTPHandler) CheckEmailOTP(ctx context.Context, in *otpGRPC.CheckEmailOTPInput) (*otpGRPC.CheckEmailOTPResult, error) {
	return &otpGRPC.CheckEmailOTPResult{}, nil
}

func (handler OTPHandler) SendEmailOTP(ctx context.Context, in *otpGRPC.SendEmailOTPInput) (*otpGRPC.SendEmailOTPResult, error) {
	return &otpGRPC.SendEmailOTPResult{}, nil
}
