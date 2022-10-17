package domain

const EmailOTPLength = 6
const EmailOTPExpireMinute = 10

type EmailOTP struct {
	Email string
	Code  string
}
