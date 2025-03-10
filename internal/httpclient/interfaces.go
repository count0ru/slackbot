package httpclient

import (
	"context"
	"io"
	"net/http"
)

// HTTPClient интерфейс для HTTP-клиента.
type HTTPClient interface {
	Do(ctx context.Context, method string, url string, body io.Reader, headers map[string]string) (*http.Response, error)
}
