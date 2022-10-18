package main

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/spf13/viper"
)

func TestConfig(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		exp := map[string]interface{}{
			"key": "value",
			"list": []interface{}{
				"item1",
				"item2",
			},
			"map": map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
		}
		ret := getAllSettings()
		if !reflect.DeepEqual(ret, exp) {
			t.Errorf("bad config.\nExpected: %s.\nGot: %5s%s", exp, " ", ret)
		}
	})
	t.Run("mock config", func(t *testing.T) {
		var mockYAML = []byte(`mock: mock`)
		mockConfig := viper.New()
		mockConfig.SetConfigType("yaml")
		mockConfig.ReadConfig(bytes.NewBuffer(mockYAML))
		Config = mockConfig
		exp := map[string]interface{}{"mock": "mock"}
		ret := getAllSettings()
		if !reflect.DeepEqual(ret, exp) {
			t.Errorf("bad config.\nExpected: %s.\nGot: %5s%s", exp, " ", ret)
		}
	})
}
