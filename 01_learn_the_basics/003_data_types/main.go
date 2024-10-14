package main

import (
	"fmt"
	"math/cmplx"
	"os"
	"reflect"
)

func main() {
}

// float32/64
func sample0101() {
	// 型宣言をせずに浮動小数点数を代入
	x := 3.14

	// 変数 x の型を確認
	fmt.Printf("xの値: %v\n", x)
	fmt.Printf("xの型: %v\n", reflect.TypeOf(x))

	// 明示的に float32 を指定した場合との比較
	y := float32(3.14)
	fmt.Printf("yの値: %v\n", y)
	fmt.Printf("yの型: %v\n", reflect.TypeOf(y))
}

func sample0102() {
	var value32 float32
	var value64 float64

	value32 = 1234.56789
	value64 = 1234.56789

	fmt.Printf("float32: %f\n", value32)
	fmt.Printf("float64: %f\n", value64)
}

// complex64/128
func sample0201() {
	// complex64 の複素数
	var c1 complex64 = complex(1.2, 3.4) // 実数部1.2, 虚数部3.4
	fmt.Printf("c1: %v (complex64)\n", c1)

	// complex128 の複素数
	var c2 complex128 = complex(1.2, 3.4) // 実数部1.2, 虚数部3.4
	fmt.Printf("c2: %v (complex128)\n", c2)

	// 実数部と虚数部を分離して取得
	fmt.Printf("実数部: %f, 虚数部: %f\n", real(c2), imag(c2))
}

func sample0202() {
	// complex64の使用例
	var c64 complex64 = 3 + 4i
	fmt.Printf("complex64: %v\n", c64)
	fmt.Printf("実部: %f, 虚部: %f\n", real(c64), imag(c64))

	// complex128 の使用例
	var c128 complex128 = 3 + 4i
	fmt.Printf("complex128: %v\n", c128)
	fmt.Printf("実部: %f, 虚部: %f\n", real(c128), imag(c128))

	// 複素数の演算
	fmt.Printf("絶対値: %f\n", cmplx.Abs(c128))
	fmt.Printf("指数関数: %v\n", cmplx.Exp(c128))

	// 精度の違いを示す例
	c64precise := complex64(0.1 + 0.2i)
	c128precise := complex128(0.1 + 0.2i)
	fmt.Printf("complex64精度: %v\n", c64precise*c64precise)
	fmt.Printf("complex128精度: %v\n", c128precise*c128precise)
}

// byte
func sample03() {
	// --- 文字列の処理 ---
	str := "Hello"
	bytes := []byte(str) // 文字列をバイトスライスに変換
	fmt.Println(bytes)   // [72 101 108 108 111]

	// --- バイナリデータの操作 ---
	// ファイルからバイナリデータを読み込む
	data, err := os.ReadFile("/path/to/01_learn_the_basics/003_data_types/sample.png") // os.ReadFile を使用
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(data) // バイトのスライスとして出力
}

// rune
func sample04() {
	str := "こんにちは"

	// 文字列を1文字ずつ（rune単位）で処理
	for i, r := range str {
		fmt.Printf("%d: %c (Unicode: %U)\n", i, r, r)
	}
}
