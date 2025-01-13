package main

import (
	"bytes"
	"fmt"
	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	main02()
}

// --  -- //
func main01() {
	// Heimdall HTTP クライアントを作成
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(10*time.Second), // 10秒のタイムアウトを設定
		httpclient.WithRetryCount(3),               // リトライ回数を3回に設定
	)

	// GET リクエストを送信
	resp, err := client.Get("https://httpbin.org/get", nil)
	if err != nil {
		log.Fatalf("リクエスト失敗: %v", err)
	}
	defer resp.Body.Close()

	// レスポンスの内容を出力
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("ステータスコード: %d\n", resp.StatusCode)
	fmt.Printf("レスポンスボディ: %s\n", string(body))
}

// --  -- //
func main02() {
	// Heimdall HTTP クライアントを作成
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(10*time.Second), // 10秒のタイムアウトを設定
		httpclient.WithRetryCount(3),               // リトライ回数を3回に設定
	)

	// JSON データを送信
	data := []byte(`{"message": "Hello, Heimdall!"}`)
	headers := http.Header{
		"Content-Type": []string{"application/json"},
	}

	// POST リクエストを送信
	resp, err := client.Post("https://httpbin.org/post", bytes.NewBuffer(data), headers)
	if err != nil {
		log.Fatalf("リクエスト失敗: %v", err)
	}
	defer resp.Body.Close()

	// レスポンスの内容を出力
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("ステータスコード: %d\n", resp.StatusCode)
	fmt.Printf("レスポンスボディ: %s\n", string(body))
}

// --  -- //
func main03() {
	// 固定バックオフを設定（リトライ間隔: 2秒）
	backoff := heimdall.NewConstantBackoff(2*time.Second, 1*time.Second)

	// HTTP クライアントを作成
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(10*time.Second),           // タイムアウトを10秒に設定
		httpclient.WithRetrier(heimdall.NewRetrier(backoff)), // リトライ機能を設定
	)

	// リクエストを実行
	response, err := client.Get("https://example.com", nil)
	if err != nil {
		fmt.Printf("リクエスト失敗: %v\n", err)
		return
	}

	fmt.Printf("レスポンスコード: %d\n", response.StatusCode)
}

// --  -- //
func main04() {
	// 指数バックオフを設定
	backoff := heimdall.NewExponentialBackoff(
		500*time.Millisecond,    // 初期待機時間（リトライの最初の間隔）
		5*time.Second,           // 最大待機時間（リトライ間隔の上限）
		float64(15*time.Second), // 最大経過時間（リトライ全体の制限時間）
		2.0,                     // 指数増加率（リトライ間隔を倍々に増加）
	)

	// HTTP クライアントを作成
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(10*time.Second),           // タイムアウトを 10 秒に設定
		httpclient.WithRetrier(heimdall.NewRetrier(backoff)), // リトライ機能を設定
	)

	// GET リクエストを実行
	response, err := client.Get("https://example.com", nil)
	if err != nil {
		fmt.Printf("リクエスト失敗: %v\n", err)
		return
	}

	// レスポンスコードを出力
	fmt.Printf("レスポンスコード: %d\n", response.StatusCode)
}

// --  -- //
func main05() {

}

// --  -- //
func main06() {

}

// --  -- //
func main07() {

}

// --  -- //
func main08() {

}

// --  -- //
func main09() {

}

// --  -- //
func main10() {

}
