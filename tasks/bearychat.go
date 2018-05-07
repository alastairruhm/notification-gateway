package tasks

import (
	"net/http"

	"github.com/alastairruhm/notification-gateway/server/service/mongodb"

	"github.com/alastairruhm/notification-gateway/server/model"

	"github.com/bearyinnovative/bearychat-go"
)

type BearychatIncomingMessage struct {
	bearychat.Incoming
	URL string
}

func BearychatNotify(queue string, args ...interface{}) error {
	m := &model.Notification{
		Conn: mongodb.NewCollectionSession("notification"),
	}
	// must be nofitication id
	notificationID := args[0].(string)
	n, err := m.FindByID(notificationID)
	if err != nil {
		return err
	}
	bm := &BearychatIncomingMessage{
		bearychat.Incoming{
			Text:         n.Param["text"].(string),
			Notification: n.Param["notification"].(string),
			Markdown:     n.Param["markdown"].(bool),
			Channel:      n.Param["channel"].(string),
			User:         n.Param["user"].(string),
			Attachments:  nil,
		},
		n.Param["url"].(string),
	}
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
