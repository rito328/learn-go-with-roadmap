package main

import (
	"errors"
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

type Address struct {
	City  string
	State string
}

type Employee struct {
	Person
	Position string
	Address
}

func main() {

}

// --- array --- //
func sample01() {
	// var を使った配列の宣言（ゼロ値で初期化）
	var arr1 [3]int
	fmt.Println("var で宣言した配列 (初期状態):", arr1)

	// 要素を代入
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	fmt.Println("var で宣言した配列 (代入後):", arr1)

	// 配列の宣言と初期化
	arr := [3]int{1, 2, 3}
	fmt.Println("配列:", arr)

	// 配列の長さを取得
	fmt.Println("配列の長さ:", len(arr))

	// インデックスを使って要素にアクセス
	fmt.Println("インデックス 0 の要素:", arr[0])

	// 配列のコピー
	arrCopy := arr
	arrCopy[0] = 100
	fmt.Println("元の配列:", arr) // コピーされるため、元の配列は変更されない
	fmt.Println("コピーされた配列:", arrCopy)

	// 配列を関数に渡す
	modifyArray(arr)
	fmt.Println("関数呼び出し後の元の配列:", arr) // コピーされるため、元の配列は変更されない

	//// すべての要素は同じ型でなければならない
	//arr := [3]int{1, "2"}
	//// => cannot use "2" (untyped string constant) as int value in array or slice literal

	//arr := [3]int{1, 2, 3}
	//// 範囲外のインデックスにアクセスするとランタイムパニックが発生
	//fmt.Println("範囲外のインデックス 5 の要素:", arr[5])
	//// => invalid argument: index 5 out of bounds [0:3]

	//// -- コンパイルエラーの例 --
	//arr := [3]int{1, 2, 3}
	//// 範囲外のインデックスにアクセスするとコンパイルエラーが発生
	//fmt.Println("範囲外のインデックス 5 の要素:", arr[5])
	//// => invalid argument: index 5 out of bounds [0:3]
	//
	//// -- ランタイムパニックの例 --
	//var arr [3]int
	//index := 3 // 動的にインデックスを指定
	//fmt.Println("動的インデックス:", index)
	//
	//// 動的なインデックスによりランタイムパニックが発生
	//fmt.Println(arr[index])
	//// => panic: runtime error: index out of range [3] with length 3
}

// 配列を受け取る関数（配列全体がコピーされる）
func modifyArray(a [3]int) {
	a[0] = 999
	fmt.Println("関数内での配列:", a)
}

// --- slice --- //
func sample0201() {
	// --- リテラルを使ったスライスの宣言 --- //
	s := []int{1, 2, 3, 4, 5} // 配列を作らずに直接スライスを宣言
	fmt.Println(s)            // [1 2 3 4 5]

	// --- 配列からスライスを作成 --- //
	arr := [5]int{1, 2, 3, 4, 5}
	s1 := arr[1:4]  // 配列の一部をスライスとして取り出す
	fmt.Println(s1) // [2 3 4]

	// -- make 関数を使ったスライスの宣言 -- //
	s2 := make([]int, 5) // 長さ 5, 要素が全て 0 のスライスを作成
	fmt.Println(s2)      // [0 0 0 0 0]

	// --- append 関数を使った要素の追加 --- //
	s3 := []int{1, 2, 3}
	s3 = append(s3, 4, 5) // スライスに新しい要素を追加
	fmt.Println(s3)       // [1 2 3 4 5]

	// -- copy 関数を使ってスライスをコピー -- //
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)   // srcの内容をdstにコピー
	fmt.Println(dst) // [1 2 3]
}

func sample0202() {
	// -- スライスの拡張に関する例 -- //
	s := make([]int, 3, 5)     // 長さ3、容量5のスライスを作成
	fmt.Println("初期スライス:", s)  // [0 0 0]
	fmt.Println("長さ:", len(s)) // 3
	fmt.Println("容量:", cap(s)) // 5

	// 容量内で要素を追加する
	s = append(s, 1, 2)              // 容量を超えない範囲での追加
	fmt.Println("容量内での追加後のスライス:", s) // [0 0 0 1 2]
	fmt.Println("長さ:", len(s))       // 5
	fmt.Println("容量:", cap(s))       // 5

	// 容量を超える要素を追加する
	s = append(s, 3)                  // 容量を超えると新しい配列が確保される
	fmt.Println("容量を超えた追加後のスライス:", s) // [0 0 0 1 2 3]
	fmt.Println("長さ:", len(s))        // 6
	fmt.Println("容量:", cap(s))        // 10（新しい配列が確保され、容量が拡大）

	// -- スライスは参照型に関する例 -- //
	arr := [5]int{1, 2, 3, 4, 5} // 配列を作成
	s1 := arr[1:4]               // 配列の部分をスライスとして取り出す
	fmt.Println("元の配列:", arr)    // [1 2 3 4 5]
	fmt.Println("スライス s1:", s1)  // [2 3 4]

	s1[0] = 100                      // スライスの要素を変更
	fmt.Println("スライス s1 の変更後:", s1) // [100 3 4]
	fmt.Println("変更後の元の配列:", arr)    // [1 100 3 4 5] （元の配列にも影響）
}

