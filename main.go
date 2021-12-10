package main

import (
	"fmt"
	"net/http"

	"github.com/csfreak/weather_station/weather_station/weather"
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
	if !(req.Method == "POST") {
		res.WriteHeader(405)
	} else {
		req.ParseForm()
		fmt.Println(req.PostForm)
		w := weather.FromEcowitt(req.PostForm)
		fmt.Println(w)
	}
}
