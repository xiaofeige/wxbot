package server

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	wechat_work_bot "github.com/xiaofeige/wechat-work-bot"
)

var (
	ConfigPath = "./config.toml"
)

type Config struct {
	ServerConf struct {
	}

	RobotConf wechat_work_bot.RobotConfig
}

// 使用viper
func init() {
	flag.StringVar(&ConfigPath, "config", "./config.toml", "/path/to/your/config.toml")
	flag.Parse()

	viper.SetConfigFile(ConfigPath)

	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
