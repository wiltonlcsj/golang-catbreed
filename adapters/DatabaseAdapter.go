package adapters

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type DatabaseAdapter struct{}

func NewDatabaseAdapter() *DatabaseAdapter {
	return &DatabaseAdapter{}
}

func (adapter *DatabaseAdapter) CreateConnection() (db *sql.DB, err error) {
	var (
		host     = os.Getenv("DATABASE_HOST")
		user     = os.Getenv("DATABASE_USER")
		password = os.Getenv("DATABASE_PASS")
		port     = os.Getenv("DATABASE_PORT")
		dbname   = os.Getenv("DATABASE_NAME")
	)

	db, err = sql.Open(
		"mysql",
		fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s`, user, password, host, port, dbname),
	)

	if err != nil {
		return nil, errors.New("cannot open sql connection")
	}

	return db, nil
}

func (adapter *DatabaseAdapter) CloseConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Println(err)
	}
}
