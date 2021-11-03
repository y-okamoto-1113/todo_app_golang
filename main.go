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

	fmt.Println("models.Db =>", models.Db)

	// u := &models.User{
	// 	Name:     "test",
	// 	Email:    "test@example.com",
	// 	Password: "testtest",
	// }
	// fmt.Println("u =>", u)
	// u.CreateUser()

	// u, _ := models.FindUser(1)
	// fmt.Println("u =>", u)
	// u.CreateTodo("First Todo")

	t, err := models.FindTodo(1)
	fmt.Println("err =>", err)
	fmt.Println("t =>", t)

}
