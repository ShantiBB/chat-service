package helpers

import (
	"context"
	"time"
)

var Cancel context.CancelFunc

func WithTimeout(ctx context.Context, d time.Duration) context.Context {
	cctx, c := context.WithTimeout(ctx, d)
	Cancel = c
	return cctx
}
