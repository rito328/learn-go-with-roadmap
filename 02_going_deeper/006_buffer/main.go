package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

func main() {
	main03()
}

// -- bytes.Buffer -- //
func main01() {
	var buffer bytes.Buffer
	buffer.WriteString("Hello, ")
	buffer.WriteString("Go!")
	fmt.Println(buffer.String()) // 出力: Hello, Go!
}

// -- bufio パッケージ -- //
func main02() {
	// ファイル読み込みの例
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	if file != nil {
		defer file.Close()
	}

	// カスタムサイズ（8KB）のバッファ付きリーダーの作成
	reader := bufio.NewReaderSize(file, 8192)

	// 1行ずつ読み込み
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(line)
	}
}

// -- バッファプール -- //
func main03() {
	var bufferPool = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}

	// バッファをプールから取得
	buffer := bufferPool.Get().(*bytes.Buffer)
	defer func() {
		buffer.Reset()
		bufferPool.Put(buffer)
	}()

	// バッファを使ってデータを書き込み
	buffer.WriteString("Hello, Go with Buffer Pool!")
	fmt.Println(buffer.String())

	// バッファをクリアしてプールに戻す
	buffer.Reset()
	bufferPool.Put(buffer)

	// 再利用可能なバッファをプールから再取得
	buffer2 := bufferPool.Get().(*bytes.Buffer)
	buffer2.WriteString("Reused Buffer!")
	fmt.Println(buffer2.String())
}
