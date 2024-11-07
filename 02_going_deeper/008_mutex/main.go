package main

import (
	"fmt"
	"sync"
)

// 使用例
func main() {
	main02()
}

// -- Mutex 使用例 -- //
var (
	counter int
	mu      sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done() // ゴルーチン終了時にWaitGroupのカウントを減らす
	mu.Lock()       // ロックを取得
	counter++       // カウンタをインクリメント
	mu.Unlock()     // ロックを解放
}
func main01() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) // ゴルーチンの数を追加
		go increment(&wg)
	}

	wg.Wait() // 全てのゴルーチンが終了するのを待つ
	fmt.Println("Final counter:", counter)
}

// -- ベストプラクティス -- //
// 良い例：カプセル化された構造体
type SafeCounter struct {
	mu     sync.Mutex
	counts map[string]int
}

// NewSafeCounter はカウンターを適切に初期化する
func NewSafeCounter() *SafeCounter {
	return &SafeCounter{
		counts: make(map[string]int),
	}
}

// Increment は安全にカウントを増やす
func (sc *SafeCounter) Increment(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.counts[key]++
}

// GetCount は安全にカウントを取得する
func (sc *SafeCounter) GetCount(key string) int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.counts[key]
}

func main02() {
	counter := NewSafeCounter()

	var wg sync.WaitGroup

	wg.Add(2)
	// 複数のゴルーチンから安全に使用可能
	go func() {
		defer wg.Done()
		counter.Increment("x")
	}()
	go func() {
		defer wg.Done()
		counter.Increment("x")
	}()

	wg.Wait()

	count := counter.GetCount("x")

	fmt.Printf("x: %d\n", count)
}

// -- RWMutex の使用例 -- //
var (
	data    int
	rwMutex sync.RWMutex
)

func read(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	rwMutex.RLock() // 読み取りロックを取得
	fmt.Printf("Reader %d: Data is %d\n", id, data)
	rwMutex.RUnlock() // 読み取りロックを解放
}

func write(wg *sync.WaitGroup, value int) {
	defer wg.Done()
	rwMutex.Lock() // 書き込みロックを取得
	data = value
	fmt.Printf("Writer: Wrote data %d\n", data)
	rwMutex.Unlock() // 書き込みロックを解放
}
func main03() {
	var wg sync.WaitGroup

	// 複数の読み取りゴルーチンを起動
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go read(&wg, i)
	}

	// 書き込みゴルーチンを起動
	wg.Add(1)
	go write(&wg, 42)

	// 再び読み取りゴルーチンを起動
	for i := 4; i <= 6; i++ {
		wg.Add(1)
		go read(&wg, i)
	}

	wg.Wait()
}
