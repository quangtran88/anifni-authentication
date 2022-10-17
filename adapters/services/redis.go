package serviceAdatpers

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/quangtran88/anifni-authentication/constants"
	baseUtils "github.com/quangtran88/anifni-base/libs/utils"
	"time"
)

const RedisTimeoutSec = 8

type RedisService struct {
	client *redis.Client
}

func NewRedisService() *RedisService {
	env := baseUtils.GetEnvManager()

	client := redis.NewClient(&redis.Options{
		Addr:     env.GetEnv(constants.RedisURIEnvKey),
		Password: env.GetEnv(constants.RedisPasswordEnvKey),
	})
	pong, err := client.Ping(context.Background()).Result()
	fmt.Println("Connect redis", pong, err)
	return &RedisService{client}
}

func (srv RedisService) Set(ctx context.Context, key string, value string, exp time.Duration) error {
	ctx, cancel := srv.wrapCtx(ctx)
	defer cancel()
	return srv.client.Set(ctx, key, value, exp).Err()
}

func (srv RedisService) Get(ctx context.Context, key string) (string, error) {
	ctx, cancel := srv.wrapCtx(ctx)
	defer cancel()
	return srv.client.Get(ctx, key).Result()
}

func (srv RedisService) Del(ctx context.Context, key string) error {
	ctx, cancel := srv.wrapCtx(ctx)
	defer cancel()
	return srv.client.Del(ctx, key).Err()
}

func (srv RedisService) wrapCtx(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, RedisTimeoutSec*time.Second)
}
