package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID      uint         `gorm:"primaryKey"`
	Name    string       `gorm:"size:100;not null"`
	Email   string       `gorm:"unique;not null"`
	Profile *UserProfile `gorm:"foreignKey:UserID"`
}

type UserProfile struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"unique;not null"`
	Age         int    `gorm:"not null"`
	PhoneNumber string `gorm:"size:20"`
	User        *User  `gorm:"foreignKey:UserID"`
}

type Book struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:100;not null"`
	PublishDate time.Time `gorm:"not null"`
	Author      Author    `gorm:"foreignKey:BookID"` // 1:1 の関連を追加
}

type Author struct {
	ID     uint   `gorm:"primaryKey"`
	BookID uint   `gorm:"unique;not null"`
	Name   string `gorm:"size:100;not null"`
	Book   *Book  `gorm:"foreignKey:BookID"` // 逆方向の関連
}

func main() {
	// section1()
	// section2()
	// section3()
}

func section3() {
	migrateUserProfile()
	createUserProfiles()
	sampleInnerJoin()
	sampleLeftJoin()
}

func sampleLeftJoin() {
	db := dbConnection()

	var users []User
	db.Joins("Profile").Find(&users)

	for _, user := range users {
		var age string
		if user.Profile == nil {
			age = "記載なし"
		} else {
			age = strconv.Itoa(user.Profile.Age)
		}
		fmt.Printf("name: %s (age: %s)\n", user.Name, age)
	}
}

func sampleInnerJoin() {
	db := dbConnection()

	var users []User
	db.InnerJoins("Profile").Find(&users)

	for _, user := range users {
		fmt.Printf("name: %s (age: %d)\n", user.Name, user.Profile.Age)
	}
}

func createUserProfiles() {
	db := dbConnection()

	// 作成する UserProfile のデータ
	profiles := []UserProfile{
		{
			UserID:      2,
			Age:         25,
			PhoneNumber: "090-1234-5678",
		},
		{
			UserID:      4,
			Age:         30,
			PhoneNumber: "090-8765-4321",
		},
	}

	// バルクインサート
	result := db.Create(&profiles)
	if result.Error != nil {
		log.Fatal("failed to create profiles: ", result.Error)
	}
}

func migrateUserProfile() {
	db := dbConnection()

	// テーブルのマイグレーション
	err := db.AutoMigrate(&UserProfile{})
	if err != nil {
		log.Fatal("マイグレーションエラー: ", err)
	}

	log.Println("マイグレーションが成功しました")
}

func section2() {
	getBookWithoutPreload()
	getBookWithPreload()
	getBookWithAuthor()
	getBooksAfterDate()
}

// 特定の出版日以降の本とその著者を取得
func getBooksAfterDate() {
	db := dbConnection()
	var books []Book
	date := time.Date(2023, 4, 7, 0, 0, 0, 0, time.UTC)
	result := db.Preload("Author").
		Where("publish_date >= ?", date).
		Find(&books)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// データの表示
	for _, book := range books {
		fmt.Printf("書籍: %s (出版日: %s)\n",
			book.Name,
			book.PublishDate.Format("2006-01-02"))
		fmt.Printf("著者: %s\n\n", book.Author.Name)
	}
}

func getBookWithAuthor() {
	db := dbConnection()
	var book Book
	bookID := 1
	result := db.Preload("Author").First(&book, bookID)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Printf("book %+v\n", book)
}

func getBookWithPreload() {
	db := dbConnection()
	var books []Book
	db.Preload("Author").Find(&books)

	fmt.Printf("%#v\n", books[0])
}

func getBookWithoutPreload() {
	db := dbConnection()
	var books []Book
	db.Find(&books) // 1回目のクエリ

	// 各本に対して著者を個別に取得（N回のクエリ）
	for _, book := range books {
		var author Author
		db.Where("book_id = ?", book.ID).First(&author)
	}
}

func section1() {
	dropTable()
	migrate()
	createUser()
	createUsers()
	getUser()
	getUserWhere()
	updateUser()
	updateUserMultiColumns()
	deleteUser()
	withTransaction()
}

func withTransaction() {
	db := dbConnection()

	// トランザクションの開始
	err := db.Transaction(func(tx *gorm.DB) error {
		user := User{Name: "斎藤花子", Email: "hanako@example.com"}
		if err := tx.Create(&user).Error; err != nil {
			return err // エラーが発生した場合ロールバック
		}

		user.Name = "斎藤 花子"
		if err := tx.Model(&user).Update("Name", "斎藤 花子").Error; err != nil {
			return err // エラーが発生した場合ロールバック
		}

		return nil // エラーが無ければコミット
	})

	if err != nil {
		log.Println("トランザクションエラー:", err)
	} else {
		log.Println("トランザクション成功")
	}
}

func deleteUser() {
	db := dbConnection()
	var user User
	db.Where("id = ?", 1).Delete(&user)
}

func updateUserMultiColumns() {
	db := dbConnection()
	user := User{Name: "山田 太郎", Email: "taro@example.com"}
	db.Model(&user).Where("id = ?", 1).Updates(user)
}

func updateUser() {
	db := dbConnection()
	var user User
	db.Model(&user).Where("email = ?", "suzuki@example.com").Update("Name", "鈴木二郎")
}

func getUserWhere() {
	db := dbConnection()
	var user User
	db.Where("email = ?", "sato@example.com").First(&user)
	log.Println(user)
}

func getUser() {
	db := dbConnection()
	var user User
	db.First(&user, 1) // ID が 1 のユーザーを取得
	log.Println(user)
}

func createUsers() {
	db := dbConnection()

	users := []User{
		{Name: "佐藤一郎", Email: "sato@example.com"},
		{Name: "鈴木次郎", Email: "suzuki@example.com"},
	}
	result := db.Create(&users)
	if result.Error != nil {
		log.Println("登録に失敗しました: ", result.Error)
	}

	var ids []uint
	for _, user := range users {
		ids = append(ids, user.ID)
	}
	log.Printf("登録成功: IDs=%v\n", ids)
}

func createUser() {
	db := dbConnection()

	user := User{Name: "山田太郎", Email: "yamada@example.com"}
	result := db.Create(&user)
	if result.Error != nil {
		log.Println("登録に失敗しました: ", result.Error)
	}
	log.Printf("登録成功: ID=%d\n", user.ID)
}

func dropTable() {
	db := dbConnection()
	err := db.Migrator().DropTable(&User{})
	if err != nil {
		return
	}
}

func migrate() {
	db := dbConnection()

	// テーブルのマイグレーション
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("マイグレーションエラー: ", err)
	}

	log.Println("マイグレーションが成功しました")
}

func dbConnection() *gorm.DB {
	// PostgresSQL接続用のDSN（データソース名）
	dsn := "host=localhost user=postgres password=postgres dbname=golang_sample_db port=5432 sslmode=disable TimeZone=Asia/Tokyo"

	// データベース接続
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("データベース接続エラー: ", err)
	}

	log.Println("PostgresSQLに接続しました")

	return db
}
