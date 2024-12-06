package handler

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest/enum"
)

// I18n国际化翻译中间件，如果发生死机，该中间件将恢复。
func I18nHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// 将lang添加到上下文
		if lang := r.Header.Get("Accept-Language"); lang == "" {
			ctx = context.WithValue(ctx, enum.I18nCtxKey, "zh-CN")
		} else {
			ctx = context.WithValue(ctx, enum.I18nCtxKey, lang)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
