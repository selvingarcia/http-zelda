package main

import (
	"database/sql"
	"encoding/csv"
	"os"
	"strconv"
)

const createCodesTable = `CREATE TABLE IF NOT EXISTS codes(
   code int primary key,
   phrase varchar(50) not null,
   title varchar(60) not null,
   content text not null,
   meaning varchar(120) not null
)`

// initSchema creates the codes table if missing and seeds it from the CSV
// when empty. Safe to run on every startup.
func initSchema(db *sql.DB, csvPath string) error {
	if _, err := db.Exec(createCodesTable); err != nil {
		return err
	}
	var n int
	if err := db.QueryRow(`SELECT count(*) FROM codes`).Scan(&n); err != nil {
		return err
	}
	if n > 0 {
		return nil
	}
	f, err := os.Open(csvPath)
	if err != nil {
		return err
	}
	defer f.Close()
	rows, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	for _, r := range rows[1:] {
		code, err := strconv.Atoi(r[0])
		if err != nil {
			return err
		}
		if _, err := tx.Exec(`INSERT INTO codes(code, phrase, title, content, meaning) VALUES($1,$2,$3,$4,$5) ON CONFLICT DO NOTHING`,
			code, r[1], r[2], r[3], r[4]); err != nil {
			return err
		}
	}
	return tx.Commit()
}
