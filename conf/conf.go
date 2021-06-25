package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type Cfg struct {
	Debug          bool
	Env            string
	LogFile        string
	LogLevel       string
	MessageChannel string
	Frontend       Site
	API            API
	MySQL          map[string]MySQL
	Redis          Redis
	Kafka          map[string]Kafka
	Path           string
}

type MySQL struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

type Site struct {
	Host      string
	Logfile   string
	Cache     bool
}

type API struct {
	Key         string
	Secret      string
	AesKey      string
	AesKey256   string
}

type Redis struct {
	Host     string
	Password string
}
type Kafka struct {
	Topic     string
	Group     string
	Broker    []string
	Zookeeper string
}

var Config *Cfg

func Init() (*Cfg, error) {
	return InitByPath("./conf")
}

func InitByPath(path string) (*Cfg, error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("can't read config file: %s", err)
	}

	c := Cfg{}
	if err := viper.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("unable to unmarshal config: %s", err)
	}

	c.Path = path
	Config = &c
	return &c, nil
}
