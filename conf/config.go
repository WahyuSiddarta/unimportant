package conf

import (
	"fmt"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

func GetConfig() (*viper.Viper, error) {
	v := viper.New()
	i := "local"
	var err error
	if i == "local" {
		v.SetConfigName("config")
		v.AddConfigPath(".")
		v.SetConfigType("json")
		//apix.Config.AutomaticEnv()
		err = v.ReadInConfig()
	} else {
		v.AddRemoteProvider("consul", "192.168.0.217:8500", "dwb-config-local")
		v.SetConfigType("json") // because there is no file extension in a stream of bytes,  supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop"
		err = v.ReadRemoteConfig()
		if err != nil {
			fmt.Println(err)
		}
	}
	return v, err
}
