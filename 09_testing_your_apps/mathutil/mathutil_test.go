package mathutil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var testSetup string // グローバル変数を使ってセットアップ情報を保持

// TestMain はすべてのテストの前後に実行される
func TestMain(m *testing.M) {
	fmt.Println("=== テストのセットアップ開始 ===")
	testSetup = "テスト環境が準備されました"

	// すべてのテストを実行
	exitCode := m.Run()

	fmt.Println("=== テストのクリーンアップ開始 ===")
	testSetup = "" // クリーンアップ処理

	// テストの終了コードを返す
	os.Exit(exitCode)
}

// `TestMain` でセットアップした値をテスト内で利用できる
func TestAdd(t *testing.T) {
	if testSetup == "" {
		t.Fatal("テストのセットアップが正しく実行されていません")
	}

	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	} else {
		t.Log("TestAdd PASSED")
	}
}

func TestAdd2(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	// アサーションを使用
	assert.Equal(t, expected, result, "Add(2, 3) should return 5")
}

func TestAddTableDriven(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 1, 2},
		{2, 3, 5},
		{-1, 1, 0},
		{0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Add(%d, %d)", tt.a, tt.b), func(t *testing.T) {
			result := Add(tt.a, tt.b)
			assert.Equal(t, tt.expected, result, "Add(%d, %d) should return %d", tt.a, tt.b, tt.expected)
		})
	}
}
