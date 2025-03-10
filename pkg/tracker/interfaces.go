package tracker

import (
	"context"
)

// Tracker интерфейс для работы с тикетами.
type Tracker interface {
	// CreateTicket создает тикет.
	CreateTicket(ctx context.Context, title, description string) (string, error)

	// GetTicketStatus возвращает статус тикета.
	GetTicketStatus(ctx context.Context, key string) (string, error)
}
