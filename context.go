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
