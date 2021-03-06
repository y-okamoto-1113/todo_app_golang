package controllers

import (
	"fmt"
	"log"
	"net/http"
	"todo_app_golang/app/models"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := getSession(w, r)
		if err != nil {
			generateHTML(w, "signup handler", "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/todos", 302)
		}
	} else if r.Method == "POST" {
		fmt.Println("POST SignupHandler before parser =>\n", r)
		err := r.ParseForm() // `Form` で送られたデータをパースする。
		fmt.Println("POST SignupHandler after parse =>\n", r)
		if err != nil {
			log.Fatalln(err)
		}
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("Password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/", 302) // 第3引数にリダイレクト先。第4引数にステータスコード。
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	_, err := getSession(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "login")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

func AuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := models.FindUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 302)
	}

	if user.Password == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}

	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		err = session.DeleteSessionByUUID()
		if err != nil {
			log.Fatalln(err)
		}
	}
	http.Redirect(w, r, "/login", 302)
}
