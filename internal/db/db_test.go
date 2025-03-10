package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNewDB(t *testing.T) {
	// Инициализация логгера
	logger, _ := zap.NewProduction()
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Error("Failed to sync logger", zap.Error(err))
		}
	}()

	// Тест: Создание новой базы данных
	dbPath := "test_new.db"
	defer os.Remove(dbPath) // Удаление файла после теста

	dbInstance, err := NewDB(dbPath, logger)
	assert.NoError(t, err)
	assert.NotNil(t, dbInstance.Conn)

	err = dbInstance.Close()
	assert.NoError(t, err)
}

func TestNewDB_ExistingFile(t *testing.T) {
	// Инициализация логгера
	logger, _ := zap.NewProduction()
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Error("Failed to sync logger", zap.Error(err))
		}
	}()

	// Тест: Использование существующего файла базы данных
	dbPath := "test_existing.db"
	_, err := os.Create(dbPath)
	assert.NoError(t, err)
	defer os.Remove(dbPath) // Удаление файла после теста

	dbInstance, err := NewDB(dbPath, logger)
	assert.NoError(t, err)
	assert.NotNil(t, dbInstance.Conn)

	err = dbInstance.Close()
	assert.NoError(t, err)
}

func TestNewDB_InvalidPath(t *testing.T) {
	// Инициализация логгера
	logger, _ := zap.NewProduction()
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Error("Failed to sync logger", zap.Error(err))
		}
	}()

	// Тест: Некорректный путь к базе данных
	dbPath := "/invalid/path/to/database.db"
	dbInstance, err := NewDB(dbPath, logger)
	assert.Error(t, err)
	assert.Nil(t, dbInstance)
}
