package mdctx

import (
	"context"
	"net/http"
)

func NewCtx(req *http.Request) context.Context {
	ctx := req.Context()
	ns := req.Header.Get(NS)
	return context.WithValue(ctx, NS, ns)
}

const (
	NS = "namespace"
)
