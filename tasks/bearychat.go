package tasks

import (
	"net/http"

	"github.com/bearyinnovative/bearychat-go"
)

type BearychatIncomingMessage struct {
	bearychat.Incoming
	URL string
}

func BearychatNotify(bm BearychatIncomingMessage) error {
	output, err := bm.Build()
	if err != nil {
		return err
	}
	_, err = http.Post(bm.URL, "application/json", output)
	if err != nil {
		return err
	}
	return nil
}
