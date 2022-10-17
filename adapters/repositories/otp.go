package repositories

import (
	"context"
	"fmt"
	"github.com/quangtran88/anifni-authentication/core/domain"
	"github.com/quangtran88/anifni-authentication/core/ports"
	"time"
)

type OTPRepository struct {
	redis ports.RedisService
}

func NewOTPRepository(redis ports.RedisService) *OTPRepository {
	return &OTPRepository{redis}
}

func (repo OTPRepository) SaveEmailOTP(ctx context.Context, otp domain.EmailOTP) error {
	key := repo.getEmailOTPKey(otp.Email)
	return repo.redis.Set(ctx, key, otp.Code, time.Minute*domain.EmailOTPExpireMinute)
}

func (repo OTPRepository) GetEmailOTP(ctx context.Context, email string) (domain.EmailOTP, error) {
	key := repo.getEmailOTPKey(email)
	code, err := repo.redis.Get(ctx, key)
	if err != nil {
		return domain.EmailOTP{}, err
	}
	return domain.EmailOTP{Email: email, Code: code}, nil
}

func (repo OTPRepository) DeleteEmailOTP(ctx context.Context, email string) error {
	key := repo.getEmailOTPKey(email)
	return repo.redis.Del(ctx, key)
}

func (repo OTPRepository) getEmailOTPKey(email string) string {
	return fmt.Sprintf("otp:%s", email)
}
