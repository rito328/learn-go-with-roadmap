package middleware

import (
	"log"
	"net/http"
	"time"
)

type RequestDurationLoggerMiddleware struct {
}

func NewRequestDurationLoggerMiddleware() *RequestDurationLoggerMiddleware {
	return &RequestDurationLoggerMiddleware{}
}

func (m *RequestDurationLoggerMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() // 処理開始時間を記録

		// 次のハンドラーを実行
		next(w, r)

		// 処理時間を計測
		duration := time.Since(start)

		// ログに出力
		log.Printf("[RequestDuration] %s %s took %v", r.Method, r.URL.Path, duration)
	}
}
