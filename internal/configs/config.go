package configs

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type backend struct {
	Name            string
	Destination_URL string
}

type server struct {
	Host string
	Port string
}

type configuration struct {
	Server   server
	Backends []backend
}

var Config *configuration

func NewConfiguration() (*configuration, error) {
	viper.AddConfigPath("data")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading config file: %s", err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %s", err)
	}

	return Config, nil
}
