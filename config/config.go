package config

import (
	"github.com/crazyhl/gopermission/base_struct"
)

var userConfig *base_struct.Config

func init() {
	userConfig = &base_struct.Config{}
}

func GetConfig() *base_struct.Config {
	return userConfig
}
