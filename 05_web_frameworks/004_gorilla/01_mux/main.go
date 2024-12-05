package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// 新しいルーターを作成
	r := mux.NewRouter()

	// ルートの定義
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", UserHandler).Methods("GET")

	// サーバーの起動
	http.ListenAndServe(":8080", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	fmt.Fprintf(w, "User ID: %s", userId)
}
