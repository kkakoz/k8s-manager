package mdctx

import "context"

func GetString(ctx context.Context, key any) string {
	return ctx.Value(key).(string)
}

func GetNs(ctx context.Context) string {
	return GetString(ctx, NS)
}
