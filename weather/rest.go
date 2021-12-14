package weather

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type JSONResponse struct {
	Stations []JSONWeather
}

type JSONWeather struct {
	StationType             string `json:"StationType"`
	Model                   string `json:"Model"`
	Timestamp               string `json:"LastUpdatedAt"`
	IndoorTempF             string `json:"IndoorTemperatureF"`
	IndoorHumidity          string `json:"IndoorHumidity"`
	OutdoorTempF            string `json:"OutdoorTemperatureF"`
	OutdoorHumidity         string `json:"OutdoorHumidity"`
	WindspeedMPH            string `json:"WindspeedMPH"`
	WindGustMPH             string `json:"WindGustMPH"`
	WindDirectionDegree     string `json:"WindDirectionDegree"`
	RainRateInch            string `json:"RainRateInch"`
	BarometricPressureABSIn string `json:"AbsoluteBarometricPressureInHg"`
	BarometricPressureRelIn string `json:"RelativeBarometricPressureInHg"`
}

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

func weatherToJSON(w Weather) JSONWeather {
	var output JSONWeather
	output.StationType = w.StationType
	output.Model = w.Model
	output.Timestamp = w.Timestamp.Format(timeformat)
	output.IndoorTempF = strconv.FormatFloat(w.IndoorTempF, 'f', 1, 64)
	output.IndoorHumidity = strconv.FormatFloat(w.IndoorHumidity, 'f', 1, 64)
	output.OutdoorTempF = strconv.FormatFloat(w.OutdoorTempF, 'f', 1, 64)
	output.OutdoorHumidity = strconv.FormatFloat(w.OutdoorHumidity, 'f', 1, 64)
	output.WindspeedMPH = strconv.FormatFloat(w.WindspeedMPH, 'f', 1, 64)
	output.WindGustMPH = strconv.FormatFloat(w.WindGustMPH, 'f', 1, 64)
	output.WindDirectionDegree = strconv.FormatUint(uint64(w.WindDirectionDegree), 10)
	output.RainRateInch = strconv.FormatFloat(w.RainRateInch, 'f', 1, 64)
	output.BarometricPressureABSIn = strconv.FormatFloat(w.BarometricPressureABSIn, 'f', 1, 64)
	output.BarometricPressureRelIn = strconv.FormatFloat(w.BarometricPressureRelIn, 'f', 1, 64)

	return output
}

func RestGetHandler(res http.ResponseWriter, req *http.Request) {
	if !(req.Method == http.MethodGet) {
		res.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		var output JSONResponse
		for _, station := range Stations {
			output.Stations = append(output.Stations, weatherToJSON(station.LastData))
		}
		err := json.NewEncoder(res).Encode(output)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
	}
}
