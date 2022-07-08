package local

import (
	"context"
	"k8s-manager/model"
)

type userLocalKey struct{}

func WithUser(ctx context.Context, value *model.User) context.Context {
	return context.WithValue(ctx, userLocalKey{}, value)
}

func GetUser(ctx context.Context) (*model.User, bool) {
	v, ok := ctx.Value(userLocalKey{}).(*model.User)
	if v == nil {
		return nil, false
	}
	return v, ok
}

type namespaceKey struct{}

func WithNamespace(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, namespaceKey{}, value)
}

func GetNamespace(ctx context.Context) string {
	val, ok := ctx.Value(namespaceKey{}).(string)
	if !ok {
		return ""
	}
	return val
}
