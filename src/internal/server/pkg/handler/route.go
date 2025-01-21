package handler

import "net/http"

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /job", createJobHandler)
	mux.HandleFunc("GET /status/{id}", jobStatusHandler)

	return mux
}
