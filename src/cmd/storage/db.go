package storage

import (
	"github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


var db *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := mysql.Config{
        User: os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net: os.Getenv("DB_NET"),
        Addr: os.Getenv("DB_ADDRESS"),
        DBName: "fitness",
    }

    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }

	fmt.Println("Successfully connected to fitness database")
}