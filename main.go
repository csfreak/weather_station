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

	http.HandleFunc("/v1/ecowitt", weather.EcowittHandler)
	http.HandleFunc("/v1/ecowitt/", weather.EcowittHandler)
	http.HandleFunc("/v1/stations", weather.RestGetHandler)
	http.HandleFunc("/v1/stations/", weather.RestGetHandler)

	http.Handle("/metrics", promhttp.Handler())

	log.Printf("Starting HTTP Server on %s", apiListenerPort)

	//Log and Exit if http server exits
	log.Fatal(http.ListenAndServe(apiListenerPort, nil))
}
