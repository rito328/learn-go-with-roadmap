package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	main020303()
}

// -- Marshalling -- //

// --- 基本的な使い方 --- //
type User struct {
	ID       int    `json:"id"` // JSON のキー名を指定
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func main0101() {
	// 構造体にデータを格納
	user := User{
		ID:       1,
		Name:     "John Doe",
		IsActive: true,
	}

	// JSON に変換
	jsonData, err := json.Marshal(user)
	if err != nil {
		// エラー処理
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// JSON を文字列として表示
	fmt.Println(string(jsonData))
}

// --- インデント付き JSON --- //
func main010101() {
	user := User{
		ID:       1,
		Name:     "John Doe",
		IsActive: true,
	}

	//jsonData, err := json.MarshalIndent(user, "", "  ")
	jsonData, err := json.MarshalIndent(user, "", "\t") // タブを使用
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
	}
	fmt.Println(string(jsonData))
}

// --- ネストされた構造体 --- //
type AddressAA struct {
	City  string `json:"city"`
	State string `json:"state"`
}

type UserAAA struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Address AddressAA `json:"address"` // ネストされた構造体
}

func main010102() {
	// ネストされた構造体を含むデータを定義
	user := UserAAA{
		ID:   1,
		Name: "Alice",
		Address: AddressAA{
			City:  "Osaka",
			State: "Osaka Prefecture",
		},
	}

	// JSON に変換
	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// JSON を表示
	fmt.Println(string(jsonData))
}

// --- 構造体タグのオプション --- //
// ---- omitempty ---- //
type UserA struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`      // 空文字列なら省略
	IsActive bool   `json:"is_active,omitempty"` // false なら省略
	Address  string `json:"address,omitempty"`   // 空文字列なら省略
}

func main010201() {
	user := UserA{
		ID: 1,
	}

	jsonData, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(jsonData))
}

// ----- ネストされた構造体と omitempty ----- //
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}
type UserB struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Address *Address `json:"address,omitempty"` // Address がゼロ値の場合は省略
}

func main01020101() {
	user := UserB{
		ID:   1,
		Name: "Alice",
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonData))
}

// ---- -（フィールドを出力しない） ---- //
type UserC struct {
	ID       int    `json:"id"`
	Password string `json:"-"` // パスワードを JSON に出力しない
}

func main010202() {
	user := UserC{
		ID:       1,
		Password: "password",
	}

	jsonData, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(jsonData))
}

// --- カスタムエンコーディング --- //
type Event struct {
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Event) MarshalJSON() ([]byte, error) {
	type Alias Event // 再帰的定義を避けるための別名型
	return json.Marshal(&struct {
		Timestamp string `json:"timestamp"` // 独自フォーマット
		*Alias
	}{
		Timestamp: e.Timestamp.Format("2006-01-02 15:04:05"),
		Alias:     (*Alias)(&e),
	})
}
func main0103() {
	event := Event{
		Name:      "Conference",
		Timestamp: time.Now(),
	}
	jsonData, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonData))
}

// --- html.EscapeString の抑制 --- //
func main0104() {
	data := map[string]string{
		"html": "<div>Hello</div>",
	}

	// `SetEscapeHTML(false)` を設定した場合
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetEscapeHTML(false) // HTML エスケープを無効化
	err := encoder.Encode(data)
	if err != nil {
		return
	}

	// エスケープしている場合（通常の `json.Marshal`）
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonData))
}

// --- Map や Interface の利用 --- //
func main0105() {
	data := map[string]interface{}{
		"id":   1,
		"name": "Alice",
		"tags": []string{"developer", "golang"},
	}
	jsonData, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(jsonData))
}

// -- Unmarshalling -- //
// --- 基本的な使い方 --- //
type User0201 struct {
	ID       int    `json:"id"` // JSON のキー名を指定
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

// --- 基本的な使い方 --- //
func main0201() {
	// JSON 形式のデータ
	jsonStr := `{"id":1,"name":"Alice","is_active":true}`

	// 構造体の変数を用意
	var user User0201

	// JSON を構造体に変換
	err := json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		// エラー処理
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// 変換されたデータを表示
	fmt.Printf("%+v\n", user)
}

// --- 構造体以外のデータ型も利用可能 --- //
// ---- マップへの変換 ---- //
func main020301() {
	// 動的な JSON データ
	jsonStr := `{"id": 1, "name": "Alice", "is_active": true}`

	// マップ型の変数を用意
	var data map[string]interface{}

	// JSON をマップに変換
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// 変換結果を表示
	fmt.Printf("ID: %v\n", data["id"])
	fmt.Printf("Name: %v\n", data["name"])
	fmt.Printf("IsActive: %v\n", data["is_active"])
}

// ---- スライスへの変換 ---- //
func main020302() {
	// 配列形式の JSON データ
	jsonArrayStr := `[{"id": 1, "name": "Alice"}, {"id": 2, "name": "Bob"}]`

	// スライス型の変数を用意
	var users []map[string]interface{}

	// JSON をスライスに変換
	err := json.Unmarshal([]byte(jsonArrayStr), &users)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// 変換結果を表示
	for _, user := range users {
		fmt.Printf("ID: %v, Name: %v\n", user["id"], user["name"])
	}
}

// ---- インターフェース型への変換 ---- //
func main020303() {
	// 不明な形式の JSON データ
	jsonStr := `{"id": 1, "name": "Alice", "tags": ["developer", "golang"]}`

	// インターフェース型の変数を用意
	var data interface{}

	// JSON をインターフェース型に変換
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// 型アサーションで具体的な型にアクセス
	if obj, ok := data.(map[string]interface{}); ok {
		fmt.Printf("ID: %v\n", obj["id"])
		fmt.Printf("Name: %v\n", obj["name"])
		fmt.Printf("Tags: %v\n", obj["tags"])
	}
}

// --  -- //
func main0202() {

}
