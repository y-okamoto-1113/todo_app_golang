package controllers

import (
	"net/http"
	"todo_app_golang/config"
)

func StartMainServer() error {
	http.HandleFunc("/", TopHandler)
	return http.ListenAndServe(":" + config.Config.Port, nil)
}
