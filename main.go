package main

import (
	"time"

	"github.com/ayoul3/phishkiller/lib"
	"github.com/prometheus/common/log"
)

var config *lib.Configuration

func init() {

	config = lib.GetConfig()
	lib.Chan = make(chan bool, config.Workers)
	config.SetLogLevel()

}

func main() {

	client := lib.CreateNewClient(config)

	log.Infof("Preparing requests")
	for _, request := range config.Requests {
		lib.PrepareRequests(client, request)
	}
	log.Infof("Launching requests")

	for i := 0; i < config.Workers; i++ {
		go lib.Perform(client)
	}
	time.Sleep(30 * time.Second)
}
