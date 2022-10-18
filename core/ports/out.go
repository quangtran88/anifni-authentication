package ports

import (
	"context"
	"github.com/quangtran88/anifni-authentication/core/domain"
	"time"
)

type OTPRepository interface {
	SaveEmailOTP(ctx context.Context, otp domain.EmailOTP) error
	GetEmailOTP(ctx context.Context, email string) (domain.EmailOTP, error)
	DeleteEmailOTP(ctx context.Context, email string) error
}

type NotificationService interface {
	SendOTPEmail(ctx context.Context, otp domain.EmailOTP) error
}

type RandomGenerator interface {
	GetDigit(size int) string
}

type HashGenerator interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type RedisService interface {
	Set(ctx context.Context, key string, value string, exp time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}

type KafkaProducer interface {
	Produce(ctx context.Context, topic string, key string, value string) error
	ProduceMultiple(ctx context.Context, topic string, messages []KafkaMessage) error
}
