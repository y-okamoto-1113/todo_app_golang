package main

import (
	"fmt"
	"log"
	"todo_app_golang/config"
)

func main() {
	fmt.Println("config.Config.Port =>", config.Config.Port)
	fmt.Println("config.Config.SQLDriver =>", config.Config.SQLDriver)
	fmt.Println("config.Config.DbName =>", config.Config.DbName)
	fmt.Println("config.Config.LogFile =>", config.Config.LogFile)

	log.Println(1111)
}
