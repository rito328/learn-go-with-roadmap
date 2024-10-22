package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"
)

func main() {
}

// --- 基本的な for 文 --- //

func sample01() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

// --- 条件付きループ --- //

func sample02() {
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
}

// --- 無限ループ --- //

func sample03() {
	for {
		fmt.Println("無限ループ")
		// break で抜けない限り永久に実行
		break
	}
}

// --- break - ループの終了 --- //
func sample0400() {
	for i := 1; i <= 3; i++ {
		fmt.Printf("外側のループ: i = %d\n", i)
		for j := 1; j <= 3; j++ {
			if j == 3 {
				break // j が 3 になったら内側のループのみを抜ける
			}
			fmt.Printf("  内側のループ: j = %d\n", j)
		}
	}
	// 出力:
	// 外側のループ: i = 1
	//   内側のループ: j = 1
	//   内側のループ: j = 2
	// 外側のループ: i = 2
	//   内側のループ: j = 1
	//   内側のループ: j = 2
	// 外側のループ: i = 3
	//   内側のループ: j = 1
	//   内側のループ: j = 2
}

// --- ラベル付きの break --- //
func sample0401() {
outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j > 5 {
				break outer // 外側のループまで一気に抜ける
			}
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}
	// 出力:
	// i: 1, j: 1
	// i: 1, j: 2
	// i: 1, j: 3
	// i: 2, j: 1
	// i: 2, j: 2
}

// --- continue - 次のループへのスキップ --- //
func sample0500() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue // i が 3 の時はスキップ
		}
		fmt.Println(i)
	}
	// 出力:
	// 1
	// 2
	// 4
	// 5
}

// --- ラベル付きの continue --- //
func sample0501() {
outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j >= 4 {
				continue outer // 外側のループの次のイテレーションに進む
			}
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}
	// 出力:
	// i: 1, j: 1
	// i: 1, j: 2
	// i: 1, j: 3
	// i: 2, j: 1
	// i: 3, j: 1
}

// --- return - 関数の終了 --- //
func printNumbers() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			return // i が 3 になったら関数を終了
		}
		fmt.Println(i)
	}
	fmt.Println("この行は実行されません")
}
func sample06() {
	printNumbers()
}

// --- 無限ループ --- //

// --- ユーザー入力の待ち受け --- //
func sample0701() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("コマンドを入力してください（終了する場合は 'quit'）: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "quit" {
			fmt.Println("プログラムを終了します")
			break
		}

		fmt.Printf("入力されたコマンド: %s\n", input)
	}
}

// --- サーバープログラム --- //
func sample0702() {
	// シグナル待ち受けのチャネル
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// HTTPサーバーの設定
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// サーバーを別のゴルーチンで起動
	go func() {
		fmt.Println("サーバーを起動しました。http://localhost:8080/")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Printf("サーバーエラー: %v\n", err)
		}
	}()

	// シグナルを待ち受け
	for {
		sig := <-sigChan
		fmt.Printf("\n%vシグナルを受信しました。シャットダウンを開始します...\n", sig)
		break
	}
}

// --- 定期的なタスク実行 --- //
func sample0703() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	count := 0
	for {
		<-ticker.C
		count++
		fmt.Printf("タスクを実行しました（%d回目）\n", count)

		if count >= 5 {
			fmt.Println("すべてのタスクが完了しました")
			break
		}
	}
}

// --- 無限ループの危険性と注意点 --- //
// --- リソースの枯渇 --- //
func badLoop() {
	count := 0
	for {
		count++ // カウンターが無限に増加
	}
}
func goodLoop() {
	count := 0
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			count++
			fmt.Printf("現在のカウント: %d\n", count)
		}
	}
}
func sample0704() {
	// CPUの使用率を表示
	go func() {
		for {
			fmt.Printf("CPUの使用率: %d%%\n", runtime.NumGoroutine())
			time.Sleep(time.Second)
		}
	}()

	// badLoop() // このループを実行するとCPU使用率が100%になる
	goodLoop() // このループは適切な間隔で実行される
}

// --- タイムアウト機構例 --- //
func sample0705() {
	// タイムアウト付きの無限ループ
	timeout := time.After(5 * time.Second)
	count := 0

	for {
		select {
		case <-timeout:
			fmt.Println("タイムアウトしました")
			return
		default:
			count++
			fmt.Printf("処理回数: %d\n", count)
			time.Sleep(1 * time.Second)
		}
	}
}

// --- エラー処理の実装 --- //
func riskyOperation() error {
	// エラーが発生する可能性のある処理
	return fmt.Errorf("エラーが発生しました")
}
func sample0706() {
	maxRetries := 3
	retryCount := 0

	for {
		err := riskyOperation()
		if err != nil {
			retryCount++
			fmt.Printf("エラーが発生しました（%d回目）: %v\n", retryCount, err)

			if retryCount >= maxRetries {
				fmt.Println("最大リトライ回数に達しました。プログラムを終了します")
				break
			}

			// 一定時間待機してから再試行
			time.Sleep(time.Second)
			continue
		}

		// 正常終了
		fmt.Println("処理が成功しました")
		break
	}
}
