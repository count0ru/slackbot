package httpclient

import (
	"context"
	"errors"
	"io"
	"net/http"
)

// Client реализация интерфейса HTTPClient.
type Client struct {
	client *http.Client
}

// NewClient создает новый экземпляр Client.
func NewClient() (*Client, error) {
	if http.DefaultClient == nil {
		return nil, errors.New("failed to create HTTP client")
	}

	return &Client{
		client: http.DefaultClient,
	}, nil
}

// Do выполняет HTTP-запрос.
func (c *Client) Do(ctx context.Context, method string, url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return c.client.Do(req)
}
