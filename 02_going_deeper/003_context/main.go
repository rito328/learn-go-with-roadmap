package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

}

// --- 1. キャンセル制御の例 --- //
func cancelExample() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("goroutine: キャンセルを検知しました")
			return
		}
	}()

	// 1秒後にキャンセル
	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(100 * time.Millisecond) // goroutineの完了を待つ
}
func main0101() {
	fmt.Println("=== キャンセルの例 ===")
	cancelExample()
}

// --- 2. タイムアウト制御の例 --- //
func timeoutExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("この行は実行されません")
	case <-ctx.Done():
		fmt.Println("timeout: タイムアウトしました")
	}
}
func main0102() {
	fmt.Println("=== タイムアウトの例 ===")
	timeoutExample()
}

// --- 3. 値の伝播の例 --- //
func valueExample() {
	type key string
	const userIDKey key = "userID"

	ctx := context.WithValue(context.Background(), userIDKey, "user123")

	// 値の取得
	if userID, ok := ctx.Value(userIDKey).(string); ok {
		fmt.Printf("value: ユーザーID %s を取得しました\n", userID)
	}
}
func main0103() {
	fmt.Println("=== 値の伝播の例 ===")
	valueExample()
}
