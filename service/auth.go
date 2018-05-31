package service

import (
	"context"

	"github.com/jehaby/webapp102/service/auth"
)

func TknFromCtx(ctx context.Context) (string, error) {
	tknRow := ctx.Value(auth.StrTokenCtxKey)
	if tknRow == nil {
		return "", ErrNotAuthorized
	}
	if strToken, _ := tknRow.(string); strToken != "" {
		return strToken, nil
	}
	return "", ErrNotAuthorized
}
