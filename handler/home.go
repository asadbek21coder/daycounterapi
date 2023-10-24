package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

type TemplateData struct {
	DaysRemaining int
}

func GetHomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	target := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	remain := int(target.Sub(time.Now()).Hours() / 24)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if remain < 0 {
		json.NewEncoder(w).Encode(map[string]string{"msg": "expired date"})
		return
	}
	if remain == 0 {
		json.NewEncoder(w).Encode(map[string]string{"msg": "today"})
		return
	}
	json.NewEncoder(w).Encode(map[string]int{"days_remain": remain})
}
