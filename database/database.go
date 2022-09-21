package database

import (
	"gorm.io/gorm"
	"os"
	"sync"
)

var (
	dbConn *gorm.DB
	once   sync.Once
)

func CreateConnection() {
	conf := dbConfig{
		User: os.Getenv("DB_USERNAME"),
		Pass: os.Getenv("db_password"),
		Host: os.Getenv("db_host"),
		Port: os.Getenv("db_port"),
		Name: os.Getenv("db_name"),
	}

	mysql := mysqlConfig{dbConfig: conf}
	once.Do(func() {
		mysql.Connect()
	})
}

func GetConnection() *gorm.DB {
	if dbConn == nil {
		CreateConnection()
	}
	return dbConn
}
