package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func main() {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", user, password, database)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Couldn't establish a connection with the database", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Couldn't ping database", err)
	}
}
