package main

import (
	"encoding/json"
	"fmt"
	"github.com/olahol/melody"
	"net/http"
)

func main() {
	main03()
}

// --  -- //
func main01() {
	m := melody.New()
	fmt.Println("Melody インスタンスが作成されました:", m)
}

// --  -- //
func main02() {
	// Melody のインスタンスを作成
	m := melody.New()

	// 静的な HTML を提供するハンドラー
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// WebSocket リクエストを処理するハンドラー
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		err := m.HandleRequest(w, r)
		if err != nil {
			fmt.Printf("error handling request: %v\n", err)
		}
	})

	// メッセージ受信時の処理
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		// 受信したメッセージをすべてのクライアントにブロードキャスト
		err := m.Broadcast(msg)
		if err != nil {
			fmt.Printf("broadcast error: %v\n", err)
		}
	})

	// サーバー起動
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

// メッセージの構造体
type Message struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
}

// --  -- //
func main03() {
	m := melody.New()

	// クライアント接続時にユーザー名をセッションに保存
	m.HandleConnect(func(s *melody.Session) {
		query := s.Request.URL.Query()
		username := query.Get("username")
		if username == "" {
			username = "匿名ユーザー"
		}
		// セッションにユーザー名を保存
		s.Set("username", username)
	})

	// メッセージ受信時の処理
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		// セッションからユーザー名を取得
		username, exists := s.Get("username")
		if !exists {
			username = "匿名ユーザー"
		}

		// 送信メッセージを構築
		message := Message{
			Sender:  username.(string),
			Content: string(msg),
		}

		// メッセージを JSON に変換
		messageJSON, err := json.Marshal(message)
		if err != nil {
			return
		}

		// ブロードキャスト
		err = m.Broadcast(messageJSON)
		if err != nil {
			fmt.Printf("broadcast error: %v\n", err)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		err := m.HandleRequest(w, r)
		if err != nil {
			fmt.Printf("error handling request: %v\n", err)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
