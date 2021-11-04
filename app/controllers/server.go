package controllers

import (
	"fmt"
	"net/http"
	"todo_app_golang/config"
)

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	fmt.Println("files =>", files) // => &{app/views}
	// localhost/static/css/bootstrap.min.css みたいに `/static/` にリクエストが来ると、`static` を無視して `app/views/` 以下の `css/bootstrap.min.css` を見るようにする。
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", TopHandler)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
