package main

import (
	"fmt"
	"github.com/rito328/20241009_rito_go_sample_project/09_testing_your_apps/mathutil"
)

func runTests() {
	fmt.Println("手動テスト実行開始")

	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 1, 2},
		{2, 3, 5},
		{-1, 1, 0},
	}

	for _, tt := range tests {
		result := mathutil.Add(tt.a, tt.b)
		if result != tt.expected {
			fmt.Printf("Test failed: Add(%d, %d) = %d; want %d\n", tt.a, tt.b, result, tt.expected)
		} else {
			fmt.Printf("Test passed: Add(%d, %d) = %d\n", tt.a, tt.b, result)
		}
	}

	fmt.Println("手動テスト実行完了")
}

func main() {
	fmt.Println("通常のアプリケーションの実行")
	fmt.Println("2 + 3 =", mathutil.Add(2, 3))

	runTests()
}
