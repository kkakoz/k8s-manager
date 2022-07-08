package mdctx

import "context"

func GetString(ctx context.Context, key any) string {
	return ctx.Value(key).(string)
}
