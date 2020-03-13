package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/v3nom/pipes"
)

func TestContextMiddleware(t *testing.T) {
	req, _ := http.NewRequest("GET", "/test", nil)

	Context(nil, nil, req, func(ctx context.Context) {
		if ctx == nil {
			t.Fatalf("Expected context to not be nil")
		}
	})
}

var aKey pipes.ContextKey = "akey"
var bKey pipes.ContextKey = "bkey"

func TestAddContextMiddleware(t *testing.T) {
	options := map[pipes.ContextKey]interface{}{
		aKey: "a",
		bKey: "b",
	}

	pipeline := pipes.New().Use(Context).Use(AddOptions(options)).Use(func(ctx context.Context, w http.ResponseWriter, r *http.Request, next pipes.Next) {
		a := ctx.Value(aKey).(string)
		b := ctx.Value(bKey).(string)
		if a+b == "ab" {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		next(ctx)
	})

	// Request
	req, _ := http.NewRequest("GET", "/test", nil)

	// Request handling
	handler := http.HandlerFunc(pipeline.Build())

	// Record
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)

	if recorder.Code == http.StatusInternalServerError {
		t.Fatalf("Expected status code: %v, Actual: %v", http.StatusOK, recorder.Code)
	}
}
