package router

import (
	"encoding/json"
	"github.com/labstack/echo"
	"k8s-manager/local"
	"k8s-manager/model"
	"k8s-manager/pkg/errno"
	"k8s-manager/pkg/keys"
	"k8s-manager/pkg/redis"
)

func authority(f echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.Request().Header.Get("X-Authorization")
		if token == "" {
			return errno.NewErr(401, 401, "请重新登录")
		}
		client := redis.GetClient()
		user := &model.User{}
		result, err := client.Get(keys.TokenKey(token)).Result()
		if err != nil {
			return errno.NewErr(401, 401, "请重新登录")
		}
		err = json.Unmarshal([]byte(result), user)
		if err != nil {
			return errno.NewErr(401, 401, "请重新登录")
		}
		ctx.Request().Header.Add(local.UserLocalKey, result)
		return f(ctx)
	}
}
