package main

import (
	"net/http"

	"github.com/ayoul3/phishkiller/lib"
	log "github.com/sirupsen/logrus"
)

var config *lib.Configuration

func init() {
	config = lib.GetConfig()
	lib.Chan = make(chan []*http.Request, config.Workers)
	config.SetLogLevel()
}

func main() {
	client := lib.CreateNewClient(config)
	log.Infof("Starting %d workers", config.Workers)
	for i := 0; i < config.Workers; i++ {
		go lib.Perform(client)
	}

	log.Info("Initiating requests")
	lib.LoopRequests(client, config.Requests)
}
