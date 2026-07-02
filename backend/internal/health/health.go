package health

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type UptimeDetails struct {
	Days    int `json:"days"`
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
	Seconds int `json:"seconds"`
}

type HealthResponse struct {
	Status string        `json:"status"`
	UpTime UptimeDetails `json:"uptime"`
}

var startTime = time.Now()

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	totalTime := time.Since(startTime)
	days := int(totalTime.Hours()) / 24
	hours := int(totalTime.Hours()) % 24
	minutes := int(totalTime.Minutes()) % 60
	seconds := int(totalTime.Seconds()) % 60

	health := HealthResponse{
		Status: "ok",
		UpTime: UptimeDetails{
			Days:    days,
			Hours:   hours,
			Minutes: minutes,
			Seconds: seconds,
		},
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(health); err != nil {
		log.Println("Encoding Failed")
		w.WriteHeader(http.StatusInternalServerError)
	}
}
