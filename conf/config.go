package conf

import "github.com/spf13/viper"

type AppConfig struct {
	JWTConfig JWTConfig `mapstructure:"jwt_op"`
}

var AppConf AppConfig

func init() {
	v := viper.New()
	configName := "dev-config.yaml"
	v.SetConfigFile(configName)
	v.ReadInConfig()
	err := v.Unmarshal(&AppConf)
	if err != nil {
		panic(err)
	}
}
