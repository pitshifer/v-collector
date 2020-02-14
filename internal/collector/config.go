package collector

// Config structure
type Config struct {
	DatabaseURL string `toml:"databaseURL"`
	MQTT        MQTTConfig
}

// MQTTConfig for MQTT broker
type MQTTConfig struct {
	URL      string
	Username string
	Password string
	ClientID string
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
