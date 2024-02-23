package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_client(t *testing.T) {

	q := WeatherQuery{Lat: 9.1633, Long: -83.7484}

	if !clientcall(&q) {
		t.Errorf("Code:%d", q.Code)
	}

}

func Test_handler(t *testing.T) {

	q := WeatherQuery{Lat: 9.1633, Long: 83.7484}

	var jsonbuffer = []byte(`{
		"latitude":9.1633,
		"longitude": -83.7484
	}`)

	req := httptest.NewRequest("POST", "http://localhost", bytes.NewBuffer(jsonbuffer))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w := httptest.NewRecorder()
	weatherHandler(w, req)
	resp := w.Result()
	if err := json.NewDecoder(resp.Body).Decode(&q); err != nil {
		t.Fatalf("Unable to decode Json:%s", err.Error())
		return
	}
	if q.Code != http.StatusOK {
		t.Fatalf("Http Error Code:%d:%s", q.Code, q.Condition)
		return
	}
}
