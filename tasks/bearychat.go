package tasks

import (
	"net/http"

	"github.com/alastairruhm/notification-gateway/server/model"
	"github.com/bearyinnovative/bearychat-go"
)

type BearychatIncomingMessage struct {
	bearychat.Incoming
	URL string
}

func BearychatNotify(notificationID string) error {
	models := model.Models{}
	m := models.Init()

	n, err := m.Notification.FindByID(notificationID)
	if err != nil {
		return err
	}
	bm := &BearychatIncomingMessage{
		bearychat.Incoming{
			n.Param["text"].(string),
			n.Param["notification"].(string),
			n.Param["markdown"].(bool),
			n.Param["channel"].(string),
			n.Param["user"].(string),
			nil,
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
