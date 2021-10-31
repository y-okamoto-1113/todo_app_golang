package config

type ConfigList struct {
	Port      string // サーバーport番号
	SQLDriver string // SQLドライバの名前
	DbName    string // データベース名
	LogFile   string // ログファイル名
}

var Config ConfigList
