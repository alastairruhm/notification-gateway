package server

import (
	"github.com/alastairruhm/notification-gateway/server/service/mongodb"
	"github.com/teambition/gear"
	"github.com/teambition/gear/middleware/cors"
)

// VERSION is app version
const VERSION = "v0.1.0"

// New returns a app instance
func New() *gear.App {
	mongodb.CheckAndInitServiceConnection()
	app := gear.New()
	app.Use(cors.New())
	app.UseHandler(newRouter())
	return app
}
