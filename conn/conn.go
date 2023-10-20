package conn

import (
	"github.com/datrine/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := "datrine:TeMi4ToPe@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(&models.Book{}, &models.Account{}, &models.User{})
	if err != nil {
		println(err.Error())
		panic(err)
	}
	println("Connection established.")
}

func InitGetDB() *gorm.DB {
	var db *gorm.DB
	dsn := "datrine:TeMi4ToPe@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.Book{}, &models.Account{}, &models.User{})
	if err != nil {
		println(err.Error())
		panic(err)
	}
	println("Connection established....")
	DB = db
	return db
}

func GetDB() *gorm.DB {
	var db *gorm.DB
	if DB != nil {
		db = DB
		return db
	}
	db = InitGetDB()
	return db
}
