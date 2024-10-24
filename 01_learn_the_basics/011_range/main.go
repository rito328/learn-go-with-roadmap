package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

}

// -- 配列に対する range -- //

func sample0101() {
	numbers := []int{1, 2, 3, 4, 5}
	for i, v := range numbers {
		fmt.Printf("インデックス: %d, 値: %d\n", i, v)
	}
}

func sample0102() {
	numbers := []int{1, 2, 3, 4, 5}
	for _, v := range numbers {
		fmt.Printf("値: %d\n", v)
	}
}

func sample0103() {
	numbers := []int{1, 2, 3, 4, 5}
	for i := range numbers {
		fmt.Printf("インデックス: %d\n", i)
	}
}

func sample0104() {
	numbers := []int{1, 2, 3, 4, 5}
	for _, v := range numbers {
		v += 10 // この操作は元のスライスに影響を与えません
	}
	fmt.Println(numbers) // 出力: [1 2 3 4 5]

	for i := range numbers {
		numbers[i] += 10 // 直接スライスの要素を変更
	}
	fmt.Println(numbers) // 出力: [11 12 13 14 15]

}

// -- マップに対する range -- //

func sample0301() {
	ages := map[string]int{"Alice": 25, "Bob": 30}
	for k, v := range ages {
		fmt.Printf("キー: %s, 値: %d\n", k, v)
	}
}

func sample0302() {
	ages := map[string]int{"Alice": 25, "Bob": 30}
	for k := range ages {
		fmt.Printf("キー: %s\n", k)
	}
}

// -- 文字列に対する range -- //

func sample0401() {
	// ASCII文字（1バイト）の場合
	word1 := "Hello"
	for i, r := range word1 {
		fmt.Printf("バイト位置: %d, ルーン: %c, バイト数: %d\n", i, r, utf8.RuneLen(r))
	}

	fmt.Println()

	// 日本語（3バイト）の場合
	word2 := "こんにちは"
	for i, r := range word2 {
		fmt.Printf("バイト位置: %d, ルーン: %c, バイト数: %d\n", i, r, utf8.RuneLen(r))
	}
}

func sample0402() {
	word := "こんにちは"
	for i, r := range word {
		fmt.Printf("インデックス: %d, ルーン: %U\n", i, r)
	}
}

// -- チャネルに対する range -- //

func sample0501() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}
