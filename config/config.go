package config

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"sync"
)

type Config struct {
	LogLevel        string
	ServerPort      string
	VideoReportPort string
	VideoRecommPort string
}

var Cfg Config

var cfgOnce sync.Once

func init() {
	getConfig()
}

func getConfig() {
	cfgOnce.Do(func() {
		c, err := goconfig.LoadConfigFile("./conf.ini")
		if err != nil {
			fmt.Errorf("init config file fail :%v", err)
			return
		}
		Cfg.LogLevel, _ = c.GetValue(goconfig.DEFAULT_SECTION, "logLevel")
		Cfg.ServerPort, _ = c.GetValue(goconfig.DEFAULT_SECTION, "serverPort")
		Cfg.VideoReportPort, _ = c.GetValue("rpc-port", "videoReportPort")
		Cfg.VideoRecommPort, _ = c.GetValue("rpc-port", "videoRecommPort")
	})
}
