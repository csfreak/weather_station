package weather_station

import "time"
import "strconv"

type Weather struct {
	StationType	string
	Model string
	Timestamp time.Time
	IndoorTempF float32
	IndoorHumidity float32
	OutdoorTempF float32
	OutdoorHumidity float32
	WindspeedMPH uint8
	WindGustMPH uint8
	WindDirectionDegree uint16
	RainRateInch float32
	LightLevelLux float32
	UV uint8
}

const timeformat = "2006-01-02 03:04:05"

func FromEcowitt(d map) (*Weather, error) {
	w := &Weather{
		StationType: d["stationtype"],
		Model: d["model"]
	}
	time, err := time.Parse(timeformat, d["dateutc"])
	if err != nil {
		return nil, err
	}
	w.Timestamp = time

	val, err := strconv.ParseFloat(d["tempinf", 32])
	if err != nil {
		return nil, err
	}
	w.IndoorTempF = val

	val, err := strconv.ParseFloat(d["tempf", 32])
	if err != nil {
		return nil, err
	}
	w.OutdoorTempF = val

	val, err := strconv.ParseFloat(d["humidityin", 32])
	if err != nil {
		return nil, err
	}
	w.IndoorHumidity = val

	val, err := strconv.ParseFloat(d["humidity", 32])
	if err != nil {
		return nil, err
	}
	w.OutdoorHumidity = val

	

	return w, nil
}
