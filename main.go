package main

import (
	"log"
	"net/http"

	"github.com/csfreak/weather_station/weather"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var apiListenerPort = ":8080"

func main() {
	weather.MetricInit()

	http.HandleFunc("/v1/ecowitt", ecowittHandler)
	http.HandleFunc("/v1/ecowitt/", ecowittHandler)

	http.Handle("/metrics", promhttp.Handler())

	log.Printf("Starting HTTP Server on %s", apiListenerPort)

	//Log and Exit if http server exits
	log.Fatal(http.ListenAndServe(apiListenerPort, nil))
}

func ecowittHandler(res http.ResponseWriter, req *http.Request) {
	if !(req.Method == http.MethodPost) {
		res.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		parseErr := req.ParseForm()
		if parseErr != nil {
			log.Println(parseErr)
			return
		}
		w, err := weather.FromEcowitt(req.PostForm)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		} else {
			weather.UpdateMetrics(w)
			res.WriteHeader(http.StatusOK)
		}
	}
}
