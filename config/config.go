package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var c *viper.Viper

func init() {
	c = viper.New()
	c.SetConfigType("yaml")
	c.SetEnvPrefix("APP")
	c.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	c.AutomaticEnv()
}

func Init(name string) {

	c.SetConfigName(name)
	c.SetConfigType("yaml")
	c.AddConfigPath("./deploy")
	err := c.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func GetString(key string) string {
	return c.GetString(key)
}

func GetBool(key string) bool {
	return c.GetBool(key)
}
