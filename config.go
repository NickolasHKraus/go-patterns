package main

// Adding configuration.
//
// Program configuration can be achieved using Viper.
//
// See: https://github.com/spf13/viper

import (
	"bytes"

	"github.com/spf13/viper"
)

// See: https://go.dev/doc/effective_go#init
func init() {
	initConfig()
}

// Instead of initializing the configuration in main(), it is useful to define
// a global variable that can be accessed anywhere in the package and
// overridden during testing. This variable can also be exported if needed.
var Config *viper.Viper

func initConfig() {
	var yaml = []byte(`key: value
list:
  - item1
  - item2
map:
  key1: value1
  key2: value2
`)
	Config = viper.New()
	Config.SetConfigType("yaml")
	Config.ReadConfig(bytes.NewBuffer(yaml))
}

func getAllSettings() map[string]interface{} {
	return Config.AllSettings()
}
