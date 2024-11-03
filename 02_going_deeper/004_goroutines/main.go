package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
}

// -- 基本形 -- //
func sayHello() {
	fmt.Println("Hello, World!")
}
func main0101() {
	go sayHello()               // ゴルーチンを起動
	time.Sleep(1 * time.Second) // メイン関数が終了しないように少し待機
}

// -- sync.WaitGroupの例 -- //
func printMessage(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(msg)
}
func main0201() {
	var wg sync.WaitGroup

	wg.Add(1)
	go printMessage("Hello from goroutine", &wg)

	wg.Wait() // ゴルーチンが終了するのを待機
	fmt.Println("Main function ends")
}

// -- 主な用途 -- //

// --- 1. 並列なI/O処理 --- //
func fetchURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching:", url, err)
		return
	}
	fmt.Println("Fetched", url, "with status", resp.Status)
}
func main0301() {
	var wg sync.WaitGroup
	urls := []string{"https://example.com", "https://example.org", "https://example.net"}

	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, &wg)
	}

	wg.Wait()
	fmt.Println("All URLs fetched.")
}

// --- 2. 並行して計算処理を行う --- //
func sum(array []int, result *int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, v := range array {
		sum += v
	}
	*result = sum
}
func main0302() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var result1, result2 int
	var wg sync.WaitGroup

	wg.Add(2) // これから開始するゴルーチンの数を引数に渡す
	go sum(array[:len(array)/2], &result1, &wg)
	go sum(array[len(array)/2:], &result2, &wg)

	wg.Wait()
	total := result1 + result2
	fmt.Println("Total sum:", total)
}

// --- 3. チャネルを使った非同期なタスク管理 --- //
// ワーカー関数
// 各ワーカーは、jobsチャネルからタスクを受け取り、完了したら結果をresultsチャネルに送信します。
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs { // jobsチャネルが閉じるまでタスクを受け取る
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second) // 処理のシミュレーション（1秒待つ）
		results <- job * 2      // 結果を results チャネルに送信
	}
}
func main0303() {
	jobs := make(chan int, 5)    // タスクを送信するチャネル
	results := make(chan int, 5) // 結果を受け取るチャネル

	// 3 つのワーカーを起動（それぞれのワーカーが並行してタスクを処理する）
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 5 つのタスクを jobs チャネルに送信
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // タスクの送信が完了したのでjobsチャネルを閉じる

	// 5つの処理結果を受け取る
	for a := 1; a <= 5; a++ {
		result := <-results
		fmt.Println("Result:", result)
	}
}

// --- 4. タイムアウト処理 --- //
func doTask(done chan bool) {
	time.Sleep(2 * time.Second) // タスクが完了するまで2秒かかる
	done <- true
}
func main0304() {
	done := make(chan bool, 1)
	go doTask(done)

	select {
	case <-done:
		fmt.Println("Task completed")
	case <-time.After(1 * time.Second):
		fmt.Println("Task timed out")
	}
}

// -- ゴルーチン使用の注意点 -- //
// --- 2. ゴルーチンリーク（ゴルーチンが終了しない） --- //
func leakyGoroutine(ch <-chan int) {
	for val := range ch {
		fmt.Println("Received:", val)
	}
	fmt.Println("Goroutine exiting")
}
func main040201() {
	ch := make(chan int)

	go leakyGoroutine(ch)

	time.Sleep(1 * time.Second)
	// ここで何も送信せずに終了する
	fmt.Println("Main function ending")
}
func fetchData(wg *sync.WaitGroup) {
	defer wg.Done() // 完了を通知する
	time.Sleep(2 * time.Second)
	fmt.Println("Data fetched")
}
func main040202() {
	var wg sync.WaitGroup
	wg.Add(1) // ゴルーチンの数を追加

	go fetchData(&wg)

	wg.Wait() // すべてのゴルーチンが完了するまで待機
	fmt.Println("Main function ending")
}

// --- 7. タイムアウトとキャンセル --- //
func main040701() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("ゴルーチンがキャンセルされました:", ctx.Err())
		}
	}(ctx)
}

// --- 8. パニック処理 --- //
func safeGoroutine() {
	defer func() {
		if r := recover(); r != nil { // パニックをキャッチ
			fmt.Println("Recovered from panic:", r) // ログに記録する
		}
	}()

	fmt.Println("Goroutine started")
	time.Sleep(1 * time.Second)
	panic("Something went wrong!") // パニック発生
	fmt.Println("This line will not execute")
}
func main040801() {
	go safeGoroutine() // パニック処理があるゴルーチンを起動

	// メイン関数の処理を続行
	time.Sleep(2 * time.Second) // ゴルーチンが終了するのを待つ
	fmt.Println("Main function completed")
}
