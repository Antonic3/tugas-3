package handler

import (
	"encoding/json"
	"net/http"
	"time"
	"tugas-3/service"
	util "tugas-3/utils"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := service.GenerateRandomStatusValue()

	err := util.WriteStatusFile(status, "status.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assignment/assignment-3/static/index.html")

}

// reload status every 15 seconds
func StatusAutoReload() {
	ticker := time.NewTicker(15 * time.Second)
	go func() {
		for range ticker.C {
			status := service.GenerateRandomStatusValue()
			err := util.WriteStatusFile(status, "status.json")
			if err != nil {
				panic(err)
			}
		}
	}()
}
