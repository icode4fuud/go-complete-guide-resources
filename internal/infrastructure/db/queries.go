package db

//Adding Query Helpers keeps the models clean

import "database/sql"

func QueryRow(query string, args ...any) *sql.Row {
	return DB.QueryRow(query, args...)
}

func Exec(query string, args ...any) (sql.Result, error) {
	return DB.Exec(query, args...)
}

func Query(query string, args ...any) (*sql.Rows, error) {
	return DB.Query(query, args...)
}

func Ping() error {
	return DB.Ping()
}

func Close() error {
	return DB.Close()
}

func Begin() (*sql.Tx, error) {
	return DB.Begin()
}

func Prepare(query string) (*sql.Stmt, error) {
	return DB.Prepare(query)
}

func LastInsertId(result sql.Result) (int64, error) {
	return result.LastInsertId()
}

func RowsAffected(result sql.Result) (int64, error) {
	return result.RowsAffected()
}
