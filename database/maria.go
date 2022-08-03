package database

import (
	"database/sql"
	"director/configs"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func MariaDB() (db *sql.DB) {
	dbDriver := "mysql"                       // Database driver
	dbUser := configs.Config("DB_USER_MARIA") // Mysql username
	dbPass := configs.Config("DB_PASS_MARIA") // Mysql password
	dbName := configs.Config("DB_DATABASE")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+"/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(50)
	return db
}
