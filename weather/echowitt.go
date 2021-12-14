package weather

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const timeformat = "2006-01-02 15:04:05"

func fromEcowitt(d url.Values) (*Weather, error) {
	w := &Weather{
		StationType: d.Get("stationtype"),
		Model:       d.Get("model"),
	}
	var err error

	timeParse, err := time.Parse(timeformat, d.Get("dateutc"))
	if err != nil {
		return nil, err
	}
	w.Timestamp = timeParse

	tempin, err := strconv.ParseFloat(d.Get("tempinf"), 64)
	if err != nil {
		return nil, err
	}
	w.IndoorTempF = tempin

	temp, err := strconv.ParseFloat(d.Get("tempf"), 64)
	if err != nil {
		return nil, err
	}
	w.OutdoorTempF = temp

	humidin, err := strconv.ParseFloat(d.Get("humidityin"), 64)
	if err != nil {
		return nil, err
	}
	w.IndoorHumidity = humidin

	humid, err := strconv.ParseFloat(d.Get("humidity"), 64)
	if err != nil {
		return nil, err
	}
	w.OutdoorHumidity = humid

	winds, err := strconv.ParseFloat(d.Get("windspeedmph"), 64)
	if err != nil {
		return nil, err
	}
	w.WindspeedMPH = winds

	windg, err := strconv.ParseFloat(d.Get("windgustmph"), 64)
	if err != nil {
		return nil, err
	}
	w.WindGustMPH = windg

	windd, err := strconv.ParseInt(d.Get("winddir"), 10, 16)
	if err != nil {
		return nil, err
	}
	w.WindDirectionDegree = uint16(windd)

	rain, err := strconv.ParseFloat(d.Get("rainratein"), 64)
	if err != nil {
		return nil, err
	}
	w.RainRateInch = rain

	apressure, err := strconv.ParseFloat(d.Get("baromabsin"), 64)
	if err != nil {
		return nil, err
	}
	w.BarometricPressureABSIn = apressure

	rpressure, err := strconv.ParseFloat(d.Get("baromrelin"), 64)
	if err != nil {
		return nil, err
	}
	w.BarometricPressureRelIn = rpressure

	return w, nil
}

func EcowittHandler(res http.ResponseWriter, req *http.Request) {
	if !(req.Method == http.MethodPost) {
		res.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		parseErr := req.ParseForm()
		if parseErr != nil {
			log.Println(parseErr)
			return
		}
		w, err := fromEcowitt(req.PostForm)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		} else {
			UpdateMetrics(w)
			UpdateStation(w, req.RemoteAddr)
			res.WriteHeader(http.StatusOK)
		}
	}
}
