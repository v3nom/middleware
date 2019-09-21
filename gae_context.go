package middleware

import (
	"context"
	"net/http"

	"github.com/v3nom/pipes"

	"google.golang.org/appengine"
)

// AppEngineContext adds Google App Engine context to the request pipeline
func AppEngineContext(ctx context.Context, w http.ResponseWriter, r *http.Request, next pipes.Next) {
	next(appengine.NewContext(r))
}
