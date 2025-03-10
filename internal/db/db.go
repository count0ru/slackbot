package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

// SQLDB реализация интерфейса DB.
type SQLDB struct {
	conn *sql.DB
}

// NewDB создает новое соединение с базой данных.
func NewDB(path string, log *zap.Logger) (DB, error) {
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Info("Database connection established", zap.String("path", path))
	return &SQLDB{conn: conn}, nil
}

// Conn возвращает соединение с базой данных.
func (db *SQLDB) Conn() *sql.DB {
	return db.conn
}

// Close закрывает соединение с базой данных.
func (db *SQLDB) Close() error {
	return db.conn.Close()
}

// Exec выполняет SQL-запрос без возвращения строк.
func (db *SQLDB) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.conn.ExecContext(ctx, query, args...)
}

// Query выполняет SQL-запрос и возвращает строки.
func (db *SQLDB) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return db.conn.QueryContext(ctx, query, args...)
}

// QueryRow выполняет SQL-запрос и возвращает одну строку.
func (db *SQLDB) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return db.conn.QueryRowContext(ctx, query, args...)
}
