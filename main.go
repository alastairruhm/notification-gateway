package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/alastairruhm/notification-gateway/server"
	"github.com/alastairruhm/notification-gateway/worker"
	"github.com/teambition/gear/logging"
)

var (
	portReg    = regexp.MustCompile(`^\d+$`)
	flagWorker = flag.Bool("worker", false, "launch worker")
	port       = flag.String("port", "8080", `Server port.`)
	version    = flag.Bool("version", false, "show app version")
)

func main() {
	flag.Parse()
	if *version {
		fmt.Println(server.VERSION)
		os.Exit(0)
	}

	if *flagWorker {
		worker := worker.Server.NewWorker("bearychat", 2)
		err := worker.Launch()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if portReg.MatchString(*port) {
		*port = ":" + *port
	}
	if *port == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	app := server.New()

	// start app
	logging.Info("Service start " + *port)

	err := app.Listen(*port)
	if err != nil {
		logging.Err(err)
		os.Exit(1)
	}
}
