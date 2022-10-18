package config

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "gorm.io/driver/mysql"
)

const (
	host   = "localhost"
	port   = "3306"
	user   = "root"
	pass   = "root"
	dbname = "assignment2"
)

var DB *sql.DB

func ConnectionMysql() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return err
	}

	DB = db

	return nil
}

func GetMysql() *sql.DB {
	return DB
}
