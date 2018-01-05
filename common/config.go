package common

import (
	cfg "github.com/Unknwon/goconfig"
)

var Cfg *cfg.ConfigFile

func InitConfig() {
	Cfg, _ = cfg.LoadConfigFile("config.ini")
}
