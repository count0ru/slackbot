package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"slackbot/internal/config"
	"slackbot/internal/handlers"
	"slackbot/internal/httpclient"
	"slackbot/internal/logger"
	"slackbot/internal/metrics" // Импортируем пакет metrics
	"slackbot/pkg/tracker"

	"go.uber.org/zap"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Panicf("Failed to load config: %v", err)
	}

	// Инициализация логгера
	logger.Init()
	defer logger.Sync()

	// Инициализация HTTP-клиента
	client, err := httpclient.NewClient()
	if err != nil {
		logger.Log.Panic("Failed to create HTTP client", zap.Error(err))
	}

	// Инициализация Tracker
	tracker, err := tracker.NewJiraTracker(cfg, client)
	if err != nil {
		logger.Log.Panic("Failed to create Jira tracker", zap.Error(err))
	}

	// Пример использования tracker
	key, err := tracker.CreateTicket(context.Background(), "Test Ticket", "Test Description")
	if err != nil {
		logger.Log.Error("Failed to create ticket", zap.Error(err))
	} else {
		logger.Log.Info("Ticket created", zap.String("key", key))
		metrics.TicketsCreated.Inc() // Увеличиваем счетчик созданных тикетов
	}

	// Инициализация метрик
	metrics.Init()
	metrics.StartServer(fmt.Sprintf("%d", cfg.Metrics.Port))

	// Запуск HTTP сервера
	http.HandleFunc("/slack/events", handlers.HandleSlackEvent)
	logger.Log.Info("Server started", zap.Int("port", cfg.General.Port))
	log.Panic(http.ListenAndServe(fmt.Sprintf(":%d", cfg.General.Port), nil))
}
