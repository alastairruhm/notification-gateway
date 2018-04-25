package server

import (
	"github.com/alastairruhm/notification-gateway/server/api"
	"github.com/alastairruhm/notification-gateway/server/bll"
	"github.com/teambition/gear"
)

func newRouter() (Router *gear.Router) {
	Router = gear.NewRouter()
	blls := new(bll.Blls).Init()
	apis := new(api.APIs).Init(blls)

	// Register a database instance
	Router.Post("/notification/", apis.Notification.Notify)
	return
}
