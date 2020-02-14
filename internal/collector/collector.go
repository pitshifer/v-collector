package collector

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

// Collector ...
type Collector struct {
	config     *Config
	mqttClient mqtt.Client
}

// CreateCollector ...
func CreateCollector(config *Config) *Collector {
	return &Collector{
		config: config,
	}
}

// Run ...
func (c *Collector) Run() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(c.config.MQTT.URL)
	opts.SetUsername(c.config.MQTT.Username)
	opts.SetPassword(c.config.MQTT.Password)
	opts.SetClientID(c.config.MQTT.ClientID)

	client := mqtt.NewClient(opts)

	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		logrus.Fatal(err)
	}

	s := client.Subscribe("/#", 0, func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("%s: %s\n", message.Topic(), string(message.Payload()))
	})
	if err := s.Error(); err != nil {
		fmt.Println(err)
	}

	timer := time.NewTicker(1 * time.Second)
	for _ = range timer.C {
	}
}
