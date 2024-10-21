package gddb

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	Id          int
	Title       string
	Description string
	IsComplete  bool
}

// Custom Psuedo Enum
type DeleteType int

const (
	title = iota + 1
	id
)

func InitDB(defaultTable string, dbDir *string) (err error) {
	// Get DB
	database, err := GetDB(dbDir)
	if err != nil {
		return err
	}

	// Close connection when this function is done
	defer database.Close()

	err = CreateTableInDB(defaultTable, dbDir)
	if err != nil {
		return err
	}

	return nil
}

func InsertTaskInDB(title *string, description *string, table *string, dbDir *string) error {

	if title == nil || description == nil || table == nil {
		err := fmt.Errorf("in InsertTaskInDB, a supplied argument is a nil string pointer")
		return err
	}

	db, err := GetDB(dbDir)
	if err != nil {
		return err
	}

	defer db.Close()

	query := fmt.Sprintf("INSERT INTO %s (title, description) VALUES (?, ?)", *table)
	statement, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("in InsertTaskInDB, something went wrong while preparing the query")
	}

	_, err = statement.Exec(*title, *description)
	if err != nil {
		return fmt.Errorf("in InsertTaskInDB, something went wrong while executing the query")
	}

	// Successfully inserted task into DB
	return nil
}

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
func ListTasksInTable(tableName *string, dbDir *string) (tasks []Task, err error) {

	if tableName == nil {
		err := fmt.Errorf("in ListTasksInTable, tableName was supplied with a nil string pointer")
		return nil, err
	}

	db, err := GetDB(dbDir)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM %s", *tableName)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	// Add all tasks from the query
	for rows.Next() {
		task := Task{}
		rows.Scan(&task.Id, &task.Title, &task.Description, &task.IsComplete)
		tasks = append(tasks, task)
	}

	return tasks, nil // Successfully printed table
}

func CreateTableInDB(newTableName string, dbDir *string) error {

	db, err := GetDB(dbDir)
	if err != nil {
		return fmt.Errorf("in CreateTableInDB, something went wrong with getting the db")
	}

	defer db.Close()

	table := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id INTEGER NOT NULL PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		isComplete BOOLEAN NOT NULL DEFAULT FALSE);`, newTableName)

	_, err = db.Exec(table)
	if err != nil {
		return fmt.Errorf("in CreateTableInDB, something went wrong with executing the query")
	}

	return nil // Table was created
}

func getQueryString(queryType DeleteType, dbTable string, thingToDelete *interface{}) string {

	query := fmt.Sprintf("DELETE FROM %s", dbTable)

	switch queryType {
	case title:
		query = fmt.Sprintf("%s WHERE title='%s'", query, *thingToDelete)
	case id:
		query = fmt.Sprintf("%s WHERE id=%d", query, *thingToDelete)
	}

	return query
}

func DeleteTaskInDB(queryType DeleteType, dbTable string, thingToDelete *interface{}, dbDir *string) error {

	if thingToDelete == nil {
		err := fmt.Errorf("in DeleteTaskInDB, thingToDelete was supplied a nil interface pointer")
		return err
	}

	db, err := GetDB(dbDir)
	if err != nil {
		return err
	}

	defer db.Close()

	query := getQueryString(queryType, dbTable, thingToDelete)
	fmt.Printf("query: %v\n", query)

	statement, err := db.Prepare(query)
	if err != nil {
		return err
	}

	statement.Exec() // Exec query

	return nil // Task successfully deleted
}

func GetDB(dbDir *string) (database *sql.DB, err error) {

	if dbDir == nil {
		err := fmt.Errorf("in GetDB, dbDir was supplied a nil string pointer")
		return nil, err
	}

	dbPath := filepath.Join(*dbDir, "godo.db")

	database, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return database, nil
}
