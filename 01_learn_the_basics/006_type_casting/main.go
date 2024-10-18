package main

import (
	"fmt"
	"strconv"
)

func main() {

}

func sample01() {
	var i int = 42
	var f float64 = float64(i)
	fmt.Printf("int: %d, float64: %f\n", i, f)
	// => int: 42, float64: 42.000000
}

func sample02() {
	var i interface{} = "Hello"
	// interface{} 型の変数 i に格納されている値が string 型であるかを確認
	s, ok := i.(string)
	if ok {
		fmt.Println("String value:", s)
	} else {
		fmt.Println("Value is not a string")
	}
}

func sample03() {
	s := "123"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
}

func sample04() {
	i := 123
	s := strconv.Itoa(i)
	fmt.Println(s)
}

func sample05() {
	var i interface{} = 42

	s := i.(string) // panic: interface conversion: interface {} is int, not string

	fmt.Println("String value:", s)
}

func sample06() {
	var i int = 12345678910111213
	var f float32 = float32(i)
	fmt.Printf("int: %d, float32: %f\n", i, f)
	// int: 12345678910111213, float32: 12345678407663616.000000
}

func sample07() {
	str := "Hello, Go!"
	bytes := []byte(str)
	fmt.Println(bytes)
	// => [72 101 108 108 111 44 32 71 111 33]
}

func sample08() {
	bytes := []byte{72, 101, 108, 108, 111, 44, 32, 71, 111, 33}
	str := string(bytes)
	fmt.Println(str)
	// => Hello, Go!
}

func sample09() {
	bytes := []byte("Hello, Go!")
	bytes[7] = 'g'
	fmt.Println(string(bytes))
	// => Hello, go!
}

func sample10() {
	str := "こんにちは"
	bytes := []byte(str)
	fmt.Println(bytes)
	// 各バイトのコードポイントが表示される
	// => [227 129 147 227 130 147 227 129 171 227 129 161 227 129 175]
	fmt.Println(len(str))
	// 文字列の長さ（文字数）: 5
	// => 15
	fmt.Println(len(bytes))
	// バイトスライスの長さ（バイト数）: 15
	// => 15
}

func sample11() {
	a := 5
	b := 2
	result := a / b
	fmt.Println(result)
	// => 2
}

func sample12() {
	a := 5.0
	b := 2.0
	result := a / b
	fmt.Println(result)
	// => 2.5
}

func sample13() {
	a := 5
	b := 2
	result := float64(a) / float64(b)
	fmt.Println(result)
	// => 2.5
}

func sample14() {
	a := 5                   // int
	b := 2.0                 // float64
	result := float64(a) / b // aをfloat64にキャスト
	fmt.Println(result)
	// => 2.5
}
