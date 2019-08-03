package db

import (
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialect
)

var (
	once     sync.Once
	instance *MySQLClient
)

const (
	username = "feedback_user"
	password = "password"
	dbname   = "feedback"
)

// MySQLClient ...
type MySQLClient struct {
	*gorm.DB
}

// User ...
type User struct {
	ID       int32  `gorm:"NOT NULL"`
	Username string `gorm:"NOT NULL"`
}

// GetInstance ...
func GetInstance() *MySQLClient {
	once.Do(func() {
		connStr := fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True&loc=Local", username, password, dbname)
		client, err := gorm.Open("mysql", connStr)
		log.Printf("connect to db: %v", connStr)
		if err != nil {
			log.Fatalf("cannot connect database: %v", err.Error())
		}
		client.LogMode(true)
		instance = &MySQLClient{client}
	})

	return instance
}

func init() {
	db := GetInstance()
	if !db.HasTable(&User{}) {
		db.AutoMigrate(&User{})
	}
}
