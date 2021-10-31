package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	multiLogFile := io.MultiWriter(os.Stdout, logfile)   // 標準出力とlogfileにログを出力する様に設定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // ログ出力のフォーマット設定。`Lshortfile` はログ出力が実行されたファイル名と行数を出力
	log.SetOutput(multiLogFile)                          // ログの出力先を設定

}
