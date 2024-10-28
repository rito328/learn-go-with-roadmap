package main

import "fmt"

func main() {
}

func sample0101() {
	var i interface{} = "hello"

	// 型アサーションを使用して、i が string 型であることを確認し、値を取得する
	s := i.(int)
	fmt.Println(s) // "hello" と出力される
}

func sample0102() {
	var i interface{} = 123

	s, ok := i.(string)
	if ok {
		fmt.Println("string型です:", s)
	} else {
		fmt.Println("string型ではありません")
	}
}

func doSomething(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("整数:", v)
	case string:
		fmt.Println("文字列:", v)
	default:
		fmt.Println("不明な型")
	}
}
func sample0103() {
	var i interface{} = "hello"
	doSomething(i)
}

// 任意の値を文字列に変換する関数
func toString(v interface{}) (string, error) {
	switch val := v.(type) {
	case string:
		return val, nil
	case int:
		return fmt.Sprintf("%d", val), nil
	case float64:
		return fmt.Sprintf("%.2f", val), nil
	case bool:
		return fmt.Sprintf("%t", val), nil
	default:
		return "", fmt.Errorf("unsupported type: %T", v)
	}
}
func sample0104() {
	values := []interface{}{
		"Hello",
		42,
		3.14,
		true,
	}

	for _, v := range values {
		result, err := toString(v)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("値: %v, 変換後: %s, 型: %T\n", v, result, v)
		// 値: Hello, 変換後: Hello, 型: string
		// 値: 42, 変換後: 42, 型: int
		// 値: 3.14, 変換後: 3.14, 型: float64
		// 値: true, 変換後: true, 型: bool
	}

	// interface{} 型の値から map[string]interface{} への型アサーション
	var data interface{} = map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	if m, ok := data.(map[string]interface{}); ok {
		// m["name"] を string として取り出せた場合（ok = true）
		if name, ok := m["name"].(string); ok {
			fmt.Printf("名前: %s\n", name)
			// 名前: John
		}
		// m["age"] を int として取り出せた場合（ok = true）
		if age, ok := m["age"].(int); ok {
			fmt.Printf("年齢: %d\n", age)
			// 年齢: 30
		}
	}
}
