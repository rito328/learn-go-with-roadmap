package main

import (
	"fmt"
	"os"
)

// Storage データの保存を抽象化するインターフェース
type Storage interface {
	Save(data []byte) error
	Load() ([]byte, error)
}

// FileStorage ファイルへの実装
type FileStorage struct {
	path string
}

func (fs FileStorage) Save(data []byte) error {
	return os.WriteFile(fs.path, data, 0644)
}

func (fs FileStorage) Load() ([]byte, error) {
	return os.ReadFile(fs.path)
}

// MemoryStorage メモリへの実装
type MemoryStorage struct {
	data []byte
}

func (ms *MemoryStorage) Save(data []byte) error {
	ms.data = make([]byte, len(data))
	copy(ms.data, data)
	return nil
}

func (ms *MemoryStorage) Load() ([]byte, error) {
	return ms.data, nil
}

func main() {

}

type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

func Greet(s Speaker) {
	fmt.Println(s.Speak())
}
func sample01() {
	person := Person{Name: "Alice"}
	Greet(person)
	// Hello, my name is Alice
}

type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return a.Name + " says woof!"
}
func sample02() {
	var s Speaker

	s = Person{Name: "Bob"}
	Greet(s)
	// Hello, my name is Bob

	s = Animal{Name: "Dog"}
	Greet(s)
	// Dog says woof!
}

func PrintAnything(a interface{}) {
	fmt.Println(a)
}
