package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func TopHandler(w http.ResponseWriter, r *http.Request) {
	_, err := getSession(w, r)
	if err != nil {
		generateHTML(w, "Hello World from generateHTML function", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	session, err := getSession(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := session.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}
		todos, _ := user.FindTodosByUser()
		user.Todos = todos
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}
