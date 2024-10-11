package dbdriver

import (
	"log"
)

func ListTablesInDB(dbDir *string) []string {
	db, err := GetDB(dbDir)
	if err != nil {
		log.Fatal(err)
	}

	query := `SELECT tbl_name FROM sqlite_master WHERE type="table"`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	var tableName string
	var tables []string
	for rows.Next() {
		rows.Scan(&tableName)
		tables = append(tables, tableName)
	}

	return tables
}
