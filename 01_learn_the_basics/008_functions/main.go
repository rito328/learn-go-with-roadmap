package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {

}

func greet(name string) string {
	return "Hello, " + name
}
func sample00() {
	// 関数の呼び出し
	message := greet("Alice")
	fmt.Println(message)
	// 出力: Hello, Alice

	// 関数の戻り値を直接使用
	fmt.Println(greet("Bob"))
	// 出力: Hello, Bob
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
func sample01() {
	result, err := divide(1.0, 2.3)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
	// => 0.4347826086956522
}

func add(a, b int) (sum int) {
	sum = a + b
	// 明記していないが sum が return される
	return
}
func sample02() {
	result := add(1, 2)
	fmt.Println(result)
	// => 3
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
func sample03() {
	total := sum(1, 2, 3, 4, 5)
	fmt.Println(total)
	// => 15
}

func sample04() {
	add := func(a, b int) int {
		return a + b
	}
	result := add(3, 4)
	fmt.Println(result)
	// => 7
}

func incrementer() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
func sample05() {
	// incrementer で生成された関数を取得
	inc := incrementer()

	// 何度か呼び出してみる
	fmt.Println(inc()) // => 1
	fmt.Println(inc()) // => 2
	fmt.Println(inc()) // => 3

	// もう一度 incrementer から別の関数を取得
	anotherInc := incrementer()

	// 新しい関数は別のクロージャとして動作する
	fmt.Println(anotherInc()) // => 1
	fmt.Println(anotherInc()) // => 2
}

func adder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
func sample06() {
	// 10を足す関数を作成
	add10 := adder(10)

	// 20を足す関数を作成
	add20 := adder(20)

	fmt.Println(add10(5)) // => 15
	fmt.Println(add20(5)) // => 25
}

func sample07() {
	result := func(a, b int) int {
		return a + b
	}(3, 4) // 無名関数を定義すると同時に実行

	fmt.Println(result)
	// => 7
}

func sample08() {
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j] // 降順にソート
	})

	fmt.Println(numbers)
	// 出力: [9 6 5 5 4 3 3 2 1 1]
}

func sample09() {
	message := "Hello, World!"

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println(message)
	}()

	message = "Goodbye, World!"
	time.Sleep(3 * time.Second)
	// 2 秒後に "Goodbye, World!" が出力される
}

func filter(numbers []int, f func(int) bool) []int {
	var result []int
	for _, num := range numbers {
		if f(num) {
			result = append(result, num)
		}
	}
	return result
}
func sample10() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evenNumbers := filter(numbers, func(n int) bool {
		return n%2 == 0
	})

	fmt.Println(evenNumbers)
	// => [2 4 6 8 10]
}
