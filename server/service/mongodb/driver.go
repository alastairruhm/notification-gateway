package mongodb

import (
	"github.com/teambition/gear/logging"
)

// MaxPool max pool size
var MaxPool = 300

func CheckAndInitServiceConnection() {
	logging.Info("check mongodb connection")
	if service.baseSession == nil {
		service.URL = "127.0.0.1:27017"
		err := service.New()
		if err != nil {
			// logger.Err(err)
			logging.Fatal(err)
		}
	}
}
