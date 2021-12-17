package weather

import (
	"encoding/json"
	"net/http"
)

type Station struct {
	Type     string
	Model    string
	LastData Weather
	Host     string
}

var (
	Stations []Station
)

func UpdateStation(w *Weather, host string) {
	for i := range Stations {
		station := Stations[i]
		if station.Host == host && station.Type == w.StationType && station.Model == w.Model {
			Stations[i].LastData = *w
			return
		}
	}
	Stations = append(Stations, Station{
		Type:     w.StationType,
		Model:    w.Model,
		LastData: *w,
		Host:     host,
	})
}

func RestGetHandler(res http.ResponseWriter, req *http.Request) {
	if !(req.Method == http.MethodGet) {
		res.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		var stationdata []Weather
		for _, station := range Stations {
			stationdata = append(stationdata, station.LastData)
		}
		err := json.NewEncoder(res).Encode(stationdata)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
	}
}
