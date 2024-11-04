package main

import (
	"fmt"
	"time"
)

func main() {
}

// -- チャネルへの送信と受信 -- //
func main01() {
	ch := make(chan int) // バッファなしチャネル

	go func() {
		ch <- 10
	}()

	value := <-ch

	fmt.Printf("The value of ch is %d\n", value)
}

// -- バッファなしチャネル -- //
func main02() {
	ch := make(chan int) // バッファなしチャネル

	// 送信側の処理
	go func() {
		fmt.Println("送信準備完了")
		ch <- 42            // ここでブロックされる
		fmt.Println("送信完了") // 受信されるまでここには到達しない
	}()

	// メインルーチンで少し待機
	time.Sleep(2 * time.Second)

	// 受信側の処理
	fmt.Println("受信準備完了")
	value := <-ch // 受信処理
	fmt.Println("受信した値:", value)

	time.Sleep(1 * time.Second) // 送信完了のメッセージを確認するため待機
}

// -- バッファありチャネル -- //
func main03() {
	bufCh := make(chan int, 2)

	go func() {
		bufCh <- 1 // ブロックされない
		fmt.Println("1つ目の送信完了")

		bufCh <- 2 // ブロックされない
		fmt.Println("2つ目の送信完了")

		bufCh <- 3 // ブロックされる
		fmt.Println("3つ目の送信完了")
	}()

	time.Sleep(2 * time.Second)

	v := <-bufCh
	fmt.Println(v)

	time.Sleep(1 * time.Second)

	v = <-bufCh
	fmt.Println(v)

	time.Sleep(1 * time.Second)

	v = <-bufCh
	fmt.Println(v)

	time.Sleep(1 * time.Second)
}

// -- チャネルのクローズ -- //
func main04() {
	ch := make(chan int, 5)

	// データを送信するゴルーチン
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Printf("%d を送信しました\n", i)
		}
		close(ch) // チャネルをクローズ
		fmt.Println("チャネルをクローズしました")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("今から受信していきます")

	// チャネルからデータを受信し続ける
	for value := range ch {
		fmt.Println("受信:", value)
	}

	fmt.Println("受信が終了しました")
}

// -- Select 文 -- //
func main05() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 送信側ゴルーチン1
	go func() {
		time.Sleep(2 * time.Second) // 2 秒待機
		ch1 <- "one"
	}()

	// 送信側ゴルーチン2
	go func() {
		time.Sleep(1 * time.Second) // 1 秒待機（こちらが先に送信される）
		ch2 <- "two"
	}()

	// 受信側（select）
	select {
	case msg1 := <-ch1:
		fmt.Println("受信ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("受信ch2:", msg2)
	}
}

// --- 1. タイムアウト処理 --- //
func main0501() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "処理完了"
	}()

	select {
	case msg := <-ch:
		fmt.Println("受信:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("タイムアウト")
	}
}

// --- 2. デフォルトケース（ノンブロッキング操作）--- //
func main0502() {
	ch := make(chan string)

	select {
	case msg := <-ch:
		fmt.Println("受信:", msg)
	default:
		fmt.Println("データなし") // チャネルが空の場合すぐに実行
	}
}

// --- 3. 送受信の両方を待ち受ける --- //
func main0503() {
	ch := make(chan string) // メッセージ送信用チャネル
	done := make(chan bool) // 処理中断用チャネル

	// 受信側ゴルーチン（別のゴルーチンで待ち受ける人）
	go func() {
		time.Sleep(2 * time.Second) // 2秒後に受信する想定
		msg := <-ch
		fmt.Println("受信完了:", msg)
	}()

	// 送信用ゴルーチン
	go func() {
		time.Sleep(3 * time.Second) // 3秒後にメッセージ送信
		ch <- "hello"
	}()

	// 中断通知用ゴルーチン
	go func() {
		time.Sleep(1 * time.Second) // 1秒後に中断通知
		done <- true
	}()

	select {
	case msg := <-ch: // チャネルからの受信を待つ
		fmt.Printf("送信成功 %s", msg)
	case <-done: // 中断信号を受け取る
		fmt.Println("処理中断")
	}
}

// --- 4. 無限ループでの使用 --- //
func main0504() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	done := make(chan bool)

	// 送信側ゴルーチン
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			ch1 <- fmt.Sprintf("ch1: %d", i)
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(2 * time.Second)
			ch2 <- fmt.Sprintf("ch2: %d", i)
		}
		done <- true
	}()

	// 受信側
	finished := 0
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println("受信1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("受信2:", msg2)
		case <-done:
			finished++
			if finished == 2 {
				fmt.Println("全ての処理が完了")
				return
			}
		}
	}
}
