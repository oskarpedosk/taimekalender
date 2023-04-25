package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "taimekalender active",
		Version: "1.0.0",
	}

	_ = writeJson(w, http.StatusOK, payload, nil)
}
