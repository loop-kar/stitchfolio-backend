package config

import (
	"io"

	"github.com/hashicorp/go-multierror"
	"github.com/spf13/viper"
)

func LoadConfig(reader io.Reader, envVarBindings map[string]string, appConfig interface{}) error {

	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	//fmt.Println(envVarBindings)

	err := bind(envVarBindings)
	if err != nil {
		return err
	}

	err = viper.ReadConfig(reader)
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&appConfig)
	if err != nil {
		return err
	}

	return nil
}

func bind(envVarBindings map[string]string) error {
	var bindError error
	for key, envVar := range envVarBindings {
		if err := viper.BindEnv(key, envVar); err != nil {
			bindError = multierror.Append(bindError, err)
		}
	}

	return bindError
}
