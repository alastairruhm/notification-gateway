package api

import (
	"github.com/alastairruhm/notification-gateway/server/bll"
	"github.com/alastairruhm/notification-gateway/server/model"
)

// APIs - all APIs
type APIs struct {
	Notification *NotificationAPI
}

// CommonAPI - for blls and models dependency-injection
type CommonAPI struct {
	blls   *bll.Blls
	models *model.Models
}

// Init - 初始化
func (a *APIs) Init(blls *bll.Blls) *APIs {
	api := CommonAPI{blls, blls.Models}
	*a = APIs{
		Notification: new(NotificationAPI).Init(api),
	}
	return a
}
