package main

import (
	"fmt"
	"reflect"
)

func main() {

}

func sample01() {
	var a string = "A"
	fmt.Printf("Variable 'a' is of type %T\n", a)

	var b = "B"
	fmt.Printf("Variable 'b' is of type %T\n", b)

	// 別の型の値は代入不可
	//b = 1
	// => cannot use 1 (untyped int constant) as string value in assignment
}

func sample02() {
	a := "A" // string型と推論される
	fmt.Printf("Variable 'a' is of type %T\n", a)
}

func sample03() {
	i := 42      // 型は int と推論
	f := 3.14    // 型は float64 と推論
	s := "hello" // 型は string と推論
	b := true    // 型は bool と推論

	fmt.Printf("i: %T, f: %T, s: %T, b: %T\n", i, f, s, b)
}

func sample04() {
	numbers := []int{1, 2, 3, 4, 5}   // 型は []int と推論
	names := []string{"Alice", "Bob"} // 型は []string と推論

	fmt.Printf("numbers: %T, names: %T\n", numbers, names)
	// => numbers: []int, names: []string
}

func sample05() {
	num := getNumber() // 型は int と推論

	fmt.Printf("num: %T\n", num)
	// => num: int
}
func getNumber() int {
	return 10
}

func sample06() {
	const x = 5

	var y int = x     // x は int として推論される
	var z float64 = x // 同じ x でも文脈によって float64 として扱われる

	fmt.Printf("y: %T, z: %T\n", y, z)
	// => y: int, z: float64
}

func sample07() {
	const e int = 5 // e は int 型として固定される
	//var f float64 = e // コンパイルエラー: int から float64 への暗黙の変換は許可されない
	// =>  cannot use e (constant 5 of type int) as float64 value in variable declaration
}

// --- 構造体（struct）の型推論 --- //

func sample08() {
	// 明示的な型宣言
	var p1 Person = Person{Name: "Alice", Age: 30}

	// 型推論を使用
	p2 := Person{Name: "Bob", Age: 25}

	fmt.Printf("p1: %s, %+v\n", reflect.TypeOf(p1).Name(), p1)
	// => p1: Person, {Name:Alice Age:30}
	fmt.Printf("p2: %s, %+v\n", reflect.TypeOf(p2).Name(), p2)
	// => p2: Person, {Name:Bob Age:25}
}

type Person struct {
	Name string
	Age  int
}

// --- マップ（map）の型推論 --- //

func sample09() {
	// 明示的な型宣言
	var m1 map[string]int = map[string]int{"a": 1, "b": 2}

	// 型推論を使用
	m2 := map[string]int{"x": 10, "y": 20}

	fmt.Printf("m1: %v\n", m1)
	// => m1: map[a:1 b:2]
	fmt.Printf("m2: %v\n", m2)
	// => m2: map[x:10 y:20]
}

// --- インターフェースの型推論 --- //

// Writer インターフェースを定義
type Writer interface {
	Write([]byte) (int, error)
}

// ConsoleWriter 構造体を定義
type ConsoleWriter struct{}

// ConsoleWriter に Write メソッドを実装
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
func sample10() {
	// ConsoleWriter のインスタンスを作成
	cw := ConsoleWriter{}

	// Writer インターフェースとして使用
	var w Writer = cw

	w.Write([]byte("Hello, Interface Type Inference!"))
}
