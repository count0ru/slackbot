package metrics

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// TicketsCreated счетчик созданных тикетов.
	TicketsCreated = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "tickets_created_total",
		Help: "Total number of tickets created.",
	})

	// TicketsClosed счетчик закрытых тикетов.
	TicketsClosed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "tickets_closed_total",
		Help: "Total number of tickets closed.",
	})
)

// Init инициализирует метрики Prometheus.
func Init() {
	prometheus.MustRegister(TicketsCreated)
	prometheus.MustRegister(TicketsClosed)
}

// StartServer запускает HTTP-сервер для экспорта метрик.
func StartServer(port string) {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
			panic(err)
		}
	}()
}
