package mdctx

import (
	"context"
	"net/http"
)

func NewCtx(req *http.Request) context.Context {
	return req.Context()
}

const (
	NS = "namespace"
)
