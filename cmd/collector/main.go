package main

import (
	"flag"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/pitshiver/v-collector/internal/collector"
	"github.com/sirupsen/logrus"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "c", "config.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := collector.NewConfig()
	_, err := toml.DecodeFile(configFile, &config)
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Printf("%+v\n", config)

	collector := collector.CreateCollector(config)
	collector.Run()
}
