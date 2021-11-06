package controllers

import (
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
		if r.Method == "POST" {
			TodoCreateHandler(w, r)
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
}

func TodoNewHandler(w http.ResponseWriter, r *http.Request) {
	_, err := getSession(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

func TodoCreateHandler(w http.ResponseWriter, r *http.Request) {
	s, err := getSession(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 302)
	} else {
		u, err := s.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}
		content := r.PostFormValue("content")
		if err = u.CreateTodo(content); err != nil {
			log.Fatalln(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}

}
