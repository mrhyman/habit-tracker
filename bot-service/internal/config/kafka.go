package config

type KafkaConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func (kc KafkaConfig) GetHost() string {
	return kc.Host
}

func (kc KafkaConfig) GetPort() int {
	return kc.Port
}
