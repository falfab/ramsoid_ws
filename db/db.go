package db

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"ramsoid_ws/models"
	"log"
)

var db *sql.DB

func Init() error {
	var err error
	if db == nil {
		db, err = sql.Open("sqlite3", "./db/keys.sqlite")
		if err != nil {
			return err
		}
		_, err := db.Exec(CreateStatement)
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertRecord(record *models.Record) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(Insert)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(record.ID, record.Key)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func CheckKey(record *models.Record) bool {
	rows, err := db.Query(SelectKey, record.ID, record.Key)
	if err != nil {
		log.Fatal(err)
	}
	return rows.Next()

	// true if exists at least one record
	//return rows.Next()
}

func Close() {
	db.Close()
}

