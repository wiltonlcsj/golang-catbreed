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

func (adapter *DatabaseAdapter) Ping() error {
	conn, err := adapter.CreateConnection()
	if err != nil {
		return err
	}

	err = conn.Ping()
	if err != nil {
		return err
	}

	_ = conn.Close()
	return nil
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

func (adapter *DatabaseAdapter) QueryMultiple(sqlcommand string, args ...interface{}) (*sql.Rows, error) {
	db, err := adapter.CreateConnection()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("cannot open connection")
	}

	results, err := db.Query(sqlcommand, args...)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("cannot execute query string")
	}

	adapter.CloseConnection(db)

	return results, nil
}

func (adapter *DatabaseAdapter) ExecuteCommand(sqlcommand string, args ...interface{}) (sql.Result, error) {
	db, err := adapter.CreateConnection()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("cannot open connection")
	}

	res, err := db.Exec(sqlcommand, args...)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("cannot execute query string")
	}

	adapter.CloseConnection(db)

	return res, nil
}

func (adapter *DatabaseAdapter) QueryRow(sqlcommand string, args ...interface{}) (*sql.Row, error) {
	db, err := adapter.CreateConnection()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("cannot open connection")
	}

	row := db.QueryRow(sqlcommand, args...)

	adapter.CloseConnection(db)

	return row, nil
}