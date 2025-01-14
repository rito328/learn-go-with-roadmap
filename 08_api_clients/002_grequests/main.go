package main

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"net/http"
	"os"
	"sync"
)

func main() {
	main06()
}

// -- GET -- //
func main01() {
	// リクエストオプションの設定
	ro := &grequests.RequestOptions{
		Headers: map[string]string{
			"User-Agent": "MyGoApp",
		},
	}

	// GET リクエストを送信
	resp, err := grequests.Get("https://jsonplaceholder.typicode.com/posts/1", ro)
	if err != nil {
		fmt.Println("リクエストに失敗しました:", err)
		return
	}

	// レスポンスを出力
	fmt.Println(resp.String())
}

// -- POST -- //
func main02() {
	// 送信するデータ
	data := map[string]string{
		"title":  "foo",
		"body":   "bar",
		"userId": "1",
	}

	// リクエストオプションの設定
	ro := &grequests.RequestOptions{
		JSON: data, // JSON データを指定
	}

	// POST リクエストを送信
	resp, err := grequests.Post("https://jsonplaceholder.typicode.com/posts", ro)
	if err != nil {
		fmt.Println("リクエストに失敗しました:", err)
		return
	}

	// レスポンスを出力
	fmt.Println(resp.String())
}

// -- 並行処理との組み合わせ -- //
func fetchURL(wg *sync.WaitGroup, url string) {
	defer wg.Done()

	resp, err := grequests.Get(url, nil)
	if err != nil {
		fmt.Printf("リクエスト失敗: %s, エラー: %v\n", url, err)
		return
	}

	fmt.Printf("URL: %s, レスポンス: %s\n", url, resp.String())
}
func main03() {
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go fetchURL(&wg, url)
	}

	wg.Wait()
	fmt.Println("すべてのリクエストが完了しました")
}

// -- ファイルのダウンロード -- //
func main04() {
	resp, err := grequests.Get("https://placehold.jp/150x150.png", nil)
	if err != nil {
		fmt.Println("ファイルダウンロードに失敗しました:", err)
		return
	}
	defer resp.Close()

	file, err := os.Create("image.png")
	if err != nil {
		fmt.Println("ファイル作成に失敗しました:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(resp.Bytes())
	if err != nil {
		fmt.Println("ファイル書き込みに失敗しました:", err)
		return
	}

	fmt.Println("ファイルのダウンロードが完了しました: image.png")
}

// -- カスタムクッキーの利用 -- //
func main05() {
	// クッキーの設定
	cookies := []*http.Cookie{
		{
			Name:  "session_id",
			Value: "abc123",
		},
	}

	// リクエストオプションの設定
	ro := &grequests.RequestOptions{
		Cookies: cookies,
	}

	// リクエストを送信
	resp, err := grequests.Get("https://httpbin.org/cookies", ro)
	if err != nil {
		fmt.Println("リクエスト失敗:", err)
		return
	}

	// レスポンスを出力
	fmt.Println("レスポンス:", resp.String())
}

// -- レスポンスデータのパース -- //
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main06() {
	resp, err := grequests.Get("https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Println("リクエスト失敗:", err)
		return
	}

	var post Post
	err = json.Unmarshal(resp.Bytes(), &post)
	if err != nil {
		fmt.Println("JSON パース失敗:", err)
		return
	}

	fmt.Printf("投稿情報: %+v\n", post)
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
