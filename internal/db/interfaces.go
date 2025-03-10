package db

import (
	"context"
	"database/sql"
)

// DB интерфейс для работы с базой данных.
type DB interface {
	// Conn возвращает соединение с базой данных.
	Conn() *sql.DB

	// Close закрывает соединение с базой данных.
	Close() error

	// Exec выполняет SQL-запрос без возвращения строк.
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	// Query выполняет SQL-запрос и возвращает строки.
	Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)

	// QueryRow выполняет SQL-запрос и возвращает одну строку.
	QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row
}
