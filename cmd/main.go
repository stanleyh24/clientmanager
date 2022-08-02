package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	db, err := createConnection()
	if err != nil {
		panic(err)
	}

	err = GetAll(ctx, db)
	if err != nil {
		panic(err)
	}
}

func createConnection() (*sql.DB, error) {
	DSN := "root:admin123@tcp(127.0.0.1:3306)/networkmanager?charset=utf8mb4&parseTime=True&loc=Local"
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

func GetAll(ctx context.Context, db *sql.DB) error {

	qry := `select id,ip,name,username,password from router`

	rows, err := db.QueryContext(ctx, qry)

	if err != nil {
		return err
	}
	fmt.Println(rows)
	return nil

}
