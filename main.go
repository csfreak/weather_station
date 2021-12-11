package main

import (
	"fmt"
	"net/http"

	"github.com/csfreak/weather_station/weather"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/ecowitt", ecowittHandler)
	mux.HandleFunc("/v1/ecowitt/", ecowittHandler)

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	s.ListenAndServe()
}

func ecowittHandler(res http.ResponseWriter, req *http.Request) {
	if !(req.Method == http.MethodPost) {
		res.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		req.ParseForm()
		fmt.Println(req.PostForm)
		w, err := weather.FromEcowitt(req.PostForm)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Println(w)
		res.WriteHeader(http.StatusOK)
	}
}
