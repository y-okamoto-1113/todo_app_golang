package config

import (
	"log"
	"todo_app_golang/utils"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string // サーバーport番号
	SQLDriver string // SQLドライバの名前
	DbName    string // データベース名
	LogFile   string // ログファイル名
	Static    string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
		Static:    cfg.Section("web").Key("static").String(),
	}
}
