package api

import (
	"fmt"
	"time"

	"github.com/alastairruhm/notification-gateway/server/bll"
	"github.com/alastairruhm/notification-gateway/server/schema"
	"github.com/benmanns/goworker"
	"github.com/mitchellh/mapstructure"
	"github.com/teambition/gear"
	"github.com/teambition/gear/logging"

	_ "github.com/alastairruhm/notification-gateway/worker"
)

type NotificationAPI struct {
	CommonAPI
	notificationBll *bll.NotificationBll
}

// Init ...
func (i *NotificationAPI) Init(api CommonAPI) *NotificationAPI {
	i.CommonAPI = api
	i.notificationBll = api.blls.Notification
	return i
}

type Message struct {
	Channel string                 `json:"channel"`
	Param   map[string]interface{} `json:"param"`
	ETA     time.Time              `json:"eta,omitempty"`
}

// Validate message template validate implementation
func (m *Message) Validate() error {
	if m.Channel == "" {
		return gear.ErrBadRequest.WithMsg("channel required")
	}
	return nil
}

type ParamBearychatIncoming struct {
	Text         string `json:"text"`
	Notification string `json:"notification"`
	Markdown     bool   `json:"markdown"`
	Channel      string `json:"channel"`
	User         string `json:"user"`
	// 暂不支持 attachments
	// Attachments  []IncomingAttachment `json:"attachments,omitempty"`
	URL string `json:"url"`
}

func (i *NotificationAPI) Notify(ctx *gear.Context) error {

	message := Message{}
	if err := ctx.ParseBody(&message); err != nil {
		fmt.Printf(err.Error())
		logging.Err(err)
		return gear.ErrBadRequest.From(err)
	}

	switch message.Channel {
	case "bearychat":
		var param ParamBearychatIncoming
		err := mapstructure.Decode(message.Param, &param)
		if err != nil {
			logging.Err(err)
			return gear.ErrBadRequest.From(err)
		}

		nRecord := schema.Notification{
			Channel: message.Channel,
			Param:   message.Param,
		}
		n, err := i.notificationBll.Create(&nRecord)

		if err != nil {
			logging.Err(err)
			return gear.ErrBadRequest.From(err)
		}
		notificationID := n.ID.Hex()

		err = goworker.Enqueue(&goworker.Job{
			Queue: "bearychat",
			Payload: goworker.Payload{
				Class: "Bearychat",
				Args:  []interface{}{notificationID},
			},
		})

		if err != nil {
			logging.Err(err)
			return gear.ErrBadRequest.From(err)
		}

		logging.Info("enqueue bearychat notification " + notificationID)
	}

	return ctx.JSON(200, map[string]string{"status": "ok"})
}
