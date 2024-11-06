package main

import (
	"fmt"
	"time"
)

func main() {
}

// -- 具体的な使用例 -- //
// --- 1. 複数チャネルの監視 --- //
func main0101() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 送信側ゴルーチン1
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "メッセージ1"
	}()

	// 送信側ゴルーチン2
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "メッセージ2"
	}()

	// 複数のチャネルを監視
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("ch1から受信:", msg1)
		case msg2 := <-ch2:
			fmt.Println("ch2から受信:", msg2)
		}
	}
}

// --- 2. タイムアウト付きのチャネル待機 --- //
func main0102() {
	ch := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		ch <- "Hello!"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout!")
	}
}

// --- 3. for-select パターン --- //
func worker(done chan bool, dataCh chan int) {
	for {
		select {
		case <-done:
			fmt.Println("ワーカーを終了します")
			return
		case data := <-dataCh:
			// データを処理
			fmt.Printf("受信したデータ: %d\n", data)
			time.Sleep(1 * time.Second) // 処理時間をシミュレート
		}
	}
}
func main0103() {
	done := make(chan bool)
	dataCh := make(chan int)

	// ワーカーの起動
	go worker(done, dataCh)

	// データを送信
	for i := 1; i <= 3; i++ {
		dataCh <- i
	}

	// 終了シグナルを送信
	done <- true
	fmt.Println("メインプログラムを終了します")
}
