package middlewares

import (
    "context"
    "go.uber.org/zap"
    "net/http"
)

func InjectLogger(logger *zap.Logger) func(next http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx := context.WithValue(r.Context(), "logger", logger)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}
