package health

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Health struct {
	Status string `json:"status"`
	UpTime string `json:"uptime"`
}

var startTime = time.Now()

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	totalTime := time.Since(startTime)
	days := int(totalTime.Hours()) / 24
	hours := int(totalTime.Hours()) % 24
	minutes := int(totalTime.Minutes()) % 60
	seconds := int(totalTime.Seconds()) % 60

	health := Health{
		Status: "ok",
		UpTime: fmt.Sprintf("App Uptime: %dd %dh %dm %ds", days, hours, minutes, seconds),
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(health); err != nil {
		log.Println("Encoding Failed")
		w.WriteHeader(http.StatusInternalServerError)
	}
}
