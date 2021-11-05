package controllers

import (
	"net/http"
)

func TopHandler(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "Hello World from generateHTML function", "layout", "public_navbar", "top")
}
