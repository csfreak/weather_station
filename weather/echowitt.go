package weather

import (
	"net/url"
	"strconv"
	"time"
)

const timeformat = "2006-01-02 15:04:05"

func FromEcowitt(d url.Values) (*Weather, error) {
	w := &Weather{
		StationType: d.Get("stationtype"),
		Model:       d.Get("model"),
	}
	var err error

	time, err := time.Parse(timeformat, d.Get("dateutc"))
	if err != nil {
		return nil, err
	}
	w.Timestamp = time

	tempin, err := strconv.ParseFloat(d.Get("tempinf"), 32)
	if err != nil {
		return nil, err
	}
	w.IndoorTempF = tempin

	temp, err := strconv.ParseFloat(d.Get("tempf"), 32)
	if err != nil {
		return nil, err
	}
	w.OutdoorTempF = temp

	humidin, err := strconv.ParseFloat(d.Get("humidityin"), 32)
	if err != nil {
		return nil, err
	}
	w.IndoorHumidity = humidin

	humid, err := strconv.ParseFloat(d.Get("humidity"), 32)
	if err != nil {
		return nil, err
	}
	w.OutdoorHumidity = humid

	winds, err := strconv.ParseFloat(d.Get("windspeedmph"), 32)
	if err != nil {
		return nil, err
	}
	w.WindspeedMPH = winds

	windg, err := strconv.ParseFloat(d.Get("windgustmph"), 32)
	if err != nil {
		return nil, err
	}
	w.WindGustMPH = windg

	windd, err := strconv.ParseInt(d.Get("winddir"), 10, 16)
	if err != nil {
		return nil, err
	}
	w.WindDirectionDegree = uint16(windd)

	rain, err := strconv.ParseFloat(d.Get("rainratein"), 32)
	if err != nil {
		return nil, err
	}
	w.RainRateInch = rain

	apressure, err := strconv.ParseFloat(d.Get("baromabsin"), 32)
	if err != nil {
		return nil, err
	}
	w.BarometricPressureABSIn = apressure

	rpressure, err := strconv.ParseFloat(d.Get("baromrelin"), 32)
	if err != nil {
		return nil, err
	}
	w.BarometricPressureRelIn = rpressure

	return w, nil
}
