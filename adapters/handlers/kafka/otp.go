package kafkaHandlers

import (
	"context"
	"github.com/quangtran88/anifni-authentication/core/ports"
	"github.com/quangtran88/anifni-base/libs/event"
	basePorts "github.com/quangtran88/anifni-base/libs/ports"
	baseUtils "github.com/quangtran88/anifni-base/libs/utils"
)

type OTPHandler struct {
	otpSrv ports.OTPService
}

func NewOTPHandler(otpService ports.OTPService) *OTPHandler {
	return &OTPHandler{otpService}
}

func (h OTPHandler) HandleSendOTP(msg basePorts.EventMessage) {
	var body event.SendOTPRequestMessage
	err := baseUtils.ParseJSON(msg.Value, &body)
	if err != nil {
		return
	}
	_ = h.otpSrv.SendEmailOTP(context.Background(), body.Email)
}
