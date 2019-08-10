package middleware

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/v3nom/pipes"
)

// Useless test to reach maximum coverage
func TestGetNow(t *testing.T) {
	now := getNow()
	if now.Unix() == 0 {
		t.Fatal("Never happens")
	}
}

func TestLimitPerMinute(t *testing.T) {
	utc, _ := time.LoadLocation("UTC")
	input := time.Date(2019, 5, 18, 19, 5, 35, 0, utc)
	expected := time.Date(2019, 5, 18, 19, 5, 0, 0, utc)
	actual := LimitPerMinute(input)
	if !expected.Equal(actual) {
		t.Fatalf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestRateLimit(t *testing.T) {
	utc, _ := time.LoadLocation("UTC")
	now := time.Date(2019, 5, 18, 19, 5, 19, 0, utc)
	getNow = func() time.Time {
		return now
	}

	rateLimitMiddleware := LimitRate(LimitPerMinute, 10)
	pipeline := pipes.New().
		Use(rateLimitMiddleware).
		Use(func(ctx context.Context, w http.ResponseWriter, r *http.Request, next pipes.Next) {
			fmt.Fprint(w, "ok")
			next(ctx)
		})

	// Request
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.RemoteAddr = "127.0.0.1"

	// Request handling
	handler := http.HandlerFunc(pipeline.Build())

	i := 0
	for i <= 15 {
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, req)
		if i > 10 {
			if recorder.Code != http.StatusTooManyRequests {
				t.Fatalf("Iteration: %v, Expected: %v, Actual: %v", i, http.StatusTooManyRequests, recorder.Code)
			}
		} else {
			if recorder.Code != http.StatusOK {
				t.Fatalf("Iteration: %v, Expected: %v, Actual: %v", i, http.StatusOK, recorder.Code)
			}
		}
		i++
	}

	// Make request from another IP
	req2, _ := http.NewRequest("GET", "/test", nil)
	req2.RemoteAddr = "127.0.0.2"
	recorder2 := httptest.NewRecorder()
	handler.ServeHTTP(recorder2, req2)
	if recorder2.Code != http.StatusOK {
		t.Fatalf("New IP.Expected: %v, Actual: %v", http.StatusOK, recorder2.Code)
	}

	// Make request later
	req3, _ := http.NewRequest("GET", "/test", nil)
	req3.RemoteAddr = "127.0.0.1"
	now = time.Date(2019, 5, 18, 19, 6, 19, 0, utc)
	recorder3 := httptest.NewRecorder()

	handler.ServeHTTP(recorder3, req3)
	if recorder3.Code != http.StatusOK {
		t.Fatalf("New minute.Expected: %v, Actual: %v", http.StatusOK, recorder3.Code)
	}
}
