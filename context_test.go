package middleware

import (
	"context"
	"net/http"
	"testing"
)

func TestContextMiddleware(t *testing.T) {
	req, _ := http.NewRequest("GET", "/test", nil)

	AppEngineContext(nil, nil, req, func(ctx context.Context) {
		if ctx == nil {
			t.Fatalf("Expected context to not be nil")
		}
	})
}
