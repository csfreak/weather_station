package weather

import (
	"time"
)

type Weather struct {
	StationType         string
	Model               string
	Timestamp           time.Time
	IndoorTempF         float32
	IndoorHumidity      float32
	OutdoorTempF        float32
	OutdoorHumidity     float32
	WindspeedMPH        float32
	WindGustMPH         float32
	WindDirectionDegree uint16
	RainRateInch        float32
}
