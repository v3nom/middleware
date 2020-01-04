package middleware

import (
	"net/http"
	"time"

	"context"

	"github.com/v3nom/pipes"
)

type rateEntry struct {
	IPs map[string]int
}

var getNow = func() time.Time {
	return time.Now()
}

// LimitRate creates rate limitter middleware using provided time bucketing function
func LimitRate(timeFun func(time.Time) time.Time, limit int) pipes.Middleware {
	var rateWindow string
	entry := &rateEntry{
		IPs: map[string]int{},
	}

	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, next pipes.Next) {
		ip := r.RemoteAddr
		currentWindow := timeFun(getNow()).String()

		if currentWindow != rateWindow {
			rateWindow = currentWindow
			entry = &rateEntry{
				IPs: map[string]int{
					ip: 1,
				},
			}
		} else if entry.IPs[ip] > limit {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		} else {
			entry.IPs[ip]++
		}
		next(ctx)
	}
}

// LimitPerMinute round time to a minute
func LimitPerMinute(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
}
