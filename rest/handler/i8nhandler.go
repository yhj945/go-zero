package handler

import (
	"context"
	"net/http"
)

// I18n国际化翻译中间件，如果发生死机，该中间件将恢复。
func I18nHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// 将lang添加到上下文
		ctx = context.WithValue(ctx, "lang", r.Header.Get("Accept-Language"))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
