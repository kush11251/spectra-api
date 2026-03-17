package controllers

import (
    "encoding/json"
    "net/http"
)

func GetMetrics(w http.ResponseWriter, r *http.Request) {
    metrics := map[string]int{
        "cpu": 50,
        "memory": 70,
    }
    json.NewEncoder(w).Encode(metrics)
}