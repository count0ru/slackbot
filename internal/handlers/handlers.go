package handlers

import (
	"net/http"

	"slackbot/internal/logger"

	"go.uber.org/zap"
)

// HandleSlackEvent обработчик HTTP-запросов от Slack.
func HandleSlackEvent(w http.ResponseWriter, r *http.Request) {
	// Логирование входящего запроса
	logger.Log.Info("Received Slack event",
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
	)

	// Пример обработки запроса
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Event received"))
	if err != nil {
		logger.Log.Error("Failed to write response", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
