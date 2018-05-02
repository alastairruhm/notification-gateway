package worker

import (
	"fmt"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/alastairruhm/notification-gateway/tasks"
	// "github.com/RichardKnop/machinery/v1/"
)

var Server *machinery.Server

func init() {
	var err error
	Server, err = NewServer()
	if err != nil {
		panic(err)
	}
}

// NewServer creates new server
func NewServer() (*machinery.Server, error) {
	cnf := &config.Config{
		Broker:        "redis://127.0.0.1:6379",
		DefaultQueue:  "machinery_tasks",
		ResultBackend: "mongodb://127.0.0.1:27017/taskresults",
	}

	server, err := machinery.NewServer(cnf)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// register task
	server.RegisterTask("bearychat", tasks.BearychatNotify)

	return server, nil
}
