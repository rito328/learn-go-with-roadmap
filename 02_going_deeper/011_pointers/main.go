package main

import "fmt"

func main() {
}

// -- ポインターの基本概念 -- //
func main01() {
	// 通常の変数
	x := 10

	// ポインター変数の宣言 (*int は int型を指すポインター)
	var p *int

	// x のアドレスを p に格納
	p = &x

	// ポインターを通じて x の値を変更
	*p = 20

	fmt.Printf("x の値: %d\n", x)       // 20
	fmt.Printf("p の値(アドレス): %v\n", p) // メモリアドレス
	fmt.Printf("p の指す値: %d\n", *p)    // 20
}

// -- ポインターを使った値の変更 -- //
func main02() {
	x := 42
	p := &x
	*p = 100       // ポインターを通じて値を変更
	fmt.Println(x) // 100
}

// -- 関数とポインタ -- //
func updateValue(p *int) {
	*p = 99 // ポインターを使って値を更新
}
func main03() {
	x := 42
	updateValue(&x) // x のアドレスを渡す
	fmt.Println(x)  // 99
}

// -- 構造体とポインタ -- //
type Person struct {
	Name string
	Age  int
}

func (p *Person) Birthday() {
	p.Age++ // (*p).Age++ と同じ
}
func main04() {
	person := &Person{
		Name: "田中",
		Age:  25,
	}
	person.Birthday()
	fmt.Printf("年齢: %d\n", person.Age) // 26

}
