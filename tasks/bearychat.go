package tasks

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/alastairruhm/notification-gateway/server/model"
	"github.com/bearyinnovative/bearychat-go"
	"github.com/mitchellh/mapstructure"
	"log"
)

type BearychatParam struct {
	Text         string
	Notification string
	Markdown     bool
	Channel      string
	User         string
	// 暂不支持 attachments
	// Attachments  []IncomingAttachment `json:"attachments,omitempty"`
	URL string
}

func BearychatNotify(queue string, args ...interface{}) error {
	fmt.Printf("queue: %s args: %v\n", queue, args)

	m := &model.Notification{}
	fmt.Println(m)
	// must be nofitication id
	notificationID, ok := args[0].(string)

	if !ok {
		fmt.Println("type assertion is error")
		return errors.New("args is error")
	}
	fmt.Println("notificationID: " + notificationID)

	n, err := m.FindByID(notificationID)
	if err != nil {
		return err
	}

	fmt.Printf("notification: %v\n", n)
	// here why not print out error ?
	var bm BearychatParam
	err = mapstructure.Decode(n.Param, &bm)
	if err != nil {
		return err
	}
	incoming := bearychat.Incoming{
		Text:         bm.Text,
		Notification: bm.Notification,
		Markdown:     bm.Markdown,
		Channel:      bm.Channel,
		User:         bm.Channel,
	}
	output, err := incoming.Build()
	if err != nil {
		return err
	}
	_, err = http.Post(bm.URL, "application/json", output)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("done")
	return nil
}
