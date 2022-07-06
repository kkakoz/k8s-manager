package logic

import (
	"context"
	"github.com/kkakoz/ormx/opt"
	"k8s-manager/model"
	"k8s-manager/pkg/cryption"
	"k8s-manager/pkg/errno"
	"k8s-manager/repo"
	"k8s-manager/request"
)

type UserLogic struct {
	userRepo *repo.UserRepo
}

func NewUserLogic(userRepo *repo.UserRepo) *UserLogic {
	return &UserLogic{userRepo: userRepo}

}

func (item *UserLogic) Login(ctx context.Context, req *request.LoginReq) (string, error) {
	user, err := item.userRepo.Get(ctx, opt.Where("name = ?", req.Name))
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errno.New400("账号不存在")
	}
	if user.Password != cryption.Md5Str(req.Password+user.Salt) {
		return "", errno.New400("密码错误")
	}
	return item.userRepo.CacheAdd(ctx, user)
}

func (item *UserLogic) Current(ctx context.Context, token string) (*model.User, error) {
	return item.userRepo.CacheGet(ctx, token)
}
