package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/v3nom/pipes"
)

func TestPanic(t *testing.T) {
	pipeline := pipes.New().Use(PanicHandler(func(ctx context.Context, w http.ResponseWriter, r *http.Request, err interface{}) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)
	})).Use(func(ctx context.Context, w http.ResponseWriter, r *http.Request, next pipes.Next) {
		panic(errors.New("Test Error"))
	})

	// Request
	req, _ := http.NewRequest("GET", "/test", nil)

	// Request handling
	handler := http.HandlerFunc(pipeline.Build())

	// Record
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("Expected status code: %v, Actual: %v", http.StatusInternalServerError, recorder.Code)
	}

	if recorder.Body.String() != "Error: Test Error" {
		t.Fatal("Expected body text")
	}
}
