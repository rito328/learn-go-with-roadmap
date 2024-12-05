package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
)

// セッションストアの作成（32バイトのキーを使用）
var store = sessions.NewCookieStore([]byte("secret-key-32-bytes-long-password"))

func main() {
	r := mux.NewRouter()

	// ルートの定義
	r.HandleFunc("/login", LoginHandler).Methods("GET")
	r.HandleFunc("/profile", ProfileHandler).Methods("GET")
	r.HandleFunc("/logout", LogoutHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}

// ログイン処理
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")

	// セッションへの値の設定
	session.Values["authenticated"] = true
	session.Values["user"] = "John Doe"

	// セッションの保存
	session.Save(r, w)

	fmt.Fprintln(w, "You have been logged in!")
}

// プロフィール表示（セッションチェック付き）
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")

	// 認証チェック
	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		http.Error(w, "Please log in first", http.StatusUnauthorized)
		return
	}

	// ユーザー名の取得
	username := session.Values["user"].(string)
	fmt.Fprintf(w, "Welcome, %s!", username)
}

// ログアウト処理
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")

	// セッション値のクリア
	session.Values["authenticated"] = false
	delete(session.Values, "user")

	session.Save(r, w)

	fmt.Fprintln(w, "You have been logged out!")
}
