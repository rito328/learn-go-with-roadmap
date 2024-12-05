package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	// 新しいルーターを作成
	r := mux.NewRouter()

	// 基本的なルートの設定
	r.HandleFunc("/", HomeHandler).Methods("GET")

	// ミドルウェアの適用
	// 1. ログ出力
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	// 2. CORS設定
	corsRouter := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(loggedRouter)

	// 3. レスポンス圧縮
	compressedRouter := handlers.CompressHandler(corsRouter)

	// 4. パニックリカバリー
	finalRouter := handlers.RecoveryHandler()(compressedRouter)

	// サーバーの起動
	http.ListenAndServe(":8080", finalRouter)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page!"))
}
