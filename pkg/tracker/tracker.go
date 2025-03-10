package tracker

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"slackbot/internal/config"
	"slackbot/internal/httpclient"
)

type Headers struct {
	Authorization string
	ContentType   string
}

type JiraTracker struct {
	URL          string
	Token        string
	TicketConfig config.TicketConfig
	Client       httpclient.HTTPClient
	Headers      Headers
}

// NewJiraTracker создает новый экземпляр JiraTracker.
func NewJiraTracker(cfg *config.Config, client httpclient.HTTPClient) (*JiraTracker, error) {
	if cfg.Tracker.URL == "" {
		return nil, errors.New("Jira URL is required") //nolint:stylecheck
	}
	if cfg.Tracker.Token == "" {
		return nil, errors.New("Jira token is required") //nolint:stylecheck
	}
	if client == nil {
		return nil, errors.New("HTTP client is required") //nolint:stylecheck
	}

	return &JiraTracker{
		URL:          cfg.Tracker.URL,
		Token:        cfg.Tracker.Token,
		TicketConfig: cfg.Ticket,
		Client:       client,
		Headers: Headers{
			Authorization: "Basic " + cfg.Tracker.Token,
			ContentType:   "application/json",
		},
	}, nil
}

// CreateTicket создает тикет в Jira.
func (t *JiraTracker) CreateTicket(ctx context.Context, title, description string) (string, error) {
	payload := map[string]interface{}{
		"fields": map[string]interface{}{
			"project":     t.TicketConfig.Project,
			"summary":     title,
			"description": description,
			"issuetype":   t.TicketConfig.Issuetype,
			"labels":      t.TicketConfig.Labels,
			"priority":    t.TicketConfig.Priority,
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := t.Client.Do(ctx, "POST", t.URL, bytes.NewBuffer(jsonPayload), map[string]string{
		"Authorization": t.Headers.Authorization,
		"Content-Type":  t.Headers.ContentType,
	})
	if err != nil {
		return "", fmt.Errorf("failed to create ticket: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("Jira API returned status: %d", resp.StatusCode) //nolint:stylecheck
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return result["key"].(string), nil
}

// GetTicketStatus возвращает статус тикета в Jira.
func (t *JiraTracker) GetTicketStatus(ctx context.Context, key string) (string, error) {
	resp, err := t.Client.Do(ctx, "GET", fmt.Sprintf("%s/%s", t.URL, key), nil, map[string]string{
		"Authorization": t.Headers.Authorization,
		"Content-Type":  t.Headers.ContentType,
	})
	if err != nil {
		return "", fmt.Errorf("failed to get ticket status: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Jira API returned status: %d", resp.StatusCode) //nolint:stylecheck
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return result["fields"].(map[string]interface{})["status"].(map[string]interface{})["name"].(string), nil
}
