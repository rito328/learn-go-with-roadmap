package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {

}

// -- Work Stealing -- //
func generateWork(id int) {
	// 重い計算を模擬
	sum := 0
	for i := 0; i < 1000000; i++ {
		sum += i
	}
	fmt.Printf("Work %d completed\n", id)
}

func main01() {
	// 2つのプロセッサのみを使用
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup

	// P1に多くのタスクを割り当て
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			generateWork(id)
		}(i)

		// 最初の5つのタスクは少し待って実行開始
		if i < 5 {
			time.Sleep(10 * time.Millisecond)
		}
	}

	// P2は初期状態では仕事が少ない
	for i := 10; i < 12; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			generateWork(id)
		}(i)
	}

	wg.Wait()
}

// -- システムコールの扱い -- //
func main02() {
	runtime.GOMAXPROCS(1) // 1つのPのみを使用

	var wg sync.WaitGroup

	// システムコールを実行するゴルーチン
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1: Starting file operation (syscall)")

		// システムコールの例（ファイル操作）
		f, err := os.Create("temp.txt")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer os.Remove("temp.txt")
		defer f.Close()

		// ファイル書き込み（システムコール）
		f.WriteString("Hello, World!")
		fmt.Println("Goroutine 1: File operation completed")
	}()

	// CPU負荷の高い処理を実行する他のゴルーチン
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d: Starting CPU-bound work\n", id+2)

			// CPU負荷の高い処理を模擬
			sum := 0
			for i := 0; i < 1000000; i++ {
				sum += i
			}

			fmt.Printf("Goroutine %d: Completed CPU-bound work\n", id+2)
		}(i)
	}

	wg.Wait()
}

// -- パフォーマンス面の特徴 -- //
func memStats() runtime.MemStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m
}
func main03() {
	// メモリ使用量を表示
	fmt.Println("Initial memory stats:")
	initial := memStats()
	fmt.Printf("Alloc: %v MiB\n", initial.Alloc/1024/1024)

	start := time.Now()
	var wg sync.WaitGroup

	// 100,000個のゴルーチンを作成
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// 軽い処理
			time.Sleep(time.Millisecond)
		}(i)
	}

	wg.Wait()
	duration := time.Since(start)

	// 終了時のメモリ使用量を表示
	final := memStats()
	fmt.Printf("\nFinal memory stats:\n")
	fmt.Printf("Alloc: %v MiB\n", final.Alloc/1024/1024)
	fmt.Printf("Total time: %v\n", duration)
	fmt.Printf("Goroutines created: 100,000\n")
	fmt.Printf("Average creation+execution time per goroutine: %v\n", duration/100000)
}
