package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string // テスト名
		a, b     int    // テスト時に使用する値
		expected int    // 期待する結果
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -1, -2, -3},
		{"mixed numbers", -1, 5, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
