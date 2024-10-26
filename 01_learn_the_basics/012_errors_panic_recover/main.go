package main

import (
	"errors"
	"fmt"
)

func main() {

}

// -- Errors -- //
// -- Errors -- //
// -- Errors -- //
// -- Errors -- //

// --- エラーの作成

func DoSomething() error {
	return errors.New("何かがうまく動作しませんでした")
}
func sample010101() {
	err := DoSomething()
	if err != nil {
		fmt.Printf("エラーが発生しました %w", err)
	}
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("'%d' をゼロで除算することはできません", a)
	}
	return a / b, nil
}
func sample010102() {
	err := DoSomething()
	if err != nil {
		fmt.Printf("エラーが発生しました %w", err)
	}
}

// --- 基本的なパターン

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
func sample010201() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("エラーが発生:", err)
		return
	}
	fmt.Println("結果:", result)
}

// --- 想定されるエラーの定義

// ErrDivideByZero 想定されるエラーを定義
var ErrDivideByZero = errors.New("divide by zero")

func Divide2(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}
func sample010301() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("エラーが発生:", err)
		return
	}
	fmt.Println("結果:", result)
}

// --- カスタムエラー型の実装

// ResourceNotFoundError カスタムエラー型
type ResourceNotFoundError struct {
	Code    int
	Message string
}

func (e *ResourceNotFoundError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}
func sample010401() {
	err := &ResourceNotFoundError{
		Code:    404,
		Message: "Resource not found",
	}
	fmt.Println(err)
}

// --- エラーのラッピング

func add(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("両方 0 以上にしてください")
	}

	return a + b, nil
}
func addExecute() (int, error) {
	i, err := add(1, 0)
	if err != nil {
		// エラーをラップして、追加のコンテキストを与える
		return 0, fmt.Errorf("addExecute でエラー: %w", err)
	}
	return i, nil
}
func sample010501() {
	i, err := addExecute()
	if err != nil {
		// エラーを表示
		fmt.Println("エラー:", err)
		// => エラー: addExecute でエラー: 両方 0 以上にしてください
		return
	}
	fmt.Println(i)
}

// ---- errors.Unwrap
func sample010502() {
	i, err := addExecute()
	if err != nil {
		fmt.Println("エラー:", err)
		// => エラー: addExecute でエラー: 両方 0 以上にしてください

		// ラップされたエラーを解除して元のエラーを取得
		origErr := errors.Unwrap(err)
		fmt.Println("元のエラー:", origErr)
		// => 元のエラー: 両方 0 以上にしてください

		return
	}
	fmt.Println(i)
}

// ---- errors.Is と errors.As
var ErrLowLevel = errors.New("低レベルのエラーが発生しました")

func fuga() error {
	return ErrLowLevel
}
func hoge() error {
	err := fuga()
	if err != nil {
		return fmt.Errorf("hoge でエラー: %w", err)
	}
	return nil
}

func sample010503() {
	err := hoge()
	if err != nil {
		fmt.Println("エラー:", err)

		// エラーが特定のエラー（ErrLowLevel）かどうかをチェック
		if errors.Is(err, ErrLowLevel) {
			fmt.Println("ErrLowLevel が原因のエラーです")
		}
	}
}

type MyError2 struct {
	Message string
}

func (e *MyError2) Error() string {
	return e.Message
}
func DoSomething2() error {
	return &MyError2{"カスタムエラーが発生しました"}
}
func sample010504() {
	err := DoSomething2()

	var myErr *MyError2
	// エラーが MyError2 型かどうかを確認
	if errors.As(err, &myErr) {
		fmt.Println("カスタムエラー:", myErr)
	} else {
		fmt.Println("異なる型のエラー:", err)
	}
}

// -- Panic -- //
// -- Panic -- //
// -- Panic -- //
// -- Panic -- //
func sample0201() {
	fmt.Println("Start")
	panic("something went wrong")
	fmt.Println("End") // この行は実行されません
}

// ---- パニックの伝播と回復

// パニックの伝播例
func level3() {
	panic("level 3 panic")
}

func level2() {
	defer func() {
		fmt.Println("level 2 defer")
	}()
	level3()
}

func level1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered: %v\n", r)
		}
	}()
	level2()
}
func sample0202() {
	level1()
}

// -- Recover -- //
func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}
func riskyFunction() {
	defer handlePanic()

	// 何らかの処理
	panic("something went wrong")
}
func sample0301() {
	riskyFunction()
	fmt.Println("プログラムは続行します")
}

// -- エラーハンドリングのベストプラクティス -- //

// ErrNotFound 基本的なエラー定数
var ErrNotFound = errors.New("not found")

// MyError カスタムエラー型
type MyError struct {
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}

func getData() error {
	// ErrNotFound をラップして返す（errors.Is で補足する用）
	return fmt.Errorf("failed to get data: %w", ErrNotFound)

	// カスタムエラーを作成してラップ（errors.As で補足する用）
	//myErr := &MyError{Message: "something went wrong"}
	//return fmt.Errorf("failed to get data: %w", myErr)
}
func sample0401() {
	err := getData()
	if err != nil {
		// Is: エラー値の比較
		if errors.Is(err, ErrNotFound) {
			fmt.Println("見つかりませんでした")
			return
		}

		// As: エラー型の検査
		var myErr *MyError
		if errors.As(err, &myErr) {
			fmt.Printf("MyErrorです: %s\n", myErr.Message)
			return
		}

		fmt.Printf("その他のエラー: %v\n", err)
	}
}
