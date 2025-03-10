package metrics

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/assert"
)

func TestMetrics(t *testing.T) {
	// Инициализация метрик
	Init()

	// Увеличиваем счетчики
	TicketsCreated.Inc()
	TicketsClosed.Inc()

	// Создаем тестовый HTTP-запрос
	req := httptest.NewRequest("GET", "/metrics", nil)
	w := httptest.NewRecorder()

	// Обрабатываем запрос
	promhttp.Handler().ServeHTTP(w, req)

	// Проверяем статус ответа
	assert.Equal(t, http.StatusOK, w.Code, "Unexpected status code")

	// Проверяем, что метрики присутствуют в ответе
	assert.Contains(t, w.Body.String(), "tickets_created_total", "Metric 'tickets_created_total' not found")
	assert.Contains(t, w.Body.String(), "tickets_closed_total", "Metric 'tickets_closed_total' not found")
}
