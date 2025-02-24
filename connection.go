package connections

import (
	"bytes"
	"database/sql"
	"errors"
	"net/http"
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

func NewQuery(query string) *Query {
	return &Query{query}
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

type Api struct {
	Host   string
	ApiKey string
}

func NewApi(host, apiKey string) (*Api, error) {
	api := &Api{
		Host:   host,
		ApiKey: apiKey,
	}

	if host == "" || apiKey == "" {
		return nil, errors.New("todos os campos são obrigatorios")
	}
	return api, nil
}

func (api *Api) UrlForEndpoint(endpoint string, params map[string]string) string {
	url := api.Host + "/" + endpoint
	if len(params) > 0 {
		url += "?"
		for k, v := range params {
			url += k + "=" + v + "&"
		}
	}

	return url
}

func (api *Api) NewRequest(method, url string, data []byte, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return req, nil
}

func NewClient() *http.Client {
	return &http.Client{}
}
