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

	u, _ := models.FindUser(3)
	fmt.Println("u =>", u)

	u.Name = "test3"
	u.Email = "test3@example.com"
	u.UpdateUser()
	u, _ = models.FindUser(3)
	fmt.Println(u)

	u.DeleteUser()
	u, _ = models.FindUser(3)
	fmt.Println(u)

}
