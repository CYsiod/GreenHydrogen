package config

type Mqtt struct {
	Broker    string `json:"broker" yaml:"broker"`
	UserName  string `json:"userName" yaml:"userName"`
	PassWord  string `json:"passWord" yaml:"pass_word"`
	ClientID  string `mapstructure:"client_id" json:"client_id" yaml:"client_id"`
	Topic     string `json:"topic" yaml:"topic"`
	KeepAlive int64  `mapstructure:"keep_alive" json:"keep_alive" yaml:"keep_alive"`
}
