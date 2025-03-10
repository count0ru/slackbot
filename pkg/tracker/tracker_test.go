package tracker

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"slackbot/internal/config"
	"slackbot/internal/httpclient"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTicket(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := httpclient.NewMockHTTPClient(ctrl)

	// Настройка мока для HTTP-запроса
	mockClient.EXPECT().
		Do(gomock.Any(), "POST", gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&http.Response{
			StatusCode: http.StatusCreated,
			Body:       io.NopCloser(bytes.NewBufferString(`{"key": "TEST-123"}`)),
		}, nil)

	cfg := &config.Config{
		Tracker: config.TrackerConfig{
			URL:   "https://test-domain.atlassian.net/rest/api/3/issue",
			Token: "test-token",
		},
		Ticket: config.TicketConfig{
			Project:   config.ProjectConfig{Key: "TEST"},
			Issuetype: config.IssuetypeConfig{Name: "Task"},
			Labels:    []string{"Critical"},
			Priority:  config.PriorityConfig{Name: "High"},
		},
	}

	// Исправляем вызов NewJiraTracker
	tracker, err := NewJiraTracker(cfg, mockClient)
	assert.NoError(t, err, "Failed to create Jira tracker")

	// Пример использования tracker
	key, err := tracker.CreateTicket(context.Background(), "Test Ticket", "Test Description")
	assert.NoError(t, err, "Failed to create ticket")
	assert.Equal(t, "TEST-123", key, "Unexpected ticket key")
}
