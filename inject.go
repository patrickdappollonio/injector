package injector

import (
	"context"
	"net/http"
)

// Context allows you to use it as a middleware to store a value with a given name
// on the request context. You can use it in a router globally or per-route (if supported
// by your router / muxer)
func Context(name string, value interface{}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), name, value)
			next.ServeHTTP(w, r.WithContext(ctx))

		}
		return http.HandlerFunc(fn)
	}
}