// --- map --- //
func sample0301() {
	// map[キーの型]値の型 の形式で宣言します
	var m map[string]int

	// または、make関数を使用して初期化します
	m = make(map[string]int)

	// マップリテラルを使用して宣言と初期化を同時に行うこともできます
	m1 := map[string]int{
		"apple":  1,
		"banana": 2,
	}

	fmt.Println(m, m1)
}

func sample0302() {
	var m map[string]int

	if m == nil {
		fmt.Println("m is nil")
	} else {
		fmt.Println("m is not nil")
	}

	// => m is nil

	var m1 = make(map[string]int)

	if m1 == nil {
		fmt.Println("m1 is nil")
	} else {
		fmt.Println("m1 is not nil")
	}

	// => m1 is nil
}

func sample0303() {
	var m map[string]int
	m["key"] = 1 // => panic: assignment to entry in nil map

	var m1 = make(map[string]int)
	m1["key"] = 1

	var m2 = map[string]int{
		"key": 1,
	}

	fmt.Println(m1, m2)
	// => map[key:1] map[key:1]
}

func sample0304() {
	ages := make(map[string]int)

	// キーと値の追加
	ages["Alice"] = 25
	ages["Bob"] = 30

	fmt.Println(ages)
	// => map[Alice:25 Bob:30]

	// 値の取得
	aliceAge := ages["Alice"]
	fmt.Println(aliceAge)

	// 値の更新
	ages["Alice"] = 26

	// 値の削除
	delete(ages, "Bob")

	fmt.Println(ages)
	// => map[Alice:26]

	// キーが存在するかを確認する
	age, exists := ages["Alice"]
	if exists {
		fmt.Println("Aliceの年齢:", age)
	} else {
		fmt.Println("Aliceは存在しません")
	}
}

func sample0305() {
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}

	for key, value := range ages {
		fmt.Println(key, value)
		// => Bob 30
		// => Alice 25
	}
}

func sample0306() {
	a := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}

	b := a

	a["Alice"] = 26

	fmt.Println(a, b)
	// => map[Alice:26 Bob:30] map[Alice:26 Bob:30]
}

// --- 構造体 struct --- //

func sample0401() {
	// Name 以外のフィールドは初期化されない
	p := Person{
		Name: "Alice",
	}

	// フィールドのゼロ値が自動的に適用される
	fmt.Printf("%#v\n", p)
	// => main.Person{Name:"Alice", Age:0, Email:"", Active:false}
}

func sample0402() {
	// --- 構造体のインスタンス化 --- //
	p := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}
	fmt.Println(p) // {Alice 30 alice@example.com}

	// --- フィールドへのアクセス --- //
	fmt.Println(p.Name) // Alice
	fmt.Println(p.Age)  // 30

	// --- 構造体のフィールドの変更 --- //
	p.Age = 31
	fmt.Println(p.Age) // 31
}

func sample0403() {
	p := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}

	updateEmail(&p, "new_email@example.com")
	fmt.Println(p.Email) // new_email@example.com
}

func updateEmail(p *Person, newEmail string) {
	p.Email = newEmail
}

func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

func sample0404() {
	p := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}

	fmt.Println(p.Greet()) // Hello, my name is Alice
}

func newPerson(name string, age int, email string) (*Person, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if age < 0 {
		return nil, errors.New("age cannot be negative")
	}
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	return &Person{
		Name:  name,
		Age:   age,
		Email: email,
	}, nil
}

func sample0405() {
	p, err := newPerson("John Doe", 18, "johndoe@example.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Name: %s, Age: %d, Email: %s\n", p.Name, p.Age, p.Email)
	// => Name: John Doe, Age: 18, Email: johndoe@example.com
}

func sample0406() {
	e := Employee{
		Person: Person{
			Name:  "Bob",
			Age:   40,
			Email: "bob@example.com",
		},
		Position: "Manager",
		Address: Address{
			City:  "New York",
			State: "NY",
		},
	}

	fmt.Println(e.Name) // Bob（Person構造体のフィールドに直接アクセス可能）
	fmt.Println(e.City) // New York（Address構造体のフィールドに直接アクセス可能）
}

func sample0407() {
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}

// --- make --- //
func sample05() {
	// --- スライスの生成 --- //
	s := make([]int, 3, 10)

	fmt.Println(s)
	// => [0 0 0]

	// --- マップの生成 --- //
	m := make(map[string]int, 5)

	fmt.Println(m)
	// => map[]
	fmt.Println(len(m))
	// => 0

	// --- チャネルの生成 --- //
	ch := make(chan int, 3)

	fmt.Println(ch)
	// => 0xc000120000 (メモリアドレス。実行ごとに異なる）
	fmt.Println(cap(ch))
	// => 3

	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	// => 1
}
