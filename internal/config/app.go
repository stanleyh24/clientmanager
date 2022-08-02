package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DSN = "root:admin123@tcp(127.0.0.1:3306)/networkmanager?charset=utf8mb4&parseTime=True&loc=Local"

func CreateConnection() (*sql.DB, error) {

	db, err := sql.Open("mysql", DSN)
	if err != nil {
		return nil, err
	} else {
		log.Println("DB connection established successfully")
	}
	db.SetMaxIdleConns(5)
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
