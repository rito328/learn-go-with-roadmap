package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {

}

// ---- if --- //

// 基本的な構文
func sample0101() {
	x := 15
	if x > 10 {
		fmt.Println("x is greater than 10")
	}
}

// else if と else の使用
func sample0102() {
	x := 7
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else if x == 10 {
		fmt.Println("x is equal to 10")
	} else {
		fmt.Println("x is less than 10")
	}
}

// 変数の宣言と使用
func sample0103() {
	if y := 20; y > 10 {
		fmt.Println("y is greater than 10")
	}
}

// 複数の条件
func sample0104() {
	a := 5
	b := 8
	if a < 10 && b > 5 {
		fmt.Println("Both conditions are true")
	}
}

func sample0105() {
	a := 5
	b := 0
	if a == 0 || b == 0 {
		fmt.Println("a or b is 0.")
	}
}

// if 文でのエラーハンドリングと不要な else 文の省略
func checkNumber(n int) error {
	if n < 0 {
		return errors.New("negative number not allowed")
	}
	if n == 0 {
		return errors.New("zero is not a valid number")
	}
	fmt.Println("Valid number:", n)
	return nil
}
func sample0106() {
	n := 10
	err := checkNumber(n)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(n)
}

// --- switch...case 文 --- //

// 基本的な構文
func sample0201() {
	value := 1

	switch value {
	case 1:
		fmt.Println("Value is 1")
	case 2:
		fmt.Println("Value is 2")
	default:
		fmt.Println("Value is neither 1 nor 2")
	}
}

// 複数の条件
func sample0202() {
	value := 4

	switch value {
	case 1, 3, 5:
		fmt.Println("Value is odd")
	case 2, 4, 6:
		fmt.Println("Value is even")
	default:
		fmt.Println("Value is unknown")
	}
}

// 式を使用した case
func sample0203() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("午前中です")
	case t.Hour() < 17:
		fmt.Println("午後です")
	case t.Hour() < 19:
		fmt.Println("夕方です")
	default:
		fmt.Println("夜です")
	}
}

// 型に対する switch
func sample0204() {
	var x interface{} = "Hello"
	switch v := x.(type) {
	case int:
		fmt.Println("x is an int:", v)
	case string:
		fmt.Println("x is a string:", v)
	default:
		fmt.Println("x is of a different type")
	}
}

// fallthrough キーワード
func sample0205() {
	switch value := 2; value {
	case 1:
		fmt.Println("Case 1")
	case 2:
		fmt.Println("Case 2")
		fallthrough
	case 3:
		fmt.Println("Case 3")
	default:
		fmt.Println("Default case")
	}
	// => Case 2
	// => Case 3
}

// fallthrough 特徴と注意点
func sample0206() {
	switch value := 2; value {
	case 1:
		fmt.Println("Case 1")
	case 2:
		fmt.Println("Case 2")
		fallthrough
	case 4:
		fmt.Println("Case 4") // value は 4 ではないが、fallthrough で実行される
	default:
		fmt.Println("Default case")
	}
	// => Case 2
	// => Case 4
}

// fallthrough 使いどころ例
func sample0207() {
	rank := 2
	switch rank {
	case 1:
		fmt.Println("Rank 1 benefits")
		fallthrough
	case 2:
		fmt.Println("Rank 2 benefits")
		fallthrough
	case 3:
		fmt.Println("Rank 3 benefits")
	}
	// => Rank 2 benefits
	// => Rank 3 benefits
}
