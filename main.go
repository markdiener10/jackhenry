package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type WeatherQuery struct {
	Lat       float64 `json:"latitude"`
	Long      float64 `json:"longitude"`
	Code      int     `json:"code,omitempty"`
	Stamp     string  `json:"timestamp,omitempty"`
	Condition string  `json:"condition,omitempty"`
}

const (
	APIKEY = "TBD"
)

func clientcall(q *WeatherQuery) bool {

	queryString := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%.4f&lon=%.4f&units=imperial&appid=%s", q.Lat, q.Long, APIKEY)

	resp, err := http.Get(queryString)
	if err != nil {
		q.Code = http.StatusInternalServerError
		q.Condition = err.Error()
		return false
	}
	defer resp.Body.Close()
	q.Code = resp.StatusCode
	if q.Code != http.StatusOK {
		q.Condition = "Error:NotOk"
		return false
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		q.Condition = "Error:" + err.Error()
		return false
	}
	var js map[string]any
	err = json.Unmarshal(body, &js)
	if err != nil {
		q.Code = http.StatusInternalServerError
		q.Condition = "Error:" + err.Error()
		return false
	}

	mainany, ok := js["main"]
	if !ok {
		q.Code = http.StatusInternalServerError
		q.Condition = "Error: Unable to get temperature"
		return false
	}
	main, ok := mainany.(map[string]any)
	if !ok {
		q.Code = http.StatusInternalServerError
		q.Condition = "Error: Unable to get temperature"
		return false
	}

	tempany, ok := main["temp"]
	if !ok {
		q.Code = http.StatusInternalServerError
		q.Condition = "Error: Unable to get temperature"
		return false
	}
	temp, ok := tempany.(float64)
	if !ok {
		q.Code = http.StatusInternalServerError
		q.Condition = "Error: Unable to get temperature"
		return false
	}

	switch {
	case temp > 100.0:
		q.Condition = "Very Hot"
	case temp > 90:
		q.Condition = "Hot"
	case temp > 70:
		q.Condition = "Warm"
	case temp > 50:
		q.Condition = "Brisk"
	case temp > 32:
		q.Condition = "Cold"
	default:
		q.Condition = "Bloody Cold"
	}
	return true
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	var q WeatherQuery
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !clientcall(&q) {
		http.Error(w, "error", q.Code)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(q)
}

func main() {
	log.Println("Jack Henry code sample starting")
	mux := http.NewServeMux()
	mux.HandleFunc("/", weatherHandler)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
	log.Println("Jack Henry code sample terminated")
}
