package weather

import (
	"time"
)

type Weather struct {
	StationType             string
	Model                   string
	Timestamp               time.Time
	IndoorTempF             float64
	IndoorHumidity          float64
	OutdoorTempF            float64
	OutdoorHumidity         float64
	WindspeedMPH            float64
	WindGustMPH             float64
	WindDirectionDegree     uint16
	RainRateInch            float64
	BarometricPressureABSIn float64
	BarometricPressureRelIn float64
}
