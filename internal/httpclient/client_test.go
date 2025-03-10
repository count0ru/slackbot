package httpclient

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	// Создаем тестовый сервер
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	// Создаем клиент
	client, err := NewClient()
	assert.NoError(t, err, "Failed to create HTTP client")

	// Выполняем GET-запрос
	resp, err := client.Do(context.Background(), "GET", ts.URL, nil, nil)
	assert.NoError(t, err, "Failed to perform GET request")
	defer resp.Body.Close() // Закрываем тело ответа
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Unexpected status code")
}
