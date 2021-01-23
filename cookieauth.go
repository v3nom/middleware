package middleware

import (
	"net/http"

	"github.com/v3nom/pipes"

	"context"
)

type cookieDecoder = func(context.Context, *http.Cookie) map[string]string
type userFactory = func(map[string]string) interface{}

// User context key
const User = pipes.ContextKey("User")

// CookieAuth creates cookie auth middleware.
func CookieAuth(authCookieName string, cookieDec cookieDecoder, userFac userFactory) pipes.Middleware {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, next pipes.Next) {
		// Check if cookie exists
		cookie, err := r.Cookie(authCookieName)
		if err != nil {
			next(ctx)
			return
		}

		// Decode cookie
		values := cookieDec(ctx, cookie)
		if values == nil {
			next(ctx)
			return
		}

		// Create user
		user := userFac(values)
		if user == nil {
			next(ctx)
			return
		}

		next(context.WithValue(ctx, User, user))
	}
}
