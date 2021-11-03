package controllers

import (
	"html/template"
	"log"
	"net/http"
)

func TopHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("app/views/templates/top.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, "Hello World!!!") // 第2引数はHTMLに渡せる。HTML側で`{{.}}`で展開できる
}
