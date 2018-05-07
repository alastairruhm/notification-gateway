package worker

import (
	"github.com/alastairruhm/notification-gateway/tasks"

	"github.com/benmanns/goworker"
)

func init() {
	settings := goworker.WorkerSettings{
		URI:            "redis://localhost:6379/",
		Connections:    100,
		Queues:         []string{"bearychat", "delimited", "queues"},
		UseNumber:      true,
		ExitOnComplete: false,
		Concurrency:    2,
		Namespace:      "resque:",
		Interval:       5.0,
	}
	goworker.SetSettings(settings)
	goworker.Register("Bearychat", tasks.BearychatNotify)
}
