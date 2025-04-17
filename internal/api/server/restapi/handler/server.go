package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/Maksim646/tokens/internal/api/definition"
	"go.uber.org/zap"
)

func (h *Handler) WsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		path := r.URL.Path
		authHeader := r.Header.Get("Authorization")

		var principal *definition.Principal
		var err error

		switch {
		case strings.HasPrefix(path, "/auth/token"), path == "/docs", path == "/swagger.json":
			zap.L().Debug("Unauthenticated route accessed", zap.String("path", path))
		case path == "/auth/refresh":
			principal, err = h.ValidateExpiredOrValidToken(authHeader)
		default:
			principal, err = h.ValidateAccessTokenOnly(authHeader)
		}

		if err != nil {
			zap.L().Warn("Unauthorized access", zap.String("path", path), zap.Error(err))
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		if principal != nil {
			ctx := context.WithValue(r.Context(), contextKeyUserID, principal.ID)
			ctx = context.WithValue(ctx, contextKeyIP, principal.IP)
			ctx = context.WithValue(ctx, contextKeyRefreshID, principal.RefreshID)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}
