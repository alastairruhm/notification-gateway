package api

import "github.com/alastairruhm/notification-gateway/server/bll"

// APIs - all APIs
type APIs struct {
	Notification *NotificationAPI
}

// CommonAPI - for blls
type CommonAPI struct {
	blls *bll.Blls
}

// Init - 初始化
func (a *APIs) Init(blls *bll.Blls) *APIs {
	api := CommonAPI{blls}
	*a = APIs{
		Notification: new(NotificationAPI).Init(api),
	}
	return a
}
