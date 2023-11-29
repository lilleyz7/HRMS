package database

import (
	"database/sql"
	"log"
	"os"
	"time"
)

var DBconn *sql.DB

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		println(err)
	}

	return db, nil
}

func ConnectToDB() {
	dsn := os.Getenv("DB_URL")
	counts := 0

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not ready yet")
			counts++
		} else {
			log.Println("Connected successfully")

			DBconn = connection
			return
		}

		if counts > 10 {
			log.Println(err)
			DBconn = nil
			return
		}

		log.Println("Back off for 2 sec")
		time.Sleep(2 * time.Second)
		continue
	}
}
