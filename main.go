package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/alastairruhm/notification-gateway/server"
	"github.com/alastairruhm/notification-gateway/server/service/mongodb"
	"github.com/benmanns/goworker"
	"github.com/teambition/gear/logging"

	_ "github.com/alastairruhm/notification-gateway/worker"
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
		logging.Info("worker starts.")
		mongodb.CheckAndInitServiceConnection()
		if err := goworker.Work(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		return
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
