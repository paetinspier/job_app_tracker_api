package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/fatih/color"
	_ "github.com/lib/pq"
)

type DatabaseService struct {
	DB *sql.DB
}

var connectionInstance *sql.Conn

func Init() (*DatabaseService, error) {
	dsn := os.Getenv("dsn")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	s := &DatabaseService{DB: db}
	
	color.Green("Connected to database")
	return s, nil
}

func (ds *DatabaseService) PingDatabase() error {
	if err := ds.DB.Ping(); err != nil {
		return err
	}

	return nil
}

func (ds *DatabaseService) DeferClose() error {
	defer ds.DB.Close()

	return nil
}

// Query is an example method to interact with the database.
func (ds *DatabaseService) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return ds.DB.Query(query, args...)
}

// Exec is an example method to execute a non-query SQL statement.
func (ds *DatabaseService) Exec(query string, args ...interface{}) (sql.Result, error) {
	return ds.DB.Exec(query, args...)
}
