package bll

import (
	"github.com/alastairruhm/notification-gateway/server/schema"
)

// NotificationBll ...
type NotificationBll struct {
	*Bll
}

// Create will generate a record of database
func (n *NotificationBll) Create(notification *schema.Notification) (*schema.Notification, error) {
	e, err := n.Models.Notification.Save(notification)

	if err != nil {
		return nil, err
	}
	return e, nil
}
