//go:build !windows
// +build !windows

package Exercise_1

import (
	"encoding/json"
)

func getConfig() Config {
	data, _ := ioutil.ReadFile("C:\\Users\\vitno\\Documents\\Config\\config.json")
	var c Config
	_ = json.Unmarshal(data, &c)
	return c
}
