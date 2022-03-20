package configs

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func InitConfigFile() error {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./configs/")
	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "Initilize config file error")
	}
	return nil
}
