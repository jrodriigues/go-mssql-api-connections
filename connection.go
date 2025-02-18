package connections

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/microsoft/go-mssqldb"
)

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Pool     *sql.DB
}

func NewDatabase(host, port, user, password, database string) (*Database, error) {
	db := &Database{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Database: database,
	}

	if host == "" || port == "" || user == "" || password == "" || database == "" {
		return nil, errors.New("todos os campos são obrigatorios")
	}
	return db, nil
}

func (db *Database) ConnString() string {
	return "sqlserver://" + db.User + ":" + db.Password + "@" + db.Host + ":" + db.Port + "?database=" + db.Database
}

func (db *Database) Connect() error {
	conn, err := sql.Open("sqlserver", db.ConnString())

	if err != nil {
		return err
	} else {
		db.Pool = conn
		return nil
	}
}

func (db *Database) ExecQuery(q *Query) (*QueryResult, error) {
	var result *sql.Rows

	result, err := db.Pool.Query(q.String)

	return &QueryResult{result}, err
}

type Query struct {
	String string
}

func NewQueryFromFile(filename string) (*Query, error) {
	query, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Query{string(query)}, nil
}

type QueryResult struct {
	Rows *sql.Rows
}

type Destination struct {
	Host     string
	Endpoint string
	ApiKey   string
}
