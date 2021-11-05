package controllers

import "net/http"

func SignupHandler(w http.ResponseWriter, r *http.Request){
	generateHTML(w, "signup handler", "layout", "public_navbar", "signup")
}
