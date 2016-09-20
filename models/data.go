package models

import (
	"database/sql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func Open(prod bool) {
	var err error
	if prod {
		db, err = sql.Open("sqlite3", "prod.db")
	} else {
		db, err = sql.Open("sqlite3", "test.db")
	}
	if err != nil {
		log.Fatal(err)
	}

	mkChTab()
	mkLsnTab()
	mkGramTab()

	// TODO: make testable JSON vs prod JSON
	if prod {
		parseJSON("json/babayaga.json")
	} else {
		parseJSON("../json/test.json")
	}
}

func Close() {
	defer db.Close()
}
