package main

import (
	"fmt"
	"log"
	"todo_app_golang/app/controllers"
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
	// 	Name:     "test1",
	// 	Email:    "test1@example.com",
	// 	Password: "test1",
	// }
	// fmt.Println("u =>", u)
	// u.CreateUser()

	// u, _ := models.FindUser(1)
	// fmt.Println("u =>", u)
	// u.CreateTodo("First Todo")
	// u.CreateTodo("Second Todo")
	// u.CreateTodo("Third Todo")

	// t, err := models.FindTodo(1)
	// fmt.Println("err =>", err)
	// fmt.Println("t =>", t)

	// todos, _ := models.FindAllTodos()
	// for index, todo := range todos {
	// 	fmt.Println(index, todo)
	// }

	// t, _ := models.FindTodo(1)
	// t.Content = "this content is updated by user"
	// t.UpdateTodo()
	// fmt.Println(t.UpdatedAt)

	// u, _ := models.FindUser(1)
	// _ = u.CreateTodo("todo to be deleted")
	// todos, _ := u.FindTodosByUser()
	// fmt.Println("todos before delete =>", todos)
	// last_index := len(todos) - 1
	// t := todos[last_index]
	// t.DeleteTodo()
	// todos, _ = u.FindTodosByUser()
	// fmt.Println("todos after delete =>", todos)

	err := controllers.StartMainServer()
	if err != nil {
		log.Fatalln(err)
	}

	// user, err := models.FindUserByEmail("test1@example.com")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println("user =>", user)
	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println("session =>", session)
	// valid, _ := session.CheckSession()
	// fmt.Println(valid)

}
