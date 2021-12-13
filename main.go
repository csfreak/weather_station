package main

import (
	"fmt"
	"github.com/csfreak/weather_station/weather"
	"log"
	"net/http"
)

var apiListenerPort = ":8080"

func main() {
	http.HandleFunc("/v1/ecowitt", ecowittHandler)
	http.HandleFunc("/v1/ecowitt/", ecowittHandler)

	fmt.Println("Starting HTTP Server on", apiListenerPort)
	log.Fatal(http.ListenAndServe(apiListenerPort, nil))
}

func ecowittHandler(res http.ResponseWriter, req *http.Request) {
	if !(req.Method == http.MethodPost) {
		res.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		parseErr := req.ParseForm()
		if parseErr != nil {
			fmt.Println(parseErr)
			return
		}
		fmt.Println(req.PostForm)
		w, err := weather.FromEcowitt(req.PostForm)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
		} else {
			fmt.Println(w)
			res.WriteHeader(http.StatusOK)
		}
	}
}
