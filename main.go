package main

import (
	"fmt"
	"todo_app_golang/app/models"
)

func main() {
	// fmt.Println("config.Config.Port =>", config.Config.Port)
	// fmt.Println("config.Config.SQLDriver =>", config.Config.SQLDriver)
	// fmt.Println("config.Config.DbName =>", config.Config.DbName)
	// fmt.Println("config.Config.LogFile =>", config.Config.LogFile)

	// log.Println(1111)

	fmt.Println(models.Db)
}
