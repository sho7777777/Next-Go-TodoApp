package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Connect() *sql.DB {
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbNet := os.Getenv("DB_NET")
	c := mysql.Config{
		DBName: dbName,
		User:   dbUser,
		Passwd: dbPasswd,
		Addr:   dbAddr,
		Net:    dbNet,
	}
	Db, err := sql.Open("mysql", c.FormatDSN())

	// Db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/todo")
	if err != nil {
		fmt.Printf("Error while connecting database: %v", err)
	}
	fmt.Println("connect.go", Db)
	err = Db.Ping()
	if err != nil {
		// fmt.Printf("Failed to connect to the database: %v", err)
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	return Db
}
