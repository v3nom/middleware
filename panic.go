package middleware

import (
	"context"
	"net/http"

	"github.com/v3nom/pipes"
)

// PanicHandler creates middleware which recovers from panic in the pipeline and can return user friedly message or page.
// Should be added as early as possible in the pipeline to handle panics from other middlewares.
func PanicHandler(handler func(ctx context.Context, w http.ResponseWriter, r *http.Request, err interface{})) pipes.Middleware {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, next pipes.Next) {
		defer func() {
			if err := recover(); err != nil {
				handler(ctx, w, r, err)
			}
		}()
		next(ctx)
	}
}
