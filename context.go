package middleware

import (
	"context"
	"net/http"

	"github.com/v3nom/pipes"
)

// Context adds context to the request pipeline
func Context(ctx context.Context, w http.ResponseWriter, r *http.Request, next pipes.Next) {
	next(r.Context())
}

// AddOptions adds context options
func AddOptions(options map[pipes.ContextKey]interface{}) pipes.Middleware {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, next pipes.Next) {
		if options != nil {
			for k, v := range options {
				ctx = context.WithValue(ctx, k, v)
			}
		}
		next(ctx)
	}
}
