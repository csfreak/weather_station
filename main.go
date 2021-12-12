package main

import (
	"fmt"
	"net/http"

	"github.com/csfreak/weather_station/weather"
)

func main() {
	fmt.Println("Staring main func")
	http.HandleFunc("/v1/ecowitt", ecowittHandler)
	http.HandleFunc("/v1/ecowitt/", ecowittHandler)

	fmt.Println("Starting HTTP Server")
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
	fmt.Println("After HTTP Server")
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
			fmt.Println(err)
		} else {
			fmt.Println(w)
			res.WriteHeader(http.StatusOK)
		}
	}
}
