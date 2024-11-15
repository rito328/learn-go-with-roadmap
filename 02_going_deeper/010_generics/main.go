package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func main() {
	main01()
}

// -- インターフェースを使用した型制約 -- //

// Number 数値型を表す型制約インターフェース
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

// Sum 数値のスライスの合計を計算する関数
func Sum[T Number](numbers []T) T {
	var sum T
	for _, n := range numbers {
		sum += n
	}
	return sum
}
func main01() {
	ints := []int{1, 2, 3, 4, 5}
	floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	fmt.Println(Sum(ints))   // 出力: 15
	fmt.Println(Sum(floats)) // 出力: 16.5
}

// -- ジェネリックな関数の例 -- //

// Max ジェネリックな関数の定義
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
func main02() {
	// int 型で Max 関数を使用
	intA, intB := 3, 5
	fmt.Printf("Max of %d and %d is %d\n", intA, intB, Max(intA, intB))

	// float64 型で Max 関数を使用
	floatA, floatB := 4.5, 2.3
	fmt.Printf("Max of %.1f and %.1f is %.1f\n", floatA, floatB, Max(floatA, floatB))

	// string 型で Max 関数を使用
	strA, strB := "apple", "banana"
	fmt.Printf("Max of %s and %s is %s\n", strA, strB, Max(strA, strB))
}

// -- ジェネリックな構造体の例 -- //

// Stack ジェネリックなスタック構造体の定義
type Stack[T any] struct {
	items []T
}

// Push - スタックに要素を追加
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop - スタックから要素を取り出し、削除
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false // スタックが空の場合
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}
func main03() {
	// int 型のスタックを作成
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)

	fmt.Println("Integer Stack:")
	for {
		item, ok := intStack.Pop()
		if !ok {
			break // スタックが空になった場合
		}
		fmt.Println(item)
	}

	// string 型のスタックを作成
	stringStack := Stack[string]{}
	stringStack.Push("Hello")
	stringStack.Push("World")

	fmt.Println("\nString Stack:")
	for {
		item, ok := stringStack.Pop()
		if !ok {
			break // スタックが空になった場合
		}
		fmt.Println(item)
	}
}

// --  -- //
func main04() {

}
