package bll

import "github.com/alastairruhm/notification-gateway/server/model"

// Bll is Business Logic Layer with all models
type Bll struct {
	Models *model.Models
}

// Blls ...
type Blls struct {
	Models       *model.Models
	Notification *NotificationBll
}

// Init ...
func (bs *Blls) Init() *Blls {
	bs.Models = new(model.Models).Init()
	b := &Bll{bs.Models}
	bs.Notification = &NotificationBll{b}
	return bs
}
