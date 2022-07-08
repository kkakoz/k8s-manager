package repo

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/kkakoz/ormx"
	"k8s-manager/model"
	"k8s-manager/pkg/cryption"
	"k8s-manager/pkg/keys"
	"time"
)

type UserRepo struct {
	ormx.IRepo[model.User]
	redis *redis.Client
}

func NewUserRepo(redis *redis.Client) *UserRepo {
	return &UserRepo{IRepo: ormx.NewRepo[model.User](), redis: redis}
}

func (item *UserRepo) CacheAdd(ctx context.Context, user *model.User) (string, error) {
	token := cryption.UUID()
	data, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	err = item.redis.Set(keys.TokenKey(token), data, time.Hour*24*3).Err()
	return token, err
}

func (item *UserRepo) CacheGet(ctx context.Context, token string) (*model.User, error) {
	res, err := item.redis.WithContext(ctx).Get(keys.TokenKey(token)).Result()
	if err != nil {
		return nil, err
	}
	user := &model.User{}
	err = json.Unmarshal([]byte(res), user)
	return user, err
}
