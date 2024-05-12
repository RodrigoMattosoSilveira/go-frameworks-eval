package storage

import (
	"github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
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
        ParseTime: true,
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

// Incliuded in the repository, but not in the blog;
// I was not able to send the author a note to inform him
// 
func GetDB() *sql.DB {
	return db
}