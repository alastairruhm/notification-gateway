package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alastairruhm/notification-gateway/server/bll"
	"github.com/alastairruhm/notification-gateway/server/schema"
	"github.com/bearyinnovative/bearychat-go"
	"github.com/mitchellh/mapstructure"
	"github.com/teambition/gear"
	"github.com/teambition/gear/logging"
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
	Notification string `json:"notification,omitempty"`
	Markdown     bool   `json:"markdown,omitempty"`
	Channel      string `json:"channel,omitempty"`
	User         string `json:"user,omitempty"`
	// 暂不支持 attachments
	// Attachments  []IncomingAttachment `json:"attachments,omitempty"`
}

func (i *NotificationAPI) Notify(ctx *gear.Context) error {

	message := Message{}
	if err := ctx.ParseBody(&message); err != nil {
		fmt.Errorf("%v", err)
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
		m := bearychat.Incoming{
			Text:         param.Text,
			Notification: param.Notification,
			Markdown:     param.Markdown,
			Channel:      param.Channel,
			User:         param.User,
		}
		output, _ := m.Build()
		nRecord := schema.Notification{
			Channel: message.Channel,
			// Param: param,
		}
		_, err = i.notificationBll.Create(&nRecord)

		if err != nil {
			logging.Err(err)
			return gear.ErrBadRequest.From(err)
		}

		http.Post("https://hook.bearychat.com/=bw74N/incoming/", "application/json", output)
	}

	return ctx.JSON(200, map[string]string{"status": "ok"})
}
