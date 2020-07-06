package lib

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configuration struct {
	Requests    []Request
	Workers     int
	ChanWorkers chan []bool
	LogLevel    int64
	Proxy       string
}

type Request struct {
	Path    string
	Method  string
	Headers map[string]string
	Params  []Param
}

type Param struct {
	Name  string
	Value string
	Type  string
}

func GetConfig() *Configuration {
	var config Configuration
	var err error

	viper.AddConfigPath(".")
	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	if err = viper.Unmarshal(&config); err != nil {
		log.Fatalf("Config file not formatted properly %s \n", err)
	}

	return &config
}

func (c *Configuration) SetLogLevel() {
	logLevels := map[int64]log.Level{0: log.WarnLevel, 1: log.InfoLevel, 2: log.DebugLevel}
	log.SetLevel(logLevels[c.LogLevel])
}
