package main

import "fmt"

// パッケージレベルの変数
var globalVar int = 10

func main() {
	// main 関数内のローカル変数
	var localVar int = 20

	fmt.Println("グローバル変数:", globalVar)
	fmt.Println("main 関数のローカル変数:", localVar)

	// 他の関数を呼び出す
	printGlobal()
	tryToAccessLocal()

	// ローカル変数の値を変更
	localVar = 30
	fmt.Println("変更後の main 関数のローカル変数:", localVar)
}

func printGlobal() {
	fmt.Println("printGlobal 関数内からのグローバル変数:", globalVar)
	// ここでlocalVarにアクセスしようとするとコンパイルエラーになります
	// fmt.Println(localVar) // この行をコメントアウトを外すとエラーになります
}

func tryToAccessLocal() {
	fmt.Println("tryToAccessLocal 関数内からのグローバル変数:", globalVar)
	// この関数内で新しいローカル変数を定義
	var localVar int = 100 // これはmain 関数の localVar とは別の変数です
	fmt.Println("tryToAccessLocal関数内の新しいローカル変数:", localVar)
}
