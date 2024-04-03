package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"unique"`
	Password string
}

type InvitationCode struct {
	ID   uint   `gorm:"primary_key"`
	Code string `gorm:"unique"`
	Used bool
}

const (
	Host     = "localhost"
	Port     = "5432"
	UserName = "root"
	DBName   = "User_Management"
	Password = "Hiren123"
)

func InitDB() {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", Host, Port, UserName, DBName, Password)
	var err error
	Db, err = gorm.Open("postgres", connectionString)
	if err != nil {
		panic("failed to connect database")
	}
	// AutoMigrate will create missing tables based on model structs
	Db.AutoMigrate(&User{}, &InvitationCode{})
}
