package main

import (
	"net/http"
	"tugas-3/handler"
)

func main() {
	handler.StatusAutoReload()

	http.HandleFunc("/status", handler.StatusHandler)
	http.HandleFunc("/", handler.IndexHandler)

	http.ListenAndServe(":8080", nil)
}
